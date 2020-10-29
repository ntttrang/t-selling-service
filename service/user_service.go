package service

import (
	"strconv"

	"github.com/ntttrang/t-selling-service/entity"
	"github.com/ntttrang/t-selling-service/messagehandler"
	"github.com/ntttrang/t-selling-service/request"
)

// GetUsers GetUsers
func (s *BaseService) GetUsers(req request.UserSearch) ([]entity.User, *messagehandler.Mess) {
	var users []entity.User
	var err *messagehandler.Mess
	users, err = s.repo.SelectUser(req)
	return users, err
}

// InsertUser insert a new user
func (s *BaseService) InsertUser(req request.UserCreateReq) *messagehandler.Mess {
	var err *messagehandler.Mess
	// Validate param
	err = req.Validate()
	if nil != err {
		return err
	}
	userSearch := request.UserSearch{UserID: req.UserID}
	userList, err := s.repo.SelectUser(userSearch)
	if nil != err {
		return err
	}
	// userId already existed
	if len(userList) != 0 {
		err = messagehandler.W000009.GetMessage("userId")
		return err
	}

	// Do the logic
	countryCode, e := strconv.Atoi(req.CountryCode)
	if nil != e {
		err = messagehandler.W000006.GetMessage("countryCode")
		return err
	}
	u := entity.User{
		UserID:      req.UserID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		CountryCode: countryCode,
	}
	err = s.repo.InsertUser(u)
	return err

}

//UpdateUser update a user
func (s *BaseService) UpdateUser(req request.UserUpdateReq) *messagehandler.Mess {
	var err *messagehandler.Mess
	// Validate param
	err = req.Validate()
	if nil != err {
		return err
	}
	userSearch := request.UserSearch{UserID: req.UserID}
	userList, err := s.repo.SelectUser(userSearch)
	if nil != err {
		return err
	}
	// Not found userId
	if len(userList) == 0 {
		err = messagehandler.W000008.GetMessage("userId")
		return err
	}

	countryCode, e := strconv.Atoi(req.CountryCode)
	if nil != e {
		err = messagehandler.W000006.GetMessage("countryCode")
		return err
	}
	u := entity.User{
		UserID:      req.UserID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		CountryCode: countryCode,
	}
	err = s.repo.UpdateUser(u)
	return err
}

// DeleteUser delete a user by primary key
func (s *BaseService) DeleteUser(userID string) *messagehandler.Mess {
	var err *messagehandler.Mess
	userSearch := request.UserSearch{UserID: userID}
	userList, err := s.repo.SelectUser(userSearch)
	if nil != err {
		return err
	}
	// Not found userId
	if len(userList) == 0 {
		err = messagehandler.W000008.GetMessage("userId")
		return err
	}
	err = s.repo.DeleteUser(userID)
	return err
}
