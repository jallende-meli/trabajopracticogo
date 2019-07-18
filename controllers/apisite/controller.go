package controllers

import (
	"../../services/apisite"
	"../../utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	paramSiteID = "siteID"
	paramUserID = "userID"
	paramcountryID = "countryID"
	paramitemID = "itemID"
)

func GetUserFromAPI(c *gin.Context) {
	userID := c.Param(paramUserID)
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiError := &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		c.JSON(apiError.Status, apiError)
		return
	}

	user, apiError := apisite.GetUserFromAPI(id)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, user)
}


func GetCountryFromAPI(c *gin.Context) {
	countryID := c.Param(paramcountryID)
	country, apiError := apisite.GetCountryFromAPI(countryID)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, country)
}

func GetSiteFromAPI(c *gin.Context) {
	siteID := c.Param(paramSiteID)
	site, apiError := apisite.GetSiteFromAPI(siteID)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, site)
}

func GetResultFromAPI(c *gin.Context) {
	userID := c.Param(paramUserID)
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiError := &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		c.JSON(apiError.Status, apiError)
		return
	}

	result, apiError := apisite.GetResultFromAPI(id)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.JSON(http.StatusOK, result)
}