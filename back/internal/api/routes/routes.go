package routes

import "github.com/gin-gonic/gin"

func RouterSetup(r *gin.Engine) {
	MiddlewaresSetup(r)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
}

func MiddlewaresSetup(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
}
