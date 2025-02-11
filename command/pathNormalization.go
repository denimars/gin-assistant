package command

import (
	"runtime"
	"strings"
)

func PathNormalization(path string) string {
	switch runtime.GOOS {
	case "windows":
		path_ := strings.ReplaceAll(path, `\`, `\\`)
		path__ := strings.ReplaceAll(path_, `/`, `\\`)
		return path__
	default:
		return path
	}
}
