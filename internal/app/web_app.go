package app

import (
	"L0_task/internal/config"
	"L0_task/internal/database"
	"L0_task/internal/midlleware"
	"L0_task/internal/route"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RunWebServer(cfg *config.Config) {
	// подключение к хранилищам
	db := database.MustLoad(cfg)
	engine := gin.New()

	// подключение мидлвари для всех роутов
	engine.Use(gin.Recovery(), midlleware.DBMiddleware(db))

	// подключение роутов
	route.AddOrderRoutes(engine)

	// запуск сервера
	address := fmt.Sprintf("%s:%s", cfg.AppIP, cfg.AppPort)
	err := engine.Run(address)
	if err != nil {
		panic(err)
	}
}
