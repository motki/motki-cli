package command

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"

	"github.com/motki/core/log"
	"github.com/motki/core/model"
	"github.com/motki/core/proto/client"

	"github.com/motki/cli"
	"github.com/motki/cli/text"
)

// ProductCommand provides an interactive manager for production chains.
type ProductCommand struct {
	character *model.Character
	corp      *model.Corporation
	corpID    int

	env    *cli.Prompter
	logger log.Logger
	client client.Client
}

func NewProductCommand(cl client.Client, p *cli.Prompter, logger log.Logger) ProductCommand {
	var corp *model.Corporation
	var char *model.Character
	var corpID int
	char, err := cl.CharacterForRole(model.RoleLogistics)
	if err == nil {
		corp, err = cl.GetCorporation(char.CorporationID)
		if err != nil {
			logger.Warnf("command: unable to get corporation details: %s", err.Error())
		} else {
			corpID = corp.CorporationID
		}
	}
	if err != nil && err != client.ErrNotAuthenticated {
		logger.Debugf("command: unable to load auth details: %s", err.Error())
	}
	return ProductCommand{
		char,
		corp,
		corpID,
		p,
		logger,
		cl}
}

func (c ProductCommand) RequiresAuth() bool {
	return true
}

func (c ProductCommand) Prefixes() []string {
	return []string{"product", "prod"}
}

func (c ProductCommand) Description() string {
	if c.corp == nil {
		return "Manipulate production chains for corpID 0."
	}
	return fmt.Sprintf("Manipulate production chains for %s.", c.corp.Name)
}

func (c ProductCommand) Handle(subcmd string, args ...string) {
	switch {
	case len(subcmd) == 0:
		c.PrintHelp()

	case subcmd == "new" || subcmd == "add" || subcmd == "create":
		c.newProduct(args...)

	case subcmd == "show":
		c.showProduct(args...)

	case subcmd == "list":
		c.listProducts()

	case subcmd == "edit":
		c.editProduct(args...)

	case subcmd == "view" || subcmd == "preview":
		c.previewProduct(args...)

	default:
		fmt.Printf("Unknown subcommand: %s\n", subcmd)
		c.PrintHelp()
	}
}

func (c ProductCommand) PrintHelp() {
	colWidth := 20
	fmt.Println()
	fmt.Println(fmt.Sprintf(`Command "%s" can be used to manipulate production chains.`, text.Boldf("product")))
	fmt.Println()
	fmt.Println(text.WrapText(fmt.Sprintf(`When invoking a subcommand, if the optional parameter is omitted, an interactive prompt will begin to collect the necessary details.`), text.StandardTerminalWidthInChars))
	fmt.Println()
	if c.corp != nil {
		fmt.Println(text.Boldf("Character linked!"))
		fmt.Println(fmt.Sprintf("You are logged in as %s for %s.",
			text.Boldf(c.character.Name),
			text.Boldf(c.corp.Name)))
		fmt.Println()
		fmt.Println(text.WrapText(`This command will operate on production chains for your corporation. Additionally, corporation-owned assets will be inventoried to display available and missing materials.`, text.StandardTerminalWidthInChars))
		fmt.Println()
	}
	fmt.Printf(`Subcommands:
  %s Preview production chains for a specific item type.
  %s Create a new production chain.

  %s List all production chains.
  %s Display details for a given production chain.
  %s Edit an existing production chain.
`,
		text.Boldf(text.PadTextRight("view [typeID]", colWidth)),
		text.Boldf(text.PadTextRight("add [typeID]", colWidth)),
		text.Boldf(text.PadTextRight("list", colWidth)),
		text.Boldf(text.PadTextRight("show [productID]", colWidth)),
		text.Boldf(text.PadTextRight("edit [productID]", colWidth)))
	fmt.Println()
}

// getProductName returns the given product's name.
func (c ProductCommand) getProductName(p *model.Product) string {
	t, err := c.client.GetItemType(p.TypeID)
	if err != nil {
		c.logger.Debugf("unable to get item name: %s", err.Error())
		return strconv.Itoa(p.ProductID)
	}
	return t.Name
}

// getRegionName returns the given region's name.
func (c ProductCommand) getRegionName(regionID int) string {
	r, err := c.client.GetRegion(regionID)
	if err != nil {
		c.logger.Debugf("unable to get region name: %s", err.Error())
		return strconv.Itoa(regionID)
	}
	return r.Name
}

