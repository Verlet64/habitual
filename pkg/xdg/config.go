package xdg

import (
	"os"
	"strings"
)

type XDGConfig struct{}

const (
	XDGConfigDirEnvVar    = "XDG_CONFIG_DIRS"
	FallbackXDGConfigPath = "/etc/xdg"
)

func (x XDGConfig) Path() string {
	if dirs := os.Getenv(XDGConfigDirEnvVar); dirs != "" {
		d := strings.Split(dirs, ":")
		return d[0]
	}

	return FallbackXDGConfigPath
}
