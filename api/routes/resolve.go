package routes

import (
	"os"
	"url-shortener/initializers"
	"url-shortener/models"

	"github.com/gin-gonic/gin"
)

func ResolveURL(c *gin.Context) {

	// get the short from the url
	id := c.Param("id")

	//get id of url

	var urls models.ResponseURL
	initializers.DB.First(&urls, id)

	// respond with data

	c.JSON(200, gin.H{

		"shortLink": os.Getenv("DOMAIN") + "/" + urls.CustomShort,
	})

}