func (c ProductCommand) getBlueprintIndex(p *model.Product) (map[*model.Product][]*model.Blueprint, []*model.Product) {
	if c.corpID == 0 {
		return nil, nil
	}
	needed := map[int]*model.Product{}
	var fillNeeded func(p *model.Product)
	fillNeeded = func(p *model.Product) {
		if p.Kind == model.ProductBuild {
			item, err := c.client.GetItemTypeDetail(p.TypeID)
			if err != nil {
				c.logger.Warnf("unable to get product item type detail: %s", err.Error())
				return
			}
			if item.BlueprintID == 0 {
				c.logger.Warnf("got blueprintID of 0 for typeID %d", p.TypeID)
				return
			}
			needed[item.BlueprintID] = p
			for _, m := range p.Materials {
				fillNeeded(m)
			}
		}

	}
	fillNeeded(p)
	index := map[*model.Product][]*model.Blueprint{}
	bps, err := c.client.GetCorpBlueprints()
	if err != nil {
		c.logger.Warnf("unable to get corporation blueprints: %s", err.Error())
		return nil, nil
	}
	for need, prod := range needed {
		for _, bp := range bps {
			if bp.Kind != model.BlueprintCopy {
				continue
			}
			if int(bp.TypeID) == need {
				index[prod] = append(index[prod], bp)
			}
		}
	}

	var missing []*model.Product
	for _, prod := range needed {
		if _, ok := index[prod]; !ok {
			missing = append(missing, prod)
		}
	}
	return index, missing
}

// printProductInfo prints production chain details.
func (c ProductCommand) printProductInfo(p *model.Product) {
	batchSize := decimal.NewFromFloat(float64(p.BatchSize))
	costEach := p.Cost().Mul(batchSize) // Cost has quantity baked in.
	bp, err := c.client.GetMaterialSheet(p.TypeID)
	if err != nil {
		fmt.Println("Unable to print production chain detail:", err.Error())
		return
	}
	batchQuantity := decimal.NewFromFloat(float64(bp.ProducesQty)).Mul(batchSize)
	sellEach := p.MarketPrice.Mul(batchQuantity)
	profitEach := sellEach.Sub(costEach)
	marginEach := decimal.Zero
	if sellEach.Cmp(decimal.Zero) != 0 {
		marginEach = profitEach.Div(sellEach).Mul(decimal.NewFromFloat(100))
	}
	unitLabel := "unit"
	if batchQuantity.GreaterThan(decimal.NewFromFloat(1)) {
		unitLabel = fmt.Sprintf("%s units", batchQuantity)
	}
	hasBuildComponent := false
	regionName := c.getRegionName(p.MarketRegionID)
	fmt.Println(text.Boldf(text.CenterText(c.getProductName(p), text.StandardTerminalWidthInChars)))
	fmt.Println(text.CenterText(regionName, text.StandardTerminalWidthInChars))
	fmt.Println()
	if len(p.Materials) == 0 {
		fmt.Println("This component cannot be manufactured. It must be procured by other means.")
		fmt.Println()
		fmt.Println("Best sell price in "+text.Boldf(regionName)+":", text.PadCurrencyLeft(p.MarketPrice, 15)+"/ea")
		fmt.Println()
		return
	}
	fmt.Printf(
		" #  %s%s%s%s\n",
		text.PadTextRight("Material Name", 28),
		text.PadTextLeft("Cost/ea", 17),
		text.PadTextLeft("Qty Req", 12),
		text.PadTextLeft("Cost/"+unitLabel, 19))
	index := new(int)
	for _, part := range p.Materials {
		if part.Kind == model.ProductBuild {
			hasBuildComponent = true
		}
		c.printChildProductInfo(part, batchSize, p.MaterialEfficiency, index, 0)
	}
	fmt.Println()
	fmt.Printf("%s%s%s\n", text.PadTextLeft(fmt.Sprintf("Per %s", unitLabel), 49), text.PadTextLeft("Revenue", 12), text.PadCurrencyLeft(sellEach, 19))
	fmt.Printf("%s%s%s\n", text.PadTextLeft(fmt.Sprintf("%s%% ME", p.MaterialEfficiency.Mul(decimal.NewFromFloat(100)).StringFixed(0)), 49), text.PadTextLeft("Cost", 12), text.PadCurrencyLeft(costEach, 19))
	fmt.Printf("%s%s\n", text.PadTextLeft("Profit", 61), text.PadCurrencyLeft(profitEach, 19))
	fmt.Printf("%s%s\n", text.PadTextLeft("Margin", 61), "      %"+text.PadTextLeft(marginEach.StringFixed(2), 12))

	if hasBuildComponent {
		fmt.Println()
		fmt.Println("* 'M' indicates the component will be produced in-house.")
	}
	fmt.Println()
}

