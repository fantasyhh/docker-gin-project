package models

import "github.com/jinzhu/gorm"

// Auth struct for user Model
type Auth struct {
	ID       int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Username string `json:"username"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Email    string `json:"email"`
}

// TableName 代表Auth struct对应的数据库表名
func (Auth) TableName() string {
	return "auth"
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) (bool, error) {
	var auth Auth
	err := db.Select("id").Where(&Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}
