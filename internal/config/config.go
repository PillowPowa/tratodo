package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string     `yaml:"env" env-default:"development"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	HTTP        HTTPConfig `yaml:"http"`
	JWT         JWTConfig  `yaml:"jwt"`
}

type HTTPConfig struct {
	Host         string `yaml:"host" env-default:"localhost"`
	Port         int    `yaml:"port" env-default:"8080"`
	ReadTimeout  int    `yaml:"read_timeout" env-default:"15"`
	WriteTimeout int    `yaml:"write_timeout" env-default:"15"`
	AllowOrigin  string `yaml:"allow_origin" env-default:"*"`
}

type JWTConfig struct {
	Secret  string `yaml:"secret"`
	ExpDays int    `yaml:"exp_days" env-default:"7"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("Config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("Cannot read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
