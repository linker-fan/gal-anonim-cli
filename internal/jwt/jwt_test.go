package jwt

import (
	"testing"

	"github.com/linker-fan/gal-anonim-cli/internal/config"
	"github.com/stretchr/testify/assert"
)

const (
	token = "a-random-generated-test-token"
)

func TestSaveTokenToFile(t *testing.T) {
	c, err := config.NewConfig("./../../config.yml")
	assert.NoError(t, err)
	err = SaveTokenToFile(c.Local.TokeFilePath, token)
	assert.NoError(t, err)
}

func TestLoadTokenFromFile(t *testing.T) {
	c, err := config.NewConfig("./../../config.yml")
	assert.NoError(t, err)
	tokenFromFile, err := LoadTokenFromFile(c.Local.TokeFilePath)
	assert.NoError(t, err)
	assert.Equal(t, tokenFromFile, token)
}
