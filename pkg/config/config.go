package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DBHost string `mapstructure:"DB_HOST"`
	DBPort int    `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("can't find the file .env :", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("environment can't be loaded: ", err)
	}

	return &env
}
