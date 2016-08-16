package gohome_test

import (
	"os"
	"testing"

	"github.com/caarlos0/gohome"
	"github.com/stretchr/testify/assert"
)

const appName = "MyWinApp"
const appdata = "C:/Users/john-doe/AppData/Local/"

func TestCacheWindows(t *testing.T) {
	os.Setenv("APPDATA", appdata)
	defer os.Unsetenv("APPDATA")
	cache := gohome.Cache(appName)
	assert.NotEmpty(t, cache)
	assert.Equal(t, appdata+appName, cache)
}

func TestConfigWindows(t *testing.T) {
	os.Setenv("APPDATA", appdata)
	defer os.Unsetenv("APPDATA")
	config := gohome.Config(appName)
	assert.NotEmpty(t, config)
	assert.Equal(t, appdata+appName, config)
}
