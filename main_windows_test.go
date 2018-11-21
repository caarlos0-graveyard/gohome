package gohome

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const appName = "MyWinApp"
const appdata = "C:/Users/john-doe/AppData/Local/"

func TestCacheWindows(t *testing.T) {
	os.Setenv("APPDATA", appdata)
	defer os.Unsetenv("APPDATA")
	cache := Cache(appName)
	assert.NotEmpty(t, cache)
	assert.Equal(t, appdata+appName, cache)
}

func TestConfigWindows(t *testing.T) {
	os.Setenv("APPDATA", appdata)
	defer os.Unsetenv("APPDATA")
	config := Config(appName)
	assert.NotEmpty(t, config)
	assert.Equal(t, appdata+appName, config)
}
