package config

import (
	"os"
)

type Config struct {
	ModelPath string
	Port      string
}

func Load() *Config {
	modelPath := os.Getenv("MITIE_MODEL_PATH")
	if modelPath == "" {
		modelPath = "models/ner_model.dat"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		ModelPath: modelPath,
		Port:      port,
	}
}
