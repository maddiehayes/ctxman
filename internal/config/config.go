// TODO: add validator package
package config

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	ApiVersion *string `yaml:"apiVersion"` // validate:"required"
}

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
