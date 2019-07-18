package apisite

import (
	"net/http"

	"../../utils"
)

type Item struct {
	ID int `json:"id"`
}

const (
	itemURL = "https://api.mercadolibre.com/items"
)

func (item *Item) Get() *utils.ApiError {

	if item.ID == 0 {
		return &utils.ApiError{
			Message: "El ID del item está vacío",
			Status:  http.StatusBadRequest,
		}
	}

	//res, err := http.Get(fmt.Sprint(itemURL, item.ID))
//	if e := hasError(err, http.StatusInternalServerError); e != nil {
//		return e
//	}
//
//	data, err := ioutil.ReadAll(res.Body)
//	if e := hasError(err, http.StatusInternalServerError); e != nil {
//		return e
//	}
//
//	var getitem Item
//	err = json.Unmarshal(data, &getitem)
//	if e := hasError(err, http.StatusInternalServerError); e != nil {
//		return e
//	}
//
	return nil
}
