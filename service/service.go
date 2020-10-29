package service

import (
	"github.com/ntttrang/t-selling-service/entity"
	"github.com/ntttrang/t-selling-service/messagehandler"
	"github.com/ntttrang/t-selling-service/repository"
	"github.com/ntttrang/t-selling-service/repository/mariadb"
	"github.com/ntttrang/t-selling-service/request"
)

//BaseService define base service
type BaseService struct {
	repo repository.Repository
}

//NewBaseService init base service
func NewBaseService() Service {
	return &BaseService{
		repo: mariadb.NewRepoMariaDb(),
	}
}

type (
	// Service Service
	Service interface {
		UserService
	}

	//UserService UserService
	UserService interface {
		GetUsers(request.UserSearch) ([]entity.User, *messagehandler.Mess)
		InsertUser(request.UserCreateReq) *messagehandler.Mess
		UpdateUser(request.UserUpdateReq) *messagehandler.Mess
		DeleteUser(string) *messagehandler.Mess
	}
)
