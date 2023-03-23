package config

import "github.com/Netflix/go-env"

type environment struct {
	ApiPort     string `env:"API_PORT"`
	DataBaseURL string `env:"DB_URL"`
	DbName      string `env:"DB_NAME"`
	DbUser      string `env:"DB_USER"`
	DbPassword  string `env:"DB_PASSWORD"`
	DbHost      string `env:"DB_HOST"`
	DbPort      string `env:"DB_PORT"`
}

var Env environment

func LoadEnv() (err error) {
	_, err = env.UnmarshalFromEnviron(&Env)
	return
}