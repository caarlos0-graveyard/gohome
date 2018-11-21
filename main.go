// +build dragonfly freebsd linux netbsd openbsd solaris

package gohome

import (
	"os"
	"path/filepath"
)

// Cache returns the correct folder to store your apps cache, according to the
// spec of each operating system.
func Cache(appName string) string {
	return filepath.Join(xdgCache(), appName)
}

// Config returns the correct folder to store your apps configuration,
// according to the spec of each operating system.
func Config(appName string) string {
	return filepath.Join(xdgConfig(), appName)
}

func xdgConfig() string {
	xdg := os.Getenv("XDG_CONFIG_HOME")
	if xdg == "" {
		return filepath.Join(os.Getenv("HOME"), ".config")
	}
	return xdg
}

func xdgCache() string {
	xdg := os.Getenv("XDG_CACHE_HOME")
	if xdg == "" {
		return filepath.Join(os.Getenv("HOME"), ".cache")
	}
	return xdg
}
