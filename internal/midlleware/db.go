package midlleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DBMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// пробрасываем DB в контекст
		c.Set("database", db)
		c.Next()
	}
}
