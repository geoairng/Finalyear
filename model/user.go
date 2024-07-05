package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          uint       `gorm:"primary_key;auto_increment"`
	CreatedAt   *time.Time `gorm:"default:CURRENT_TIMESTAMP" example:"2023-09-02"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime" example:"2023-09-02"`
	ComapnyName string     `json:"Companyname" example:"xyz"`
	Email       string     `gorm:"not null;unique" json:"Email" example:"comapany@gmail.com"`
	Password    string     `json:"-"`
}

func (user *User) VerifyHash(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
