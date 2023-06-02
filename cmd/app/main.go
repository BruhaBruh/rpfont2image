package main

import (
	"github.com/bruhabruh/rpfont2image/config"
	"github.com/bruhabruh/rpfont2image/internal/app"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
