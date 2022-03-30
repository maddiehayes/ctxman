package generator

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {
	// Return value of EnvVarName when set
	expected := Generator{
		kind: generatorExec,
		Value: ExecGenerator{
			kind: generatorExec,
		},
	}
	blob := `
kind: exec
parameters:
	command: ["yq","'.workstreams[].name'",  "/Users/maddieh/Documents/Personal/ctxman/scopes.yaml"]`

	actual := Generator{}
	if err := json.Unmarshal([]byte(blob), &actual); err != nil {
		t.Errorf("unmarshall failed")
	}
	if expected != actual {
		t.Errorf("Current() = %s; want %s", expected, actual)
	}
}
