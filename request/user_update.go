package request

import "github.com/ntttrang/t-selling-service/messagehandler"

// UserUpdateReq user for update
type UserUpdateReq struct {
	UserID      string `json:"userId" validate:"required"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	CountryCode string `json:"countryCode"`
}

//NewUserUpdateReq NewUserUpdateReq
func NewUserUpdateReq() UserUpdateReq {
	return UserUpdateReq{}
}

// Validate validate req
func (u *UserUpdateReq) Validate() *messagehandler.Mess {
	return ValidateFields(u)
}
