package repository

import (
	"github.com/ntttrang/t-selling-service/entity"
	"github.com/ntttrang/t-selling-service/messagehandler"
	"github.com/ntttrang/t-selling-service/request"
)

type (
	// Repository is the place we perform database operatons
	Repository interface {
		UserRepo
	}

	//UserRepo define user repo
	//  haha
	//  ddada
	UserRepo interface {
		SelectUser(request.UserSearch) ([]entity.User, *messagehandler.Mess)
		InsertUser(entity.User) *messagehandler.Mess
		UpdateUser(entity.User) *messagehandler.Mess
		DeleteUser(string) *messagehandler.Mess
	}
)
