package gohome

import "path/filepath"

func cache(appName string) string {
	return filepath.Join(home(), "Library", "Caches", appName)
}

func config(appName string) string {
	return filepath.Join(home(), "Library", "Application Support", appName)
}
