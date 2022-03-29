package commands

import (
	"errors"
	"fmt"

	"github.com/maddiehayes/ctxman/internal/config"

	"github.com/urfave/cli/v2"
)

var UseCmd = &cli.Command{
	Name:  "use",
	Usage: "switch to the given context",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "pretty",
			Usage: "Print one export statment per line for human readability",
		},
	},
	Before: func(c *cli.Context) error {
		// Validate user input
		numArgs := c.Args().Len()
		if numArgs != 1 {
			return errors.New("incorrect number of args provided")
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		// Load existing context from config
		ctxName := c.Args().Get(0)
		pretty := c.Bool("pretty")
		ctx, err := config.FromAppContext(c).GetContext(ctxName)
		if err != nil {
			return err
		}
		// Print environment exports for the selected context
		fmt.Println(ctx.GetEnvExports(pretty))
		return nil
	},
	ArgsUsage: "CONTEXT",
}
