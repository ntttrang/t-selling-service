package mariadb

import (
	"github.com/jinzhu/gorm"
	"github.com/ntttrang/t-selling-service/repository"
	"github.com/ntttrang/t-selling-service/startup/database"
)

// MariaDb mariadb
type MariaDb struct {
	db *gorm.DB
}

// NewRepoMariaDb init mariadb
func NewRepoMariaDb() repository.Repository {
	return &MariaDb{db: database.GetMariaDb()}
}
