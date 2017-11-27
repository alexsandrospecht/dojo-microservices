package commons

import (
	"github.com/creamdog/gonfig"
)

type Config struct {
	Yml gonfig.Gonfig
}

func (config *Config) GetValue(key string) string {
	value, _ := config.Yml.GetString(key, "")
	return value
}