package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bruhabruh/rpfont2image/config"
	"github.com/bruhabruh/rpfont2image/internal/entity"
	"log"
	"os"
	"strings"
)

var (
	ErrFontFileNotExists = errors.New("font file not exists")
)

func Run(cfg *config.Config) {
	if err := prepareCheck(cfg); err != nil {
		log.Fatal(err)
	}

	fontFileRaw, err := os.ReadFile(cfg.FontFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var providers []map[string]interface{}
	if err = json.Unmarshal(fontFileRaw, &providers); err != nil {
		log.Fatal(err)
	}
	bitMapProviders := entity.BitMapProviders(providers)

	for _, provider := range bitMapProviders {
		prefix := "minecraft"
		filePath := provider.File

		if strings.Contains(filePath, ":") {
			split := strings.SplitN(filePath, ":", -1)
			if len(split) > 0 {
				prefix = split[0]
				filePath = split[1]
			}
		}

		imageFilePath := fmt.Sprintf("%s/assets/%s/textures/%s", cfg.InputPath, prefix, filePath)
		if _, err = os.Stat(imageFilePath); os.IsNotExist(err) {
			log.Printf("File not exists: %s", imageFilePath)
			continue
		}

		imageFile, err := os.ReadFile(imageFilePath)
		if err != nil {
			log.Printf("Fail read file: %s", imageFilePath)
			continue
		}

		split := strings.SplitN(imageFilePath, ".", -1)
		imageFileExtension := split[len(split)-1]

		if _, err = os.Stat(cfg.OutputPath); os.IsNotExist(err) {
			if err = os.MkdirAll(cfg.OutputPath, os.ModePerm); err != nil {
				log.Fatal("Fail create output directory")
			}
		}

		for _, fileName := range provider.Chars {
			destination := fmt.Sprintf("%s/%s.%s", cfg.OutputPath, fileName, imageFileExtension)
			if err = os.WriteFile(destination, imageFile, os.ModePerm); err != nil {
				log.Printf("Fail write file: %s", destination)
			}
		}
	}
}

func prepareCheck(cfg *config.Config) error {
	if _, err := os.Stat(cfg.FontFilePath); os.IsNotExist(err) {
		return ErrFontFileNotExists
	}
	return nil
}
