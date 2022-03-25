package context

import (
	"os"
)

// type Context struct {
// 	Name string
// }

func Current() string {
	ctx := os.Getenv("CTX")
	if ctx == "" {
		return "no current context"
	}
	return ctx
}
