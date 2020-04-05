package query

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"strings"
	"time"

	"github.com/sohainewbie/go-hai/schemas"
	"github.com/sohainewbie/go-hai/utils"
)

func GetUser(filter schemas.ParamSearch) (usersData schemas.UsersSchema) {
	db := utils.GetInstanceMysqlDb().Model("users")
	if len(filter.Email) > 0 {
		db = db.Where("email = ?", filter.Email)
	}
	if filter.ID > 0 {
		db = db.Where("id = ?", filter.ID)
	}

	db = db.Where("deleted_at is null")
	db = db.First(&usersData)
	return
}

func UserLogin(params interface{}) (response schemas.UserAuthResponse, err error) {
	parameter := params.(*schemas.PayloadLogin)

	//try to find data from database
	checkUser := utils.GetInstanceMysqlDb().Table("users").Select(`users.id, roles.type as roleName, users.name, email, password`)
	checkUser = checkUser.Joins(`JOIN users_role ON users_role.user_id = users.id`)
	checkUser = checkUser.Joins(`JOIN roles ON roles.id = users_role.role_id`)

	if strings.Contains(parameter.Email, "@") {
		checkUser = checkUser.Where("email = ? and users.deleted_at is null", parameter.Email).Find(&response)
		if len(response.Email) == 0 {
			if len(response.Email) == 0 {
				err = fmt.Errorf("Your Email is incorect.")
				return
			}
		}
	}

	//usersData.Password base on db connection from database
	if !utils.CheckPasswordHash(parameter.Password, response.Password) {
		err = fmt.Errorf("Your Password is incorect.")
		return
	}

	var defaultTime time.Duration = 30 // in minute
	//set token expired
	response.Exp = time.Now().Add(time.Minute * defaultTime).Unix()

	//create Token
	token, err := CreateToken(response)
	if err != nil {
		return
	}

	response.Token = token
	response.AuthType = "web"
	return
}

func CreateUser(params interface{}) (usersData schemas.UsersSchema, err error) {
	parameter := params.(*schemas.UsersSchema)

	filter := schemas.ParamSearch{}
	filter.Email = parameter.Email
	checkUserRegister := GetUser(filter)
	if len(checkUserRegister.Email) > 0 || checkUserRegister.Email == parameter.Email {
		err = fmt.Errorf("Email exist")
		return
	}

	if len(checkUserRegister.PhoneNumber) > 0 || checkUserRegister.PhoneNumber == parameter.PhoneNumber {
		err = fmt.Errorf("Phone number exist")
		return
	}

	//if the payload have field password
	if len(parameter.Password) > 0 {
		hashPassword, errHash := utils.HashPassword(parameter.Password)
		if errHash != nil {
			err = errHash
			return
		}
		usersData.Password = hashPassword
	}

	//set payload before save to database
	usersData.Name = parameter.Name
	usersData.Email = parameter.Email
	usersData.PhoneNumber = parameter.PhoneNumber

	tx := utils.GetInstanceMysqlDb().Begin()
	if err = tx.Table("users").Create(&usersData).Error; err != nil {
		tx.Rollback()
		return
	}

	// create relation with table role_user
	usersRoleSchema := schemas.UsersRoleSchema{
		RoleID: 1, // user
		UserID: usersData.ID,
	}

	if err = tx.Table("users_role").Create(&usersRoleSchema).Error; err != nil {
		tx.Rollback()
		return
	}

	//commit transaction
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return
	}

	return usersData, nil
}

func CreateToken(authData schemas.UserAuthResponse) (t string, err error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = authData.ID
	claims["roleName"] = authData.RoleName
	claims["exp"] = authData.Exp

	// Generate encoded token and send it as response.
	t, err = token.SignedString([]byte(utils.Config.SecretKey))
	return
}
