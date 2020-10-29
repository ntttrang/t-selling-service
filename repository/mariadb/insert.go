package mariadb

import (
	"github.com/ntttrang/t-selling-service/entity"
	"github.com/ntttrang/t-selling-service/messagehandler"
)

// InsertUser add new user
func (mdb *MariaDb) InsertUser(u entity.User) *messagehandler.Mess {
	var err *messagehandler.Mess
	e := mdb.db.Create(&u).Error
	if nil != e {
		err = messagehandler.W000001.GetMessage("User")
		return err
	}
	return err
}
