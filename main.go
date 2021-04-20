package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"tinyurl/handler"
)

func main() {
	r := gin.Default()

	r.POST("/link/create", handler.CreateLink)
	r.GET("/l/:link", handler.RedirectLink)
	err := r.Run("localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
}
