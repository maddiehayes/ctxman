package list

import (
	"github.com/urfave/cli/v2"
)

var ListCmd = &cli.Command{
	Name:  "list",
	Usage: "list objects",
	Action: func(c *cli.Context) error {
		cli.ShowAppHelp(c)
		return nil
	},
	Subcommands: []*cli.Command{
		ListContextsCmd,
	},
}
