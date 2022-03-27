package context

import (
	"os"
	"testing"
)

func TestGetCurrentContext(t *testing.T) {
	// Return value of EnvVarName when set
	expected := "test-context"
	os.Setenv(EnvVarName, expected)
	actual := Current()
	if expected != *actual {
		t.Errorf("Current() = %s; want %s", expected, *actual)
	}

	// Return nil pointer when unset
	expected = ""
	os.Setenv(EnvVarName, expected)
	actual = Current()
	if nil != actual {
		t.Errorf("Current() = %s; want %s", expected, *actual)
	}
}
