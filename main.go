package main

import (
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"tinyurl/cache"
	_ "tinyurl/config"
	"tinyurl/handler"
)

func main() {
	mode := gin.ReleaseMode
	if viper.GetBool("base.debug") {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	cache.Init()

	r := gin.Default()
	r.Use(location.Default())

	r.POST("/link/create", handler.CreateLink)
	r.GET("/l/:link", handler.RedirectLink)

	addr := ":" + viper.GetString("base.port")
	if mode == gin.DebugMode {
		addr = "localhost:" + viper.GetString("base.port")
	}
	err := r.Run(addr)
	if err != nil {
		log.Fatalln(err)
	}
}
