package gohome

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	app := "my-app"
	cache := Cache(app)
	assert.NotEmpty(t, cache)
	assert.True(t, strings.HasSuffix(cache, app))
}

func TestConfig(t *testing.T) {
	app := "another-app"
	config := Config(app)
	assert.NotEmpty(t, config)
	assert.True(t, strings.HasSuffix(config, app))
}
