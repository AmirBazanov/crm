package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env    string       `yaml:"env" env-default:"local"`
	DbUrl  string       `yaml:"db_url" env-required:"true"`
	GRPC   GRPCConfig   `yaml:"grpc"`
	Logger LoggerConfig `yaml:"logger"`
	Redis  RedisConfig  `yaml:"redis"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	DB       int    `yaml:"db" env-default:"0"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env-default:"4000"`
	Timeout time.Duration `yaml:"timeout" env-default:"10h"`
}

type LoggerConfig struct {
	Service  string `yaml:"service" env-default:"users"`
	LogLevel string `yaml:"log_level" env-default:"debug"`
	LogFile  string `yaml:"log_file" env-default:"logs/users.log"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		return MustLoadEnv()
	}

	return MustLoadPath(configPath)
}

func MustLoadEnv() *Config {
	var config Config
	err := cleanenv.ReadEnv(&config)
	if err != nil {
		panic("config env read error: " + err.Error())
	}
	return &config
}

func MustLoadPath(configPath string) *Config {
	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "config file path")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
