package schemas

import (
	"github.com/sohainewbie/go-hai/utils"
)

type UserAuthResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	RoleName string `gorm:"column:roleName" json:"roleName"`
	AuthType string `json:"authType"`
	Exp      int64  `json:"exp"`
	Password string `json:"-"`
	Token    string `json:"token"`
}

type PayloadLogin struct {
	Email    string `json:"email" valid:"required"`
	Password string `json:"password" valid:"required"`
}

type UsersSchema struct {
	utils.BaseModel
	Name        string `json:"name" valid:"required"`
	Email       string `json:"email" valid:"required"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password" valid:"required"`
}

func (UsersSchema) TableName() string {
	return "users"
}

type UsersRoleSchema struct {
	utils.BaseModel
	UserID uint64 `json:"userId"`
	RoleID uint64 `json:"roleId"`
}

func (UsersRoleSchema) TableName() string {
	return "users_role"
}
