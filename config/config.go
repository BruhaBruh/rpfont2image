package config

import (
	"flag"
	"fmt"
	"path"
)

type Config struct {
	InputPath    string
	OutputPath   string
	FontFilePath string
}

func New() *Config {
	var config Config
	flag.StringVar(&config.InputPath, "input", "./resourcepack", "resourcepack directory")
	flag.StringVar(&config.OutputPath, "output", "./output", "output directory for images with symbols in filename")
	flag.Parse()

	config.InputPath = path.Base(config.InputPath)
	config.OutputPath = path.Base(config.OutputPath)
	config.FontFilePath = fmt.Sprintf(`%s/assets/minecraft/font/default.json`, config.InputPath)

	return &config
}
