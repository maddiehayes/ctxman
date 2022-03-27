// TODO: add validator package
package config

import (
	"errors"
	"fmt"

	"github.com/maddiehayes/ctxman/internal/context"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const (
	CliContextKey string = "config"
)

type Config struct {
	ApiVersion *string          `yaml:"apiVersion"` // validate:"required"
	Contexts   context.Contexts `yaml:"contexts"`
}

// Validate performs validation on the app config
func (c *Config) Validate() error {
	if c.ApiVersion == nil {
		return errors.New("missing apiVersion field")
	}
	if *c.ApiVersion != "v0.0.1-alpha" {
		return fmt.Errorf("unknown api version: %s", *c.ApiVersion)
	}
	log.Debugf("current version: %s", *c.ApiVersion)
	return nil
}

// FromAppContext extracts the Config from the cli app's context
func FromAppContext(c *cli.Context) *Config {
	value := c.Context.Value(CliContextKey)
	cfg, ok := value.(*Config)
	if !ok {
		log.Fatalf("failed to load config from cli context")
	}
	return cfg
}

// GetContext returns a Context matching the given name if it exists in the config
func (c *Config) GetContext(name string) (*context.Context, error) {
	// Return context with given name if it exists in the config
	for _, context := range c.Contexts {
		if *context.Name == name {
			return context, nil
		}
	}
	// Return error if DNE
	return nil, errors.New(fmt.Sprintf("context %s does not exist", name))
}
