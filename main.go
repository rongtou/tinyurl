package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	_ "tinyurl/config"
	"tinyurl/handler"
)

func main() {
	mode := gin.ReleaseMode
	if viper.GetBool("base.debug") {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	r := gin.Default()

	r.POST("/link/create", handler.CreateLink)
	r.GET("/l/:link", handler.RedirectLink)

	err := r.Run("localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
}
