package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB_Source string `mapstructure:"DB_SOURCE"`
}

func LoadEnv() *Config {
	viper.AddConfigPath("./pkg/env")
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	c := &Config{}
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalln(err)
	}
	return c
}
