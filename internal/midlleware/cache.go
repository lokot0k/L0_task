package midlleware

import (
	"L0_task/pkg/cache"
	"github.com/gin-gonic/gin"
)

func CacheMiddleware(cache *cache.Cache) gin.HandlerFunc {
	return func(c *gin.Context) {
		// пробрасываем кэш в контекст
		c.Set("cache", cache)
		c.Next()
	}
}
