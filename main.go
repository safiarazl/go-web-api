package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()
	//CARA GA BIASA (Anonymous function)
	// r.GET("/ping", func(c *gin.Context) { 
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.GET("/", rootHandler)
	r.GET("/ping", pingHandler)
	r.GET("/pages/:id/:title", pagesHandler)
	r.GET("/query", queryHandler)
	r.POST("/query", postQueryHandler)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type QueryInput struct {
	Title string
	Price int `json:"price_per_pcs"`
	Summ int
}

func postQueryHandler(c *gin.Context) {
	// localhost:8080/query
	var queryInput QueryInput
	err := c.ShouldBindJSON(&queryInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"title": queryInput.Title,
		"price_per_pcs": queryInput.Price,
		"summ": queryInput.Summ,
	})
}
func pagesHandler(c *gin.Context) {
	//localhost:8080/pages/12/Aquaman
	id:= c.Param("id")
	title:= c.Param("title")
	c.JSON(http.StatusOK, gin.H{
		"pages": id,
		"title": title,
	})
}
func queryHandler(c *gin.Context) {
	// localhost:8080/query?title=saya manusia&price=45000
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}
func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Safiar Azalia",
		"nim": "A11.2020.12715",
	})
}