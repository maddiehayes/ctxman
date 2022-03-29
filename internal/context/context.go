package context

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	// EnvVarName is the environment variable used to store current context name.
	EnvVarName string = "CTX"
)

// Context represents the state of the user environment within the current session.
type Context struct {
	Name        *string           `yaml:"name"`
	Environment map[string]string `yaml:"environment"`
}

// Contexts is a collection of Context structs.
type Contexts []*Context

// ContextSorter is used for sorting Contexts by Name attribute.
type NameSorter Contexts

func (c NameSorter) Len() int           { return len(c) }
func (c NameSorter) Swap(i, j int)      { *c[i], *c[j] = *c[j], *c[i] }
func (c NameSorter) Less(i, j int) bool { return *c[i].Name < *c[j].Name }

// SortByName sorts a Contexts collection by the Name attribute.
func (c Contexts) SortByName() {
	sort.Sort(NameSorter(c))
}

func (c Contexts) Names() []string {
	names := make([]string, len(c))
	for idx, ctx := range c {
		names[idx] = *ctx.Name
	}
	return names
}

// Current returns the name of the current context based on the current value
// of the .environment variable
func Current() *string {
	ctx := os.Getenv("CTX")
	if ctx == "" {
		return nil
	}
	return &ctx
}

// EnvExports generates a string that can be used to set all environment variables
// belonging to a specific contexts, optionally printing one export per line.
func (c *Context) GetEnvExports(pretty bool) string {
	builder := strings.Builder{}
	if pretty {
		// Write one export per line
		for name, value := range c.Environment {
			builder.WriteString(fmt.Sprintf("export %s=%s\n", name, value))
		}
	} else {
		// Write all exports on one line
		builder.WriteString("export")
		for name, value := range c.Environment {
			builder.WriteString(fmt.Sprintf(" %s=%s", name, value))
		}
	}
	return builder.String()
}
