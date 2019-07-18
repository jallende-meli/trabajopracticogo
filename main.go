package main

import (
	"./controllers/apisite"
	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

var (
	router = gin.Default()
)

func main () {
	router.GET("/user/:userID", controllers.GetUserFromAPI)
	router.GET("/country/:countryID", controllers.GetCountryFromAPI)
	router.GET("/site/:siteID", controllers.GetSiteFromAPI)
	router.GET("/result/:userID", controllers.GetResultFromAPI)
	router.Run(port)
}
