package app

import (
	"bytes"
	gocontext "context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/maddiehayes/ctxman/internal/commands"
	"github.com/maddiehayes/ctxman/internal/commands/list"
	"github.com/maddiehayes/ctxman/internal/config"
	"github.com/maddiehayes/ctxman/internal/context"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

var (
	cfg *config.Config
)

var app *cli.App = &cli.App{
	Name:   "ctx",
	Usage:  "Manage your shell environment context",
	Before: before,
	Action: action,
	Commands: []*cli.Command{
		commands.UseCmd,
		commands.ClearCmd,
		list.ListCmd,
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:      "config",
			Aliases:   []string{"f"},
			Usage:     "load configuration from `FILE`",
			Value:     config.DefaultFilePath(),
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
			log.Error(err)
		}
	},
}

func Run() error {
	return app.Run(os.Args)
}

func before(c *cli.Context) error {
	// Enable debug logs
	if c.Bool("debug") {
		log.SetLevel(log.DebugLevel)
	}
	// Load config file
	cfgFile := c.String("config")
	log.Debugf("using config file: %s\n", cfgFile)
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
}

func action(c *cli.Context) error {
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
	numArgs := c.Args().Len()
	if numArgs != 0 {
		return errors.New(fmt.Sprintf("wrong number of args: got %d, expected %d", numArgs, 0))
	}
	// Else, use fzf to select from existing contexts
	// TODO: fzf select context
	log.Debug("auto-complete to set current context")
	// log.Warn("`ctx` fuzzy finder is unimplemented")
	contextNames := listContextNames(c)
	if contextNames == nil {
		fmt.Println("no saved contexts")
	}
	interactiveSelect(config.FromAppContext(c), os.Stderr, *contextNames)
	return nil
}

func listContextNames(c *cli.Context) *string {
	// Load contexts from config
	contexts := config.FromAppContext(c).Contexts

	// Exit if no saved contexts
	if len(contexts) == 0 {
		return nil
	}
	contexts.SortByName()

	// Return string pointer of one name per line
	names := strings.Join(contexts.Names()[:], "\n")
	return &names
}

func interactiveSelect(config *config.Config, stderr io.Writer, contextNames string) error {
	cmd := exec.Command("fzf", "--ansi", "--no-preview", "--prompt", "select context: ", "--no-multi")
	var out bytes.Buffer
	cmd.Stdin = os.Stdin
	cmd.Stderr = stderr
	cmd.Stdout = &out

	cmd.Env = append(
		os.Environ(),
		fmt.Sprintf("FZF_DEFAULT_COMMAND=echo \"%s\"", contextNames),
		// fmt.Sprintf("%s=1", env.EnvForceColor),
	)
	if err := cmd.Run(); err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			return err
		}
	}
	choice := strings.TrimSpace(out.String())
	if choice == "" {
		return errors.New("you did not choose any of the options")
	}
	log.Infof("selected context %s", choice)
	ctx, _ := cfg.GetContext(choice)
	fmt.Println(ctx.GetEnvExports(false))
	return nil
}
