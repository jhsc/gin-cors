package cors

import (
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func newTestRouter(opts Options) *gin.Engine {
	router := gin.New()
	router.Use(Middleware(opts))
	router.GET("/", func(c *gin.Context) {
		c.String(200, "get")
	})
	router.POST("/", func(c *gin.Context) {
		c.String(200, "post")
	})
	router.PUT("/", func(c *gin.Context) {
		c.String(200, "put")
	})
	router.PATCH("/", func(c *gin.Context) {
		c.String(200, "patch")
	})
	return router
}
