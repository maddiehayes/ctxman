package commands

import (
	"fmt"
	"strings"

	"github.com/maddiehayes/ctxman/internal/config"
	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

var ClearCmd = &cli.Command{
	Name:  "clear",
	Usage: "clear the current context",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "pretty",
			Usage: "Print one unset statment per line for human readability",
		},
	},
	Action: func(c *cli.Context) error {
		varNames := config.FromAppContext(c).VariableNames
		if len(varNames) == 0 {
			log.Warn("no varNames specified in config file")
			return nil
		}
		unsetCmd := buildUnsetCmd(varNames, false)
		fmt.Println(unsetCmd)
		return nil
	},
	ArgsUsage: "CONTEXT",
}

// unsetCmd generates a string that can be used to unset all context-related
// environment variables, optionally printing one statement per line.
func buildUnsetCmd(varNames []string, pretty bool) string {
	cmd := "unset"
	builder := strings.Builder{}
	if pretty {
		// Write one statement per line
		for _, name := range varNames {
			builder.WriteString(fmt.Sprintf("%s %s\n", cmd, name))
		}
	} else {
		// Write all statement on one line
		builder.WriteString("unset")
		for _, name := range varNames {
			builder.WriteString(fmt.Sprintf(" %s", name))
		}
	}
	return builder.String()
}
