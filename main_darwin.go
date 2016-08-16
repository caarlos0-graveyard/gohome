package gohome

import (
	"os"
	"path/filepath"
)

func cache(appName string) string {
	return filepath.Join(os.Getenv("HOME"), "Library", "Caches", appName)
}

func config(appName string) string {
	return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", appName)
}
