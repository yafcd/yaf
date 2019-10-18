package main

import (
	"yafcd.dev/yaf/sync"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "pong",
		})
	})

	r.POST("/sync/:repo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

		sync.SyncFn()
	})

	r.Run("0.0.0.0:8081") // listen and serve on 0.0.0.0:8080
}