// printChildProductInfo displays a single component's details.
//
// This function calls itself recursively to traverse the entire production
// chain.
func (c ProductCommand) printChildProductInfo(p *model.Product, parentBatchSize decimal.Decimal, parentME decimal.Decimal, index *int, indent int) {
	*index += 1
	costEach := p.Cost()
	qtyAfterME := decimal.NewFromFloat(float64(p.Quantity)).Mul(parentBatchSize).
		Div(decimal.NewFromFloat(1).Add(parentME)).Round(0)
	costTotal := p.Cost().Mul(qtyAfterME)

	var kind string
	if p.Kind == model.ProductBuild {
		kind = "M"
	}
	fmt.Printf(
		"%s %s%s%s%s%s\n",
		text.PadTextLeft(strconv.Itoa(*index), 3),
		text.PadTextRight(strings.Repeat("  ", indent)+c.getProductName(p), 28),
		text.PadTextLeft(kind, 2),
		text.PadCurrencyLeft(costEach, 15),
		text.PadIntegerLeft(int(qtyAfterME.IntPart()), 12),
		text.PadCurrencyLeft(costTotal, 19))
	return
	indent += 1
	if p.Kind == model.ProductBuild {
		for _, part := range p.Materials {
			c.printChildProductInfo(part, parentBatchSize, p.MaterialEfficiency, index, indent)
		}
	}
}

// efficiencyValue is a Material or Time Efficiency value.
type efficiencyValue struct {
	best     int
	worst    int
	nonEmpty bool // If false, any value is worse.
}

func (e *efficiencyValue) String() string {
	worst := fmt.Sprintf("%d", e.worst)
	if len(worst) < 2 {
		worst = " " + worst
	}
	return fmt.Sprintf("%d/%s", e.best, worst)
}

func (e *efficiencyValue) track(val int) {
	if e.nonEmpty || val < e.worst {
		e.worst = val
	}
	if val > e.best {
		e.best = val
	}
	e.nonEmpty = true
}

func (c ProductCommand) printBlueprintOverview(p *model.Product) {
	if c.corpID == 0 {
		fmt.Println()
		fmt.Println("Unable to load materials inventory.")
		fmt.Println("Start the program with valid credentials to enable this feature.")
		fmt.Println()
		return
	}
	bpIndex, missing := c.getBlueprintIndex(p)

	fmt.Println()
	fmt.Println(text.Boldf(text.CenterText("Materials Inventory", text.StandardTerminalWidthInChars)))
	fmt.Println()

	if len(missing) > 0 {
		fmt.Println("Missing Blueprints")
		fmt.Println()
		for _, prod := range missing {
			fmt.Printf("%s %s\n", text.PadTextLeft(fmt.Sprintf("%d", prod.TypeID), 9), c.getProductName(prod))
		}
	}

	fmt.Println("Available Blueprints")
	col1Width := 44
	col2Width := 9
	col3Width := 9
	col4Width := 9
	col5Width := 9
	fmt.Printf(
		"%s%s\n",
		text.PadTextLeft("Best/Worst ", col1Width+col2Width+col3Width+col4Width),
		text.PadTextLeft("Total", col5Width))
	fmt.Printf(
		"%s%s%s%s%s\n",
		text.PadTextRight("Name", col1Width),
		text.PadTextLeft("Type ID", col2Width),
		text.PadTextLeft("ME%", col3Width),
		text.PadTextLeft("TE%", col4Width),
		text.PadTextLeft("Runs", col5Width))

	for prod, bps := range bpIndex {
		if len(bps) == 0 {
			c.logger.Debug("expected to have at least 1 blueprint in the result")
			continue
		}
		totalRuns := 0
		matEff := &efficiencyValue{}
		timeEff := &efficiencyValue{}
		for _, bp := range bps {
			totalRuns += int(bp.Runs)
			matEff.track(bp.MaterialEfficiency)
			timeEff.track(bp.TimeEfficiency)
		}
		fmt.Printf(
			"%s%s%s%s%s\n",
			text.PadTextRight(c.getProductName(prod), col1Width),
			text.PadTextLeft(fmt.Sprintf("%d", prod.TypeID), col2Width),
			text.PadTextLeft(matEff.String(), col3Width),
			text.PadTextLeft(timeEff.String(), col4Width),
			text.PadTextLeft(fmt.Sprintf("%d", totalRuns), col5Width))
	}
	fmt.Println()
}

