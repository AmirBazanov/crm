package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env   string     `yaml:"env" env-default:"local"`
	DbUrl string     `yaml:"db_url" env-required:"true"`
	GRPC  GRPCConfig `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env-default:"4000"`
	Timeout time.Duration `yaml:"timeout" env-default:"10h"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config file not exist")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file not exist" + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config file" + err.Error())
	}
	return &cfg
}

func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "config.yaml", "config file path")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
