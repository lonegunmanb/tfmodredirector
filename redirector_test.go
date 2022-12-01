package tfmodredirector_test

import (
	"testing"

	"github.com/lonegunmanb/tfmodredirector"
	"github.com/stretchr/testify/assert"
)

func Test_GetModules(t *testing.T) {
	tf := `
module "mod" {
  source = "../../"
}
`
	names, err := tfmodredirector.GetModules(tf, "../../")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(names))
	assert.Equal(t, "mod", names[0])
}

func Test_RewriteModuleSource(t *testing.T) {
	tf := `
module "mod" {
  source = "../../"
}
`
	actual, err := tfmodredirector.RewriteModuleSource(tf, "mod", "new_source")
	assert.Nil(t, err)
	expected := `
module "mod" {
  source = "new_source"
}
`
	assert.Equal(t, expected, actual)
}

func Test_RedirectModuleSource(t *testing.T) {
	tf := `
module "mod" {
  source = "../../"
}

module "mod2" {
  source = "github.com"
}
`
	actual, err := tfmodredirector.RedirectModuleSource(tf, "../../", "new_source")
	assert.Nil(t, err)
	expected := `
module "mod" {
  source = "new_source"
}

module "mod2" {
  source = "github.com"
}
`
	assert.Equal(t, expected, actual)
}
