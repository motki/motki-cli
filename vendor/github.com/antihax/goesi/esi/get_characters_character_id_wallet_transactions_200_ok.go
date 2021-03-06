/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.7.6
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

import (
	"time"
)

/* A list of GetCharactersCharacterIdWalletTransactions200Ok. */
//easyjson:json
type GetCharactersCharacterIdWalletTransactions200OkList []GetCharactersCharacterIdWalletTransactions200Ok

/* wallet transaction */
//easyjson:json
type GetCharactersCharacterIdWalletTransactions200Ok struct {
	TransactionId int64     `json:"transaction_id,omitempty"` /* Unique transaction ID */
	Date          time.Time `json:"date,omitempty"`           /* Date and time of transaction */
	TypeId        int32     `json:"type_id,omitempty"`        /* type_id integer */
	LocationId    int64     `json:"location_id,omitempty"`    /* location_id integer */
	UnitPrice     float64   `json:"unit_price,omitempty"`     /* Amount paid per unit */
	Quantity      int32     `json:"quantity,omitempty"`       /* quantity integer */
	ClientId      int32     `json:"client_id,omitempty"`      /* client_id integer */
	IsBuy         bool      `json:"is_buy,omitempty"`         /* is_buy boolean */
	IsPersonal    bool      `json:"is_personal,omitempty"`    /* is_personal boolean */
	JournalRefId  int64     `json:"journal_ref_id,omitempty"` /* journal_ref_id integer */
}
