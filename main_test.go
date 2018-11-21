// +build dragonfly freebsd linux netbsd openbsd solaris

package gohome_test

import (
	"os"
	"testing"

	"github.com/caarlos0/gohome"
	"github.com/stretchr/testify/assert"
)

const appName = "MyXDGApp"

var home = os.Getenv("HOME")

func TestCacheXdgUnset(t *testing.T) {
	cache := gohome.Cache(appName)
	assert.NotEmpty(t, cache)
	assert.Equal(t, home+"/.cache/"+appName, cache)
}

func TestCacheXdgSet(t *testing.T) {
	xdgCache := home + "/.xdgcache/"
	os.Setenv("XDG_CACHE_HOME", xdgCache)
	defer os.Unsetenv("XDG_CACHE_HOME")
	cache := gohome.Cache(appName)
	assert.NotEmpty(t, cache)
	assert.Equal(t, xdgCache+appName, cache)
}

func TestConfigXdgUnset(t *testing.T) {
	config := gohome.Config(appName)
	assert.NotEmpty(t, config)
	assert.Equal(t, home+"/.config/"+appName, config)
}

func TestConfigXdgSet(t *testing.T) {
	xdgConfig := home + "/.xdgconfig/"
	os.Setenv("XDG_CONFIG_HOME", xdgConfig)
	defer os.Unsetenv("XDG_CONFIG_HOME")
	config := gohome.Config(appName)
	assert.NotEmpty(t, config)
	assert.Equal(t, xdgConfig+appName, config)
}
