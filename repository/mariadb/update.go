package mariadb

import (
	"github.com/ntttrang/t-selling-service/entity"
	"github.com/ntttrang/t-selling-service/messagehandler"
)

// UpdateUser update user
func (mdb *MariaDb) UpdateUser(u entity.User) *messagehandler.Mess {
	var err *messagehandler.Mess
	userID := u.UserID
	// Not allow to update `userId`, so set userId is empty
	u.UserID = ""
	e := mdb.db.Model(u).Where("user_id=?", userID).Updates(&u).Error
	if nil != e {
		err = messagehandler.W000002.GetMessage("User")
	}
	return err
}
