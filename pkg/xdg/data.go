package xdg

import (
	"os"
	"strings"
)

type XDGData struct{}

const (
	XDGDataDirEnvVar    = "XDG_DATA_DIRS"
	FallbackXDGDataPath = "/usr/local/share/:/usr/share/"
)

func (x XDGData) Path() string {
	if dirs := os.Getenv(XDGDataDirEnvVar); dirs != "" {
		d := strings.Split(dirs, ":")
		return d[0]
	}

	return FallbackXDGDataPath
}
