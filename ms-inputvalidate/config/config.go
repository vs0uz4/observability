package config

import (
	"github.com/spf13/viper"
)

var ViperUnmarshal = viper.Unmarshal

type conf struct {
	WebServerPort    string `mapstructure:"WEB_SERVER_PORT"`
	ZipKinUrl        string `mapstructure:"ZIPKIN_URL"`
	WeatherZipAPIUrl string `mapstructure:"WEATHERZIP_API_URL"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	viper.SetDefault("WEB_SERVER_PORT", ":8000")
	viper.SetDefault("ZIPKIN_URL", "http://zipkin:9411/api/v2/spans")
	viper.SetDefault("WEATHERZIP_API_URL", "http://ms-weatherzip:8001/weather/%s")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = ViperUnmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	if cfg.WebServerPort == "" || cfg.ZipKinUrl == "" || cfg.WeatherZipAPIUrl == "" {
		panic("missing required configuration")
	}

	return cfg, err
}
