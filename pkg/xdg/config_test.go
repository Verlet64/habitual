package xdg_test

import (
	"fmt"
	"habitual/pkg/xdg"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	XDGConfigDirEnvVarTest = "XDG_CONFIG_DIRS"
)

type XDGConfigTestSuite struct {
	suite.Suite
}

func TestPathReturnsFallbackWhenEnvVarUnset(t *testing.T) {
	t.Setenv(XDGConfigDirEnvVarTest, "")

	c := xdg.XDGConfig{}

	got := c.Path()
	want := "/etc/xdg"

	assert.Equal(t, want, got, "expected spec compliant fallback value")
}

func TestPathReturnsSpecifiedValWhenEnvVarSet(t *testing.T) {
	c := xdg.XDGConfig{}

	want := t.TempDir()
	t.Setenv(XDGConfigDirEnvVarTest, fmt.Sprintf("%v", want))

	got := c.Path()
	assert.Equal(t, want, got, "expected to find value in $XDG_CONFIG_DIRS environment variable")
}

func TestPathReturnsPreferredValWhenEnvVarSet(t *testing.T) {
	c := xdg.XDGConfig{}

	want := t.TempDir()
	t.Setenv(XDGConfigDirEnvVarTest, fmt.Sprintf("%v:%v:%v", want, t.TempDir(), t.TempDir()))

	got := c.Path()

	assert.Equal(t, want, got, "expected preferred value in $XDG_CONFIG_DIRS environment variable")
}
