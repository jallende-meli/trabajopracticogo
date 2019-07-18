package apisite

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"../../utils"
)

type Site struct {
	Name string `json:"name"`
	ID string    `json:"id"`
	CountryID string `json:"country_id"`
	DefaultCurrencyID string `json:"default_currency_id"`
}

func (site *Site) Get() *utils.ApiError {

	if site.ID == "" {
		return &utils.ApiError{
			Message: "El ID del site está vacío",
			Status:  http.StatusBadRequest,
		}
	}
	url := utils.UrlSites + site.ID
	res, err := http.Get(url)
	if err != nil {
		return &utils.ApiError {
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &site); err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	return nil
}

/////////
//
//func GetAll() ([]Site, *utils.ApiError) {
//
//	res, err := http.Get(siteURL)
//	if e := hasError(err, http.StatusInternalServerError); e != nil {
//		return nil, e
//	}
//
//	data, err := ioutil.ReadAll(res.Body)
//	if e := hasError(err, http.StatusInternalServerError); e != nil {
//		return nil, e
//	}
//
//	var dataSites []Site
//
//	err = json.Unmarshal(data, &dataSites)
//	if e := hasError(err, http.StatusInternalServerError); e != nil {
//		return nil, e
//	}
//
//	return dataSites, nil
//}
