package mariadb

import (
	"github.com/ntttrang/t-selling-service/messagehandler"
	"github.com/ntttrang/t-selling-service/request"

	"github.com/ntttrang/t-selling-service/entity"
)

// SelectUser select user information
func (mdb *MariaDb) SelectUser(req request.UserSearch) ([]entity.User, *messagehandler.Mess) {
	var userList []entity.User
	var err *messagehandler.Mess

	userID := req.UserID
	firstName := req.FirstName
	lastName := req.LastName
	countryCode := req.CountryCode
	skip := req.Skip
	pageSize := req.PageSize
	sortBy := req.SortBy
	direction := req.Direction

	DB := mdb.db.Select("COUNT(*) OVER() AS total, users.*").Table("users")
	if "" != userID {
		DB = DB.Where("user_id = ?", userID)
	}

	if "" != firstName {
		DB = DB.Where("first_name LIKE ?", "%"+firstName+"%")
	}

	if "" != lastName {
		DB = DB.Where("last_name LIKE ?", "%"+lastName+"%")
	}

	if "" != countryCode {
		DB = DB.Where("country_code = ?", countryCode)
	}

	if 0 != skip {
		DB = DB.Offset(skip)
	}

	if 0 != pageSize {
		DB = DB.Limit(pageSize)
	}

	if "" != sortBy && "" != direction {
		orderBy := sortBy + " " + direction
		DB = DB.Order(orderBy, true)
	}
	e := DB.Find(&userList).Error
	if nil != e {
		err = messagehandler.W000004.GetMessage()
		return userList, err
	}
	return userList, err
}
