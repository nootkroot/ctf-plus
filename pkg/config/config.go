package config

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Theme string `yaml:"theme", envconfig:"CTF_THEME"`
}

func loadFromFile(conf *Config, path string) {
	// Load from config file first
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error while loading config file: %s", err)
		return
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		log.Fatalf("Error while getting size of config: %s", err)
	}

	if fi.Size() == 0 {
		log.Println("No config found, using defaults...")
		return
	}

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatalf("Error while parsing config file: %v", err)
	}
}

func loadFromEnv(conf *Config) {
	// check if .env file exists
	if _, err := os.Stat("sample.txt"); err != nil {
		if !os.IsNotExist(err) {
			log.Fatalf("Error while checking for .env: %s", err)
		}
	} else {
		// load .env file
		err := godotenv.Load(".env")

		if err != nil {
			log.Fatalf("Error loading .env file: %s", err)
		}
	}

	err := envconfig.Process("ctf", conf)
    if err != nil {
        log.Printf("Error while loading config from environment: %s", err)
    }
}

func LoadConfig() (*Config) {
	// Check for a custom config path
	path := os.Getenv("CTF_CONFIG_PATH")
	if path == "" {
		path = "config.yaml"
	}

	// Initiialize config with default values
	conf := new(Config)
	conf.Theme = "default"

	// Load from config file first
	loadFromFile(conf, path)

	// Then load from envirnoment variables since
	// they are given a higher priority.
	// .env file has higher priority than os env
	loadFromEnv(conf)
	
	return conf
}