package apisite

import (
	"../../utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)


type Country struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Locale string `json:"locale"`
	CurrencyID string `json:"currency_id"`
	DecimalSeparator string `json:"decimal_separator"`
	ThousandsSeparator string `json:"thousands_separator"`
	TimeZone string `json:"time_zone"`
	GeoInformation interface {} `json:"geo_information"`
	State interface{} `json:"states"`
}

func (country *Country) Get() *utils.ApiError {

	if country.ID == "" {
		return &utils.ApiError{
			Message: "El id de country es vacio",
			Status: http.StatusBadRequest,
		}
	}

	url := utils.UrlCountries + country.ID

	response, err := http.Get(url)
	if err != nil {
		return &utils.ApiError {
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}
	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}
	if err := json.Unmarshal(data, &country); err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	return nil
}
