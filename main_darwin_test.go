package gohome_test

import (
	"os"
	"testing"

	"github.com/caarlos0/gohome"
	"github.com/stretchr/testify/assert"
)

const appName = "MyApp"

var home = os.Getenv("HOME")

func TestCache(t *testing.T) {
	cache := gohome.Cache(appName)
	assert.NotEmpty(t, cache)
	assert.Equal(t, home+"/Library/Caches/"+appName, cache)
}

func TestConfig(t *testing.T) {
	config := gohome.Config(appName)
	assert.NotEmpty(t, config)
	assert.Equal(t, home+"/Library/Application Support/"+appName, config)
}
