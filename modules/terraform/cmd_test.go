package terraform

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetCommonOptionsDefaultWorkingDir(t *testing.T) {
	t.Parallel()

	expected := "/test/terraform/dir"
	additionalOptions := &Options{ TerraformDir: expected }
	options, _ := GetCommonOptions(additionalOptions)
	
	assert.Equal(t, expected, options.WorkingDir, "Options WorkingDir must default to TerraformDir")
}

func TestGetCommonOptionsDontAppendTerraformDir(t *testing.T) {
	t.Parallel()

	expected := "/test/terraform/dir"

	additionalOptions := &Options{ TerraformDir: expected, WorkingDir: expected }
	_, args := GetCommonOptions(additionalOptions, "init")
	assert.NotEqual(t, expected, args[len(args)-1], "Dont append TerraformDir to args when it's the same as WorkingDir")

	additionalOptions = &Options{ TerraformDir: expected, WorkingDir: "/some/other/dir" }

	_, args = GetCommonOptions(additionalOptions)
	if len(args) > 0 {
		assert.NotEqual(t, expected, args[len(args)-1], "Don't append TerraformDir to args when not init, get, plan, apply, destroy")
	}

	_, args = GetCommonOptions(additionalOptions, "other")
	assert.NotEqual(t, expected, args[len(args)-1], "Don't append TerraformDir to args when not init, get, plan, apply, destroy")
}

func TestGetCommonOptionsAppendTerraformDir(t *testing.T) {
	t.Parallel()

	expected := "/test/terraform/dir"
	additionalOptions := &Options{ TerraformDir: expected, WorkingDir: "/some/other/dir" }

	_, args := GetCommonOptions(additionalOptions, "init")
	assert.Equal(t, expected, args[len(args)-1], "Do append TerraformDir to args when init, get, plan, apply, destroy")

	_, args = GetCommonOptions(additionalOptions, "get")
	assert.Equal(t, expected, args[len(args)-1], "Do append TerraformDir to args when init, get, plan, apply, destroy")

	_, args = GetCommonOptions(additionalOptions, "plan")
	assert.Equal(t, expected, args[len(args)-1], "Do append TerraformDir to args when init, get, plan, apply, destroy")

	_, args = GetCommonOptions(additionalOptions, "apply")
	assert.Equal(t, expected, args[len(args)-1], "Do append TerraformDir to args when init, get, plan, apply, destroy")

	_, args = GetCommonOptions(additionalOptions, "destroy")
	assert.Equal(t, expected, args[len(args)-1], "Do append TerraformDir to args when init, get, plan, apply, destroy")

}



