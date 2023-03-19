package config

import "github.com/Netflix/go-env"

type environment struct {
	DataBaseURL string `env:"DB_URL"`
	DbName      string `env:"DB_NAME"`
	DbUser      string `env:"DB_USER"`
	DbPassword  string `env:"DB_PASSWORD"`
}

var Env environment

func LoadEnv() (err error) {
	_, err = env.UnmarshalFromEnviron(Env)
	return
}
