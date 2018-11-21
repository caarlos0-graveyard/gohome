package gohome

import (
	"os"
	"path/filepath"
)

// Cache returns the correct folder to store your apps cache, according to the
// spec of each operating system.
func Cache(appName string) string {
	return appData(appName)
}

// Config returns the correct folder to store your apps configuration,
// according to the spec of each operating system.
func Config(appName string) string {
	return appData(appName)
}

func appData(appName string) string {
	return filepath.Join(os.Getenv("APPDATA"), appName)
}
