package gohome

import (
	"os"
	"path/filepath"
)

func cache(appName string) string {
	return filepath.Join(xdgCache(), appName)
}

func config(appName string) string {
	return filepath.Join(xdgConfig(), appName)
}

func xdgConfig() string {
	xdg := os.Getenv("XDG_CONFIG_HOME")
	if xdg == "" {
		return filepath.Join(home(), ".config")
	}
	return xdg
}

func xdgCache() string {
	xdg := os.Getenv("XDG_CACHE_HOME")
	if xdg == "" {
		return filepath.Join(home(), ".cache")
	}
	return xdg
}
