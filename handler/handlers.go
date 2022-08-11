package handler

import (
	"fmt"
	_"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type QueryInput struct {
	// validator input
	Title string `json:"title" binding:"required"`
	Price int `json:"price_per_pcs" binding:"required,number"`
	Summ int `json:"summ" binding:"required,number"`
}

func PostQueryHandler(c *gin.Context) {
	// localhost:8080/query
	var queryInput QueryInput
	err := c.ShouldBindJSON(&queryInput)
	if err != nil {
		// log.Fatal(err) //akan memberhentikan server jika terjadi error
		c.JSON(http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": queryInput.Title,
		"price_per_pcs": queryInput.Price,
		"summ": queryInput.Summ,
	})
}
func PagesHandler(c *gin.Context) {
	//localhost:8080/pages/12/Aquaman
	id:= c.Param("id")
	title:= c.Param("title")
	c.JSON(http.StatusOK, gin.H{
		"pages": id,
		"title": title,
	})
}
func QueryHandler(c *gin.Context) {
	// localhost:8080/query?title=saya manusia&price=45000
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Safiar Azalia",
		"nim": "A11.2020.12715",
	})
}