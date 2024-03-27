package model

import (
	"time"

	"my_gram/pkg/helper"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	DefaultColumn
	Username     string        `gorm:"not null" json:"username" valid:"required~Username is required"`
	Email        string        `gorm:"not null" json:"email" valid:"required~Email is required,email~Invalid email format"`
	Password     string        `gorm:"not null" json:"password" valid:"required~Password is required,minstringlength(6)~Password has to have minimum length of 6 characters"`
	DoB          time.Time     `gorm:"column:dob" json:"dob"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type UserSignUp struct {
	Username string    `json:"username" valid:"required~username is required"`
	Password string    `json:"password" valid:"required~password is required,minstringlength(6)~password must have a minimum length of 6 characters"`
	Email    string    `json:"email"`
	DoB      time.Time `json:"dob"`
}

type UserSignUpCreate struct {
	DefaultColumn
	Username string    `json:"username"`
	Email    string    `json:"email"`
	DoB      time.Time `json:"dob"`
}

type UserSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInOutput struct {
	Token string `json:"token"`
}

type UserUpdate struct {
	ID       uint64 `valid:"required~ID is required"`
	Username string `form:"username" valid:"required~username is required"`
	Email    string `form:"email"`
}

type UserCreateInputSwagger struct {
	Username string `form:"username"`
	Email    string `form:"email"`
}

type UserUpdateSwagger = UserCreateInputSwagger

type UserUpdateOutput struct {
	DefaultColumn
	Username string    `form:"username" valid:"required~username is required"`
	Email    string    `form:"email"`
	DoB      time.Time `json:"dob"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helper.GeneratePass(u.Password)
	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helper.GeneratePass(u.Password)
	err = nil
	return
}
