package handler

import (
	"L0_task/internal/repository"
	"L0_task/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetOrder(c *gin.Context) {
	db := c.MustGet("database").(*gorm.DB)
	orderID := c.Param("orderID")
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
	c.JSON(http.StatusOK, order)
}
