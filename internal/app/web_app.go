package app

import (
	"L0_task/internal/config"
	"L0_task/internal/database"
	"L0_task/internal/midlleware"
	"L0_task/internal/route"
	"L0_task/pkg/cache"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunWebServer(cfg *config.Config) {
	// подключение к хранилищам
	db := database.MustLoad(cfg)
	engine := gin.New()

	// подключение мидлвари для всех роутов
	engine.Use(gin.Recovery(), midlleware.DBMiddleware(db))
	engine.Use(gin.Recovery(), midlleware.CacheMiddleware(cache.NewCache(cfg.CacheLimit)))
	engine.Use(cors.Default())

	// подключение роутов
	route.AddOrderRoutes(engine)
	route.AddSPARoutes(engine)

	// запуск сервера
	address := fmt.Sprintf("%s:%s", cfg.AppIP, cfg.AppPort)
	err := engine.Run(address)
	if err != nil {
		panic(err)
	}
}
