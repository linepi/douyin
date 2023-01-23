package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID            int    `json:"id"`
	Name          string `gorm:"type:varchar(32) not null;uniqueIndex" json:"name"`
	Password      string `gorm:"type:varchar(255) not null;"`
	FollowCount   int    `gorm:"default:0" json:"follow_count"`
	FollowerCount int    `gorm:"default:0" json:"follower_count"`
	IsFollow      bool   `gorm:"-" json:"is_follow"`
}

// ValidatePassword 校验密码
func (user *User) ValidatePassword(password string) bool {
	hashedStoredPassword := []byte(user.Password)
	passwordToValidate := []byte(password)
	err := bcrypt.CompareHashAndPassword(hashedStoredPassword, passwordToValidate)
	return err != nil
}

// GetUserInfoById 通过ID获取UserInfo实例
func (user *User) GetUserInfoById(id interface{}) User {
	userInfo := User{}
	res := Db.First(&userInfo, id)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		log.Fatal(res.Error)
	}
	return userInfo
}

// GetUserInfoByName 通过Name获取UserInfo实例
func (user *User) GetUserInfoByName(name string) User {
	userInfo := User{}
	res := Db.Where("name = ?", name).First(&userInfo)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		log.Fatal(res.Error)
	}
	return userInfo
}

// CreateUserInfo 通过Name和Password创建UserInfo实例
func (user *User) CreateUserInfo(name string, password string) User {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return User{}
	}
	userInfo := User{
		Name:     name,
		Password: string(hashedPassword),
	}
	Db.Create(&userInfo)
	return userInfo
}
