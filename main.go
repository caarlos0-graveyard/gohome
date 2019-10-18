package gohome

import (
	"os"
	"path/filepath"
)

// Cache returns the correct folder to store your apps cache, according to the
// spec of each operating system.
//
// Deprecated: use os.UserCacheDir()
func Cache(appName string) string {
	s, _ := os.UserCacheDir()
	return filepath.Join(s, appName)
}

// Config returns the correct folder to store your apps configuration,
// according to the spec of each operating system.
//
// Deprecated: use os.UserConfigDir()
func Config(appName string) string {
	s, _ := os.UserConfigDir()
	return filepath.Join(s, appName)
}
