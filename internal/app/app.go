package app

import (
	gocontext "context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/maddiehayes/ctxman/internal/commands"
	"github.com/maddiehayes/ctxman/internal/commands/list"
	"github.com/maddiehayes/ctxman/internal/config"
	"github.com/maddiehayes/ctxman/internal/context"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

var (
	cfg *config.Config
)

var app *cli.App = &cli.App{
	Name:  "ctx",
	Usage: "Manage your shell environment context",
	Before: func(c *cli.Context) error {
		// Enable debug logs
		if c.Bool("debug") {
			log.SetLevel(log.DebugLevel)
		}
		// Load config file
		cfgFile := c.String("config")
		log.Debugf("using config file: %q\n", cfgFile)
		viper.SetConfigFile(cfgFile)
		cfg = &config.Config{}
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
		if err := viper.Unmarshal(cfg); err != nil {
			return err
		}
		if err := cfg.Validate(); err != nil {
			return err
		}
		// Add Config to context
		c.Context = gocontext.WithValue(c.Context, config.CliContextKey, cfg)
		return nil
	},
	Action: func(c *cli.Context) error {
		// If --current, print the current context and exit
		if c.Bool("current") {
			ctx := context.Current()
			if ctx != nil {
				fmt.Println(*ctx)
			} else {
				log.Warn("no current context")
			}
			return nil
		}
		// Else, use fzf to select from existing contexts
		// TODO: fzf select context
		log.Debug("auto-complete to set current context")
		log.Warn("`ctx` fuzzy finder is unimplemented")
		return nil
	},
	Commands: []*cli.Command{
		commands.UseCmd,
		list.ListCmd,
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:      "config",
			Aliases:   []string{"f"},
			Usage:     "load configuration from `FILE`",
			Value:     defaultCfgFile(),
			TakesFile: true,
			Hidden:    true,
		},
		&cli.BoolFlag{
			Name:    "current",
			Aliases: []string{"c"},
			Usage:   "Print name of current context",
		},
		&cli.BoolFlag{
			Name:  "debug",
			Usage: "Enable debug logs",
		},
	},
	// ExitErrHandler handles errors returned from any cli action
	ExitErrHandler: func(context *cli.Context, err error) {
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Run() error {
	return app.Run(os.Args)
}

// defaultCfgFile loads the user home directory
func defaultCfgFile() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(homeDir, ".config", "ctxman", "config.yaml")
}
