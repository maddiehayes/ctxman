// TODO: add validator package
package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/maddiehayes/ctxman/internal/context"
	"github.com/maddiehayes/ctxman/internal/scope"

	// "github.com/maddiehayes/ctxman/internal/scope"
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
	Scopes        scope.Scopes     `yaml:"scopes"`
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
	log.Debugf("variables: %s", c.VariableNames)
	log.Debugf("Scope.Name: %s", *c.Scopes[0].Description)
	log.Debugf("Scope.Description: %s", *c.Scopes[0].Name)
	log.Debugf("Scope.ParameterName: %s", *c.Scopes[0].Parameter.Name)
	log.Debugf("Scope.ParameterGenerator: %+v", *&c.Scopes[0].Parameter.Generator.Value)
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
