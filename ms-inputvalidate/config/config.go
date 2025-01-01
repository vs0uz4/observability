package config

import (
	"github.com/spf13/viper"
)

var ViperUnmarshal = viper.Unmarshal

type conf struct {
	WebServerPort    string `mapstructure:"WEB_SERVER_PORT"`
	ZipKinUrl        string `mapstructure:"ZIPKIN_URL"`
	CepAPIUrl        string `mapstructure:"CEP_API_URL"`
	WeatherZipAPIUrl string `mapstructure:"WEATHERZIP_API_URL"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = ViperUnmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	if cfg.WebServerPort == "" || cfg.CepAPIUrl == "" || cfg.WeatherZipAPIUrl == "" {
		panic("missing required configuration")
	}

	return cfg, err
}
