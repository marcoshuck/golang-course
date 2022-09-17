package main

import (
	"log"
	"os"
)

type Config struct {
	DB struct {
		Host     string
		User     string
		Password string
		Database string
	}

	Server struct {
		Port uint16
	}

	AWS AWSConfig
}

type AWSConfig struct {
	AccessKey       string
	SecretAccessKey string
}

func main() {
	var config Config
	var err error

	if config, err = setup(); err != nil {
		log.Println("Error setting the application up.")
		os.Exit(1)
	}

	if err = run(config); err != nil {
		log.Println("Error running the application.")
		os.Exit(2)
	}
}

func setup() (Config, error) {
	return Config{
		DB: struct {
			Host     string
			User     string
			Password string
			Database string
		}{
			Host:     "",
			User:     "",
			Password: "",
			Database: "",
		},
		Server: struct {
			Port uint16
		}{
			Port: 8080,
		},
		AWS: struct {
			AccessKey       string
			SecretAccessKey string
		}{
			AccessKey:       os.Getenv("AWS_ACCESS_KEY"),
			SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		},
	}, nil
}

func run(cfg Config) error {
	return nil
}
