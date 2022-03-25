package app

import (
	"fmt"
	"os"
	"path/filepath"

	// "github.com/maddiehayes/ctxman/internal/commands"
	// "github.com/spf13/viper"
	"github.com/maddiehayes/ctxman/internal/commands"
	"github.com/maddiehayes/ctxman/internal/config"
	"github.com/maddiehayes/ctxman/internal/context"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

var (
	app *cli.App
	cfg *config.Config
)

func Run() error {
	app = &cli.App{
		Name:  "ctx",
		Usage: "Manage your shell environment context",
		Before: func(c *cli.Context) error {
			cfgFile := c.String("config")
			log.Debugf("using config file: %q\n", cfgFile)
			viper.SetConfigFile(cfgFile)
			cfg = &config.Config{}
			checkErr(viper.ReadInConfig())
			checkErr(viper.Unmarshal(cfg))
			return cfg.Validate()
		},
		Action: func(c *cli.Context) error {
			// If --current passed in, print the current context and exit
			if c.Bool("current") {
				fmt.Println(context.Current())
				return nil
			}
			// Else, autofill context
			return nil
		},
		Commands: []*cli.Command{
			commands.UseCmd,
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
		},
	}

	return app.Run(os.Args)
}

// defaultCfgFile loads the user home directory
func defaultCfgFile() string {
	homeDir, err := os.UserHomeDir()
	checkErr(err)
	return filepath.Join(homeDir, ".config", "ctxman", "config.yaml")
}

// checkErr if an error is returned, logs the error and exists with error code
func checkErr(msg interface{}) {
	if msg != nil {
		log.Fatal(msg)
	}
}