func (c ProductCommand) getProductLineIndex(p *model.Product) map[int]*model.Product {
	index := map[int]*model.Product{}
	curr := 0
	index[curr] = p // Root is 0.
	for _, part := range p.Materials {
		curr += 1
		index[curr] = part
	}
	return index
}

// newProduct creates a new production chain and opens it in the editor.
func (c ProductCommand) newProduct(args ...string) {
	if p := c.previewProduct(args...); p != nil {
		c.newProductEditor(p).Loop()
	}
}

// editProduct loads a production chain and opens it in the editor.
func (c ProductCommand) editProduct(args ...string) {
	productID := 0
	var ok bool
	var err error
	if len(args) > 0 {
		productID, err = strconv.Atoi(args[0])
	}
	if err != nil || productID <= 0 {
		productID, ok = c.env.PromptInt("Specify Product ID", nil, validateIntGreaterThan(0))
		if !ok {
			return
		}
	}
	product, err := c.client.GetProduct(productID)
	if err != nil {
		c.logger.Debugf("unable to load production chain: %s", err.Error())
		fmt.Println("Error loading production chain from db, try again.")
		return
	}
	c.printProductInfo(product)
	c.newProductEditor(product).Loop()
}

// previewProduct displays a default view for a given typeID.
func (c ProductCommand) previewProduct(args ...string) *model.Product {
	item, ok := c.env.PromptItemTypeDetail("Specify Item Type", strings.Join(args, " "))
	if !ok {
		return nil
	}
	product, err := c.client.NewProduct(item.ID)
	if err != nil {
		c.logger.Warnf("unable to create product: %s", err.Error())
		fmt.Println("Error creating production chain, try again.")
		return nil
	}
	for _, mat := range product.Materials {
		if len(mat.Materials) > 0 {
			mat.Kind = model.ProductBuild
		}
	}
	if _, err = c.client.UpdateProductPrices(product); err != nil {
		c.logger.Warnf("unable to populate production chain prices: %s", err.Error())
		fmt.Println("Error loading production chain market prices, prices may be incorrect.")
	}
	c.printProductInfo(product)
	return product
}

// showProduct loads and displays a production chain's details.
func (c ProductCommand) showProduct(args ...string) {
	productID := 0
	var ok bool
	var err error
	if len(args) > 0 {
		productID, err = strconv.Atoi(args[0])
	}
	if err != nil || productID <= 0 {
		productID, ok = c.env.PromptInt("Specify Product ID", nil, validateIntGreaterThan(0))
		if !ok {
			return
		}
	}
	product, err := c.client.GetProduct(productID)
	if err != nil {
		c.logger.Debugf("unable to load product: %s", err.Error())
		fmt.Println("Error loading production chain from db, try again.")
		return
	}
	c.printProductInfo(product)
}

// listProducts lists all the production chains.
func (c ProductCommand) listProducts() {
	products, err := c.client.GetProducts()
	if err != nil {
		c.logger.Debugf("unable to fetch production chain: %s", err.Error())
		fmt.Println("Error loading production chain from db, try again.")
		return
	}
	fmt.Println("Listing", len(products), "production chains.")
	fmt.Println()
	if len(products) == 0 {
		fmt.Println("There are no production chains. Create a new production chain with")
		fmt.Println("  product add")
		return
	}
	fmt.Printf(
		"%s%s%sType ID\n",
		text.PadTextRight("ID", 12),
		text.PadTextRight("Region", 12),
		text.PadTextRight("Name", 42))
	for _, prod := range products {
		fmt.Printf(
			"%-12.f%s%s%d\n",
			float64(prod.ProductID),
			text.PadTextRight(c.getRegionName(prod.MarketRegionID), 12),
			text.PadTextRight(c.getProductName(prod), 42),
			prod.TypeID)
	}
}
