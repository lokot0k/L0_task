package main

import (
	"L0_task/internal/app"
	"L0_task/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	app.RunWebServer(cfg)
}
