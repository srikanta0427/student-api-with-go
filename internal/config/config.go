package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServerConfig struct {
	Port string						`yaml:"port"`
	Host string						`yaml:"host"`
}

type Config struct {
	Env string						`yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string				`yaml:"storagePath"`
	HttpServer HttpServerConfig		`yaml:"httpServer"`
}

func MustLoad() *Config{
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags:= flag.String("config","","path to the configuration file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("config path is not set")
		}
	}

	if _,err:= os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file not found %s ", configPath)
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("error reading config file: %v", err.Error())
	}

	return &cfg
}