package mariadb

import (
	"github.com/ntttrang/t-selling-service/entity"
	"github.com/ntttrang/t-selling-service/messagehandler"
)

// DeleteUser delete a user
func (mdb *MariaDb) DeleteUser(userID string) *messagehandler.Mess {
	var err *messagehandler.Mess
	e := mdb.db.Where("user_id=?", userID).Delete(&entity.User{UserID: userID}).Error
	if nil != e {
		err = messagehandler.W000003.GetMessage("User")
		return err
	}
	return err
}
