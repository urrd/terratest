package terraform

import (
	"testing"
)

// Destroy runs terraform destroy with the given options and return stdout/stderr.
func Taint(t *testing.T, options *Options, module, name string) string {
	out, err := TaintE(t, options, module, name)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

// DestroyE runs terraform destroy with the given options and return stdout/stderr.
func TaintE(t *testing.T, options *Options, module, name string) (string, error) {
	if module == "" {
		return RunTerraformCommandE(t, options, FormatArgs(options.Vars, "taint", name)...)
	} else {
		return RunTerraformCommandE(t, options, FormatArgs(options.Vars, "taint", "-module="+module, name)...)
	}
}

