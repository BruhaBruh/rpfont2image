package main

import (
	"github.com/bruhabruh/rpfont2image/config"
	"github.com/bruhabruh/rpfont2image/internal/app"
)

func main() {
	cfg := config.New()

	app.Run(cfg)
}
