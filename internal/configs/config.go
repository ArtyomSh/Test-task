package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
	"sync"
)

type Config struct {
	Listen struct {
		Port string `yaml:"port" env-default:":3001"`
	} `yaml:"listen"`
	Binance struct {
		URL string `json:"url"`
	} `yaml:"binance"`
	Repository struct {
		Repo string `yaml:"repo"`
	} `yaml:"repository"`
	Redis struct {
		Port     string `yaml:"port" env-default:":6379"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`
}

var instance *Config
var once sync.Once

func GetConfig(logger *zap.Logger) *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("configs/config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
		}
	})
	return instance
}
