package config

import (
	"errors"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfigReadInConfigFails(t *testing.T) {
	invalidPath := "./invalid"

	assert.Panics(t, func() {
		_, _ = LoadConfig(invalidPath)
	}, "LoadConfig should panic when ReadInConfig fails")
}

func TestLoadConfigPanicOnUnmarshalError(t *testing.T) {
	envContent := `
WEB_SERVER_PORT='8000'
CEP_API_URL=https://api.example.com.br/ws/json/
WEATHERZIP_API_URL=http://api.example.com.br/v1/
`
	envFilePath := ".env"
	err := os.WriteFile(envFilePath, []byte(envContent), 0644)
	assert.NoError(t, err)
	defer os.Remove(envFilePath)

	ViperUnmarshal = func(rawVal any, opts ...viper.DecoderConfigOption) error {
		return errors.New("Unmarshal error")
	}

	defer func() {
		ViperUnmarshal = viper.Unmarshal
		if r := recover(); r == nil {
			t.Errorf("Expected a panic in Unmarshal, but none was triggered")
		} else {
			t.Logf("Panic successfully captured: %v", r)
		}
	}()

	if _, err := LoadConfig("."); err == nil {
		t.Errorf("Expected an error, but none was returned")
	}
}

func TestLoadConfigMissingRequiredConfigFails(t *testing.T) {
	envContent := `
WEB_SERVER_PORT='8000'
CEP_API_URL=https://api.example.com.br/ws/json/
WEATHERZIP_API_URL=
`
	envFilePath := ".env"
	err := os.WriteFile(envFilePath, []byte(envContent), 0644)
	assert.NoError(t, err)
	defer os.Remove(envFilePath)

	assert.Panics(t, func() {
		_, _ = LoadConfig(".")
	}, "LoadConfig should panic when Unmarshal fails")
}

func TestLoadConfig(t *testing.T) {
	envContent := `
WEB_SERVER_PORT=8000
CEP_API_URL=http://example.com/cep
WEATHERZIP_API_URL=http://example.com/weather
`
	envFilePath := ".env"
	err := os.WriteFile(envFilePath, []byte(envContent), 0644)
	assert.NoError(t, err)
	defer os.Remove(envFilePath)

	cfg, err := LoadConfig(".")
	assert.NoError(t, err)
	assert.NotNil(t, cfg)

	assert.Equal(t, "8000", cfg.WebServerPort)
	assert.Equal(t, "http://example.com/cep", cfg.CepAPIUrl)
	assert.Equal(t, "http://example.com/weather", cfg.WeatherZipAPIUrl)
}
