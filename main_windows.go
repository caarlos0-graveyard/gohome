package gohome

import (
	"os"
	"path/filepath"
)

func cache(appName string) string {
	return appData(appName)
}

func config(appName string) string {
	return appData(appName)
}

func appData(appName string) string {
	return filepath.Join(os.Getenv("APPDATA"), appName)
}
