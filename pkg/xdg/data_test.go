package xdg_test

import (
	"fmt"
	"habitual/pkg/xdg"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	XDGDataDirEnvVarTest = "XDG_DATA_DIRS"
)

type XDGDataTestSuite struct {
	suite.Suite
}

func TestXDGDataPathReturnsFallbackWhenEnvVarUnset(t *testing.T) {
	t.Setenv(XDGDataDirEnvVarTest, "")

	c := xdg.XDGData{}

	got := c.Path()
	want := "/usr/local/share/:/usr/share/"

	assert.Equal(t, want, got, "expected spec compliant fallback value")
}

func TestXDGDataPathReturnsSpecifiedValWhenEnvVarSet(t *testing.T) {
	c := xdg.XDGData{}

	want := t.TempDir()
	t.Setenv(XDGDataDirEnvVarTest, fmt.Sprintf("%v", want))

	got := c.Path()
	assert.Equal(t, want, got, "expected to find value in $XDG_DATA_DIRS environment variable")
}

func TestXDGDataPathReturnsPreferredValWhenEnvVarSet(t *testing.T) {
	c := xdg.XDGData{}

	want := t.TempDir()
	t.Setenv(XDGDataDirEnvVarTest, fmt.Sprintf("%v:%v:%v", want, t.TempDir(), t.TempDir()))

	got := c.Path()

	assert.Equal(t, want, got, "expected preferred value in $XDG_DATA_DIRS environment variable")
}
