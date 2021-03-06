// Command motki is a utility for interacting with a remote MOTKI application
// server.
//
//  Usage of motki:
//    -credentials string
//      	Username and password separated by a colon. (ie. "frank:mypass")
//    -history-file string
//      	Path to the CLI history file. (default ".history")
//    -insecure-skip-verify
//      	INSECURE: Skip verification of server SSL cert.
//    -log-level string
//      	Log level. Possible values: debug, info, warn, error. (default "warn")
//    -server string
//      	Backend server host and port. (default "motki.org:18443")
//    -version
//      	Display the application version.
//
// By default, the motki command connects to the public MOTKI server available
// at https://motki.org. Override this functionality by passing the -server flag on
// the command line.
//
//   motki -server myserverhost:8443
//
// To start an authenticated session, you must pass the -credentials flag on the
// command line.
//
//   motki -credentials frank:mypass
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/motki/cli/app"

	"github.com/motki/core/app/profile"

	"github.com/motki/core/log"
	"github.com/motki/core/proto"
	"github.com/motki/core/proto/client"
)

var serverAddr = flag.String("server", "motki.org:18443", "Backend server host and port.")
var historyPath = flag.String("history-file", ".history", "Path to the CLI history file.")
var credentials = flag.String("credentials", "", "Username and password separated by a colon. (ie. \"frank:mypass\")")
var logLevel = flag.String("log-level", "warn", "Log level. Possible values: debug, info, warn, error.")
var insecureSkipVerify = flag.Bool("insecure-skip-verify", false, "INSECURE: Skip verification of server SSL cert.")
var version = flag.Bool("version", false, "Display the application version.")

// fatalf creates a default logger, writes the given message, and exits.
func fatalf(format string, vals ...interface{}) {
	logger := log.New(log.Config{})
	logger.Fatalf(format, vals...)
}

var Version = "dev"

func main() {
	flag.Parse()
	if *version {
		fmt.Printf("%s %s. %s\n", os.Args[0], Version, "https://github.com/motki/cli")
		os.Exit(0)
	}

	// Start the profiler, if enabled via command-line flags.
	pr := profile.New()

	// Trimmed down motki app configuration.
	appConf := &app.Config{
		Backend: proto.Config{
			Kind: proto.BackendRemoteGRPC,
			RemoteGRPC: proto.RemoteConfig{
				ServerAddr:         *serverAddr,
				InsecureSkipVerify: *insecureSkipVerify,
			},
		},
		Logging: log.Config{
			Level: *logLevel,
		},
	}

	// Writing to stderr offers a way to redirect the logger output to a file instead of
	// interrupting the user.
	appConf.Logging.OutputType = log.OutputStderr

	conf := app.NewCLIConfig(appConf)

	// Add credentials to the configuration, if they were specified.
	if *credentials != "" {
		parts := strings.Split(*credentials, ":")
		if len(parts) != 2 {
			fatalf("motki: invalid credentials, expected format \"username:password\"")
		}
		conf = conf.WithCredentials(parts[0], parts[1])
	} else {
		user, pass := os.Getenv("MOTKI_USERNAME"), os.Getenv("MOTKI_PASSWORD")
		if user != "" && pass != "" {
			conf = conf.WithCredentials(user, pass)
		}
	}

	// Initialize the CLI environment.
	env, err := app.NewCLIEnv(conf, *historyPath)
	if err != nil {
		if err == client.ErrBadCredentials {
			fmt.Println("Invalid username or password.")
		}
		fatalf("motki: error initializing application environment: %s", err.Error())
	}

	// Start interactive command loop in another goroutine.
	go env.LoopCLI()

	// Block the main loop until SIGINT, SIGHUP, or SIGKILL is received,
	// or until CTRL+C is pressed from the command-line.
	err = env.BlockUntilSignal(nil)

	// Stop the profiler, if enabled.
	if pr != nil {
		pr.Stop()
	}

	if err != nil {
		env.Logger.Warnf("motki: failed to shutdown cleanly: %s", err.Error())
		os.Exit(1)
	}
}
