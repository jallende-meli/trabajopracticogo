package apisite

import (
	"../../utils"
)

type Result struct {
	User *User
	Country *Country
	Site *Site
}

type Ojbect interface {
	Get() *utils.ApiError
}