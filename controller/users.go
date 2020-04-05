package controller

import (
	"github.com/labstack/echo"
	"github.com/sohainewbie/go-hai/query"
	"github.com/sohainewbie/go-hai/schemas"
	"github.com/sohainewbie/go-hai/utils"
	"strconv"
)

//Register - function for register and validation
func Register(c echo.Context) error {
	// parsing
	request := new(schemas.UsersSchema)
	err, data := utils.ParsingValidate(c, request)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}

	response, err := query.CreateUser(data)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, response)
}

func Login(c echo.Context) error {

	request := new(schemas.PayloadLogin)
	err, data := utils.ParsingValidate(c, request)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}

	response, err := query.UserLogin(data)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}

	return utils.SuccessResponse(c, response)
}

func Profle(c echo.Context) error {
	// parsing
	userContext := c.Get("user").(utils.UserContext)
	filter := schemas.ParamSearch{
		ID: userContext.UserID,
	}
	response := query.GetUser(filter)
	return utils.SuccessResponse(c, response)
}

func DataUser(c echo.Context) error {
	userId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	filter := schemas.ParamSearch{
		ID: userId,
	}
	response := query.GetUser(filter)
	return utils.SuccessResponse(c, response)
}
