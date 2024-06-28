package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Port        uint16        `yaml:"port" env-default:"8080"`
	Adress      string        `yaml:"adress" env-default:"127.0.0.1"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type Config struct {
	Http HttpServer `yaml:"http"`
	Env  string     `yaml:"env" env-default:"local"`
}

func fetchConfPath() string {
	var conf string

	flag.StringVar(&conf, "config", "", "path to config file")
	flag.Parse()

	if conf == "" {
		log.Fatal("config pth is required\n\nuse flag [ --config ]\n")
	}

	return conf
}

func MustLoad() Config {
	confPath := fetchConfPath()
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		log.Fatalf("file %s - is not exist", confPath)
	}

	cfg := Config{}

	if err := cleanenv.ReadConfig(confPath, &cfg); err != nil {
		log.Fatalf("unable to parse config\nreason: %s", err.Error())
	}

	return cfg
}
