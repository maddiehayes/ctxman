package commands

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

var UseCmd = &cli.Command{
	Name:  "use",
	Usage: "switch to the given context",
	Action: func(c *cli.Context) error {
		// If --current passed in, print the current context and exit
		numArgs := c.Args().Len()
		if numArgs == 0 {
			log.Fatal("no context provided")
		} else if numArgs > 1 {
			log.Fatal("too many args provided")
		}
		ctx := c.Args().Get(0)
		fmt.Printf("switched to context %s\n", ctx)
		return nil
	},
	ArgsUsage: "CONTEXT",
}
