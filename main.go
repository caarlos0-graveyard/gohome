package gohome

import "os"

// Cache returns the correct folder to store your apps cache, according to the
// spec of each operating system.
func Cache(appName string) string {
	return cache(appName)
}

// Config returns the correct folder to store your apps configuration,
// according to the spec of each operating system.
func Config(appName string) string {
	return config(appName)
}

func home() string {
	return os.Getenv("HOME")
}
