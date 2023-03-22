package routes

import (
	"log"
	"os"
	"time"
	"url-shortener/helpers"
	"url-shortener/initializers"
	"url-shortener/models"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type response struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

func ShortenURL(c *gin.Context) {

	var body struct {
		URL string
		//CustomShort string
		Expiry time.Duration
	}

	c.Bind(&body)

	// check if the input is an actual URL
	if !govalidator.IsURL(body.URL) {

		log.Fatal("Invalid URL")
		return
	}

	// check for the domain error
	// users may abuse the shortener by shorting the domain `localhost:3000` itself
	// leading to a inifite loop, so don't accept the domain for shortening
	if !helpers.RemoveDomainError(body.URL) {

		log.Fatal("Nice Try")
		return

	}

	// enforce https
	// all url will be converted to https before storing in database
	body.URL = helpers.EnforceHTTP(body.URL)

	customShort := uuid.New().String()[:6]

	body.Expiry = 24

	// create a short

	var count int64

	err := initializers.DB.Table("response_urls").Count(&count).Error

	if err != nil {

		log.Fatal("Error fetching count", err)

	}

	if count >= 20000 {

		log.Fatal("Can't add more rows")
		return
	}

	short := models.ResponseURL{URL: body.URL, CustomShort: customShort, Expiry: body.Expiry}

	result := initializers.DB.Create(&short)

	if result.Error != nil {

		log.Fatal("Error creating the short url")
		return

	}

	// respond with the url, short, expiry in hours, calls remaining and time to reset
	resp := response{
		URL:         body.URL,
		CustomShort: "",
		Expiry:      body.Expiry,
	}

	resp.CustomShort = os.Getenv("DOMAIN") + "/" + customShort

	c.JSON(200, gin.H{

		"message": resp,
	})

}
