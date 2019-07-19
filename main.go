package main

import (
	"./controllers/apisite"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

const (
	port = ":8080"
)

var (
	router = gin.Default()
)

var limiter = rate.NewLimiter(2, 5)

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	router.GET("/user/:userID", controllers.GetUserFromAPI)
	router.GET("/country/:countryID", controllers.GetCountryFromAPI)
	router.GET("/site/:siteID", controllers.GetSiteFromAPI)
	router.GET("/result/:userID", controllers.GetResultFromAPI)
	http.ListenAndServe(port, limit(router))
	//router.Run(port)
}
