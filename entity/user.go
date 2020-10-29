package entity

import (
	"time"
)

//User user entity
type User struct {
	ID          int        `gorm:"column:id;primary_key" json:"id"`
	UserID      string     `gorm:"column:user_id" json:"userId"` // unique
	FirstName   string     `gorm:"column:first_name" json:"firstName"`
	LastName    string     `gorm:"column:last_name" json:"lastName"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"createdAt"`
	CreatedBy   string     `gorm:"column:created_by" json:"createdBy"`
	UpdateAt    *time.Time `gorm:"column:update_at" json:"updateAt"`
	UpdateBy    string     `gorm:"column:update_by" json:"updateBy"`
	CountryCode int        `gorm:"column:country_code" json:"countryCode"`
}

// TableName define table name
func (User) TableName() string {
	return "users"
}
