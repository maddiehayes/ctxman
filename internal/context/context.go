package context

import (
	"os"
)

const (
	// EnvVarName environment variable used to store current context name
	EnvVarName string = "CTX"
)

// Context represents the state of the user environment within the current session
type Context struct {
	Name *string `yaml:"name"`
}

// Contexts a collection of Context structs
type Contexts []*Context

// Current returns the name of the current context
func Current() *string {
	ctx := os.Getenv("CTX")
	if ctx == "" {
		return nil
	}
	return &ctx
}
