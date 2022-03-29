// TODO: add validator package
package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/maddiehayes/ctxman/internal/context"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const (
	CliContextKey string = "config"
)

func DefaultFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(homeDir, ".config", "ctxman", "config.yaml")
}

type Config struct {
	ApiVersion    *string          `yaml:"apiVersion"` // validate:"required"
	Contexts      context.Contexts `yaml:"contexts"`
	VariableNames []string         `yaml:"variableNames"`
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
	log.Debugf("contexts: %s", c.Contexts.Names())
	log.Debugf("variable names: %s", c.VariableNames)
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
