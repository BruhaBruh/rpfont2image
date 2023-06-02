package config

import (
	"errors"
	"flag"
	"fmt"
	"path"
)

var (
	ErrInputDirectoryEmpty = errors.New("input directory not exists")
)

type Config struct {
	InputPath    string
	OutputPath   string
	FontFilePath string
}

func New() (*Config, error) {
	var config Config
	flag.StringVar(&config.InputPath, "input", "", "resourcepack directory")
	flag.StringVar(&config.OutputPath, "output", "./output", "output directory for images with symbols in filename")
	flag.Parse()

	if config.InputPath == "" {
		return nil, ErrInputDirectoryEmpty
	}

	config.InputPath = path.Base(config.InputPath)
	config.OutputPath = path.Base(config.OutputPath)
	config.FontFilePath = fmt.Sprintf(`%s/assets/minecraft/font/default.json`, config.InputPath)

	return &config, nil
}
