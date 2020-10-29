package request

import "github.com/ntttrang/t-selling-service/messagehandler"

// UserCreateReq user for create
type UserCreateReq struct {
	UserID      string `json:"userId" validate:"required"`
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	CountryCode string `json:"countryCode" validate:"required"`
}

//NewUserCreateReq UserCreateReq
func NewUserCreateReq() UserCreateReq {
	return UserCreateReq{}
}

// Validate validate req
func (u *UserCreateReq) Validate() *messagehandler.Mess {
	return ValidateFields(u)
}
