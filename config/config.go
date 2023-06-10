package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"ozon/pkg/logger"
	"sync"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env:"is_debug" env-default:"true"`
	Listen  struct {
		Type   string `yaml:"type" env:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env:"bind_ip" env-default:"0.0.0.0"`
		Port   string `yaml:"port" env:"port" env-default:"8080"`
	} `yaml:"listen" env:"listen"`
	PsqlStorage  `yaml:"psqlStorage"`
	RedisStorage `yaml:"redisStorage"`
}

type PsqlStorage struct {
	Username string `yaml:"username" env:"POSTGRES_USER" env-default:"postgres"`
	Password string `yaml:"password" env:"POSTGRES_PASSWORD" env-default:"12345"`
	Host     string `yaml:"host" env:"POSTGRES_HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"POSTGRES_PORT" env-default:"5432"`
	Database string `yaml:"database" env:"POSTGRES_DB" env-default:"url"`
}
type RedisStorage struct {
	Username string `yaml:"username" env:"REDIS_USER" env-default:"user"`
	Host     string `yaml:"host" env:"REDIS_HOST" env-default:"redis"`
	Port     string `yaml:"port" env:"REDIS_PORT" env-default:"6379"`
	Database string `yaml:"database" env:"REDIS_DB" env-default:"0"`
}

var instance *Config
var once sync.Once

func GetConfigYml() *Config {
	once.Do(func() {
		log := logger.GetLogger()
		log.Info().Msg("read application config")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Info().Msg(help)
			log.Err(err)
		}
	})
	return instance
}
func GetConfigEnv() *Config {
	log := logger.GetLogger()
	once.Do(func() {
		log.Info().Msg("read application config")
		instance = &Config{}
		if err := cleanenv.ReadEnv(&instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Info().Msg(help)
			log.Err(err)
		}
		if err := cleanenv.ReadEnv(&instance.Listen); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Info().Msg(help)
			log.Err(err)
		}
		if err := cleanenv.ReadEnv(&instance.PsqlStorage); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Info().Msg(help)
			log.Err(err)
		}
		if err := cleanenv.ReadEnv(&instance.RedisStorage); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Info().Msg(help)
			log.Err(err)
		}
	})
	return instance
}
