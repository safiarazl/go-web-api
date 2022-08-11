package main

import (
	"github.com/gin-gonic/gin"
	"github.com/safiarazl/go-web-api/handler"
)

func main()  {
	r := gin.Default()
	//CARA GA BIASA (Anonymous function)
	// r.GET("/ping", func(c *gin.Context) { 
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.GET("/", handler.RootHandler)
	r.GET("/ping", handler.PingHandler)
	r.GET("/pages/:id/:title", handler.PagesHandler)
	r.GET("/query", handler.QueryHandler)
	r.POST("/query", handler.PostQueryHandler)
	

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}