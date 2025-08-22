package handler

import (
	"L0_task/internal/models"
	"L0_task/internal/repository"
	"L0_task/internal/service"
	"L0_task/pkg/cache"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func GetOrder(c *gin.Context) {
	db := c.MustGet("database").(*gorm.DB)
	cacheDep := c.MustGet("cache").(*cache.Cache)
	orderID := c.Param("orderID")
	cacheEntry := cacheDep.Get(orderID)
	if cacheEntry != nil {
		order := cacheEntry.(*models.Order)
		c.JSON(http.StatusOK, order)
		return
	}
	order, err := service.GetOrderByID(orderID, repository.NewOrderRepository(db))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "order not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}
		return
	}
	cacheDep.Put(orderID, &models.CachableOrder{Order: order, LastUsed: time.Now()})
	c.JSON(http.StatusOK, order)
}
