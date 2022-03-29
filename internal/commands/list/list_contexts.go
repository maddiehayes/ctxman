package list

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/maddiehayes/ctxman/internal/config"
	"github.com/maddiehayes/ctxman/internal/context"
	"github.com/maddiehayes/ctxman/internal/printers"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var ListContextsCmd = &cli.Command{
	Name:  "contexts",
	Usage: "list existing contexts",
	Before: func(c *cli.Context) error {
		// Validate user input
		numArgs := c.Args().Len()
		if numArgs != 0 {
			log.Fatal(errors.New(fmt.Sprintf("incorrect number of args provided: got %d, expected %d", numArgs, 0)))
		}
		return nil
	},
	Action: func(c *cli.Context) error {
		// Load contexts from config
		contexts := config.FromAppContext(c).Contexts

		// Exit if no saved contexts
		if len(contexts) == 0 {
			fmt.Println("no saved contexts")
			return nil
		}
		contexts.SortByName()

		// Just print one context name per line if only names requested
		if c.Bool("names-only") {
			fmt.Println(strings.Join(contexts.Names()[:], "\n"))
			return nil
		}

		// Otherwise, generate table contents
		currentCtx := context.Current()
		out := printers.NewTabWriter(os.Stdout)
		defer out.Flush()

		err := printContextHeaders(out)
		if err != nil {
			log.Fatal(err)
		}
		for _, ctx := range config.FromAppContext(c).Contexts {
			if err := printContext(ctx, out, currentCtx); err != nil {
				log.Fatal(err)
			}
		}
		return nil
	},
	ArgsUsage: "CONTEXT",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "names-only",
			Usage: "List only context names",
		},
	},
}

// printContextHeaders prints headers for the list contexts table
func printContextHeaders(out io.Writer) error {
	columnNames := []string{"CURRENT", "NAME"}
	_, err := fmt.Fprintf(out, "%s\n", strings.Join(columnNames, "\t"))
	return err
}

// printContext prints a single row of list contexts table
func printContext(ctx *context.Context, w io.Writer, currentCtx *string) error {
	prefix := " "
	if currentCtx != nil && *currentCtx == *ctx.Name {
		prefix = "*"
	}
	_, err := fmt.Fprintf(w, "%s\t%s\n", prefix, *ctx.Name)
	return err
}
