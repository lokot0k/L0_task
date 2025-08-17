package route

import (
	handler "L0_task/internal/handler/web_app"
	"github.com/gin-gonic/gin"
)

func AddOrderRoutes(r *gin.Engine) {
	group := r.Group("/order")
	group.GET("/:orderID", handler.GetOrder)
}
