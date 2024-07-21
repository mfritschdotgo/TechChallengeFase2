package configs

import (
	"os"
)

type Configs struct {
	MONGO_USER     string `mapstructure:"MONGO_USER"`
	MONGO_PASSWORD string `mapstructure:"MONGO_PASSWORD"`
	MONGO_HOST     string `mapstructure:"MONGO_HOST"`
	MONGO_PORT     string `mapstructure:"MONGO_PORT"`
	MONGO_DATABASE string `mapstructure:"MONGO_DATABASE"`
	HTTP_PORT      string `mapstructure:"HTTP_PORT"`
}

func GetConfig() *Configs {
	return &Configs{
		MONGO_USER:     os.Getenv("MONGO_USER"),
		MONGO_PASSWORD: os.Getenv("MONGO_PASSWORD"),
		MONGO_HOST:     os.Getenv("MONGO_HOST"),
		MONGO_PORT:     os.Getenv("MONGO_PORT"),
		MONGO_DATABASE: os.Getenv("MONGO_DATABASE"),
		HTTP_PORT:      os.Getenv("HTTP_PORT"),
	}
}
