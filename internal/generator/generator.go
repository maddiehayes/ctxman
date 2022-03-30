package generator

import (
	"encoding/json"
)

const (
	generatorExec string = "exec"
)

type Generatable interface {
	getKind() string
	generate() []string
	validate() error
}

type Generator struct {
	kind  string `yaml:"kind"`
	Value interface{}
}

// TODO: override Viper unmarshall hook for Generator objects. Viper uses github.com/mitchellh/mapstructure
// under the hood for unmarshaling values which uses mapstructure tags by default.
//
// <https://github.com/spf13/viper#unmarshaling>
// <https://sagikazarmark.hu/blog/decoding-custom-formats-with-viper/>

// UnmarshalJSON overrides the unmarshal interface for the generator type and acts
// as a Generator factory, creating a generator based on the `kind` attribute.
// func (g *Generator) UnmarshalJSON(data []byte) error {
// 	log.Debug("unmarshalling...")
// 	// Extract the generator kind to know which struct to create
// 	kind, err := getKind(data)
// 	if err != nil {
// 		log.Error("failed to unmarshal generator kind")
// 		return err
// 	}
// 	// Create new generator based on the kind
// 	switch kind {
// 	case generatorExec:
// 		g.Value = ExecGenerator{}
// 	default:
// 		return fmt.Errorf("generator kind '%s' is unsupported", kind)
// 	}
// 	return nil
// }

// getKind extracts the Generator kind from unmarshalled generator object
func getKind(data []byte) (string, error) {
	// Get the kind out of the unstructured data
	var kindGetter struct {
		kind string `yaml:"kind"`
	}
	if err := json.Unmarshal(data, &kindGetter); err != nil {
		return "", err
	}
	return kindGetter.kind, nil
}

type ExecParams struct {
	command []string `yaml:"command"`
}

type ExecGenerator struct {
	kind       string     `yaml:"kind"`
	parameters ExecParams `yaml:"parameters"`
}

func UnmarshalExecGenerator(data []byte) *ExecGenerator {

	return &ExecGenerator{}
}

func (g *ExecGenerator) getKind() string    { return g.kind }
func (g *ExecGenerator) generate() []string { return nil }
func (g *ExecGenerator) validate() error    { return nil }
