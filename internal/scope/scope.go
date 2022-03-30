package scope

import (
	"github.com/maddiehayes/ctxman/internal/generator"
)

type ConfigParam struct {
	Name      *string             `yaml:"name"`
	Generator generator.Generator `yaml:"generator"`
}

type Scope struct {
	Name        *string      `yaml:"name"`
	Description *string      `yaml:"description"`
	Parameter   *ConfigParam `yaml:"parameter"`
}

type Scopes []*Scope
