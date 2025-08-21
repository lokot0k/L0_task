package route

import (
	"github.com/gin-gonic/gin"
)

func AddSPARoutes(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.File("./static/index.html")
	})
}
