package main

import (
	"url-shortener/initializers"
	"url-shortener/routes"

	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDb()

}

func main() {

	r := gin.Default()

	r.POST("/api/v1", routes.ShortenURL)
	r.GET("/:url", routes.ResolveURL)

	r.Run()

}
