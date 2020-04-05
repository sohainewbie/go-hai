package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type BoxToken struct {
	AccessToken string `json:"accessToken"`
}

type UserContext struct {
	UserID   uint64 `json:"userId"`
	Exp      uint64 `json:"exp"`
	RoleName string `json:"roleName"`
}

var (
	EchoContext echo.Context
	superRole   = "admin"
	listRoles   = []string{"user"} // you can set another roles in here
)

func Authorizer(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			EchoContext = c
			req := c.Request()
			tokenString := ""
			if len(req.Header["Authorization"]) != 0 {
				tokenString = req.Header["Authorization"][0]
			} else {
				return ErrorResponse(c, errors.New("Authorization not found"))
			}
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(Config.SecretKey), nil
			})
			if err != nil {
				return ErrorResponse(c, err)
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

				roleFromJwt := strings.ToLower(claims["roleName"].(string))
				if len(roles) == 1 {
					roleFromRoutes := roles[0]

					if roleFromJwt != superRole {
						if !(intersectRoles(roleFromJwt, listRoles)) || (roleFromJwt != roleFromRoutes) {
							return UnauthorizedResponse(c)
						}
					}
				}

				if len(claims["roleName"].(string)) == 0 {
					return UnauthorizedResponse(c)
				}

				userContext := UserContext{
					Exp:      uint64(claims["exp"].(float64)),
					RoleName: claims["roleName"].(string),
				}
				userContext.UserID, _ = strconv.ParseUint(claims["userId"].(string), 10, 64)
				c.Set("user", userContext)
			} else {
				return UnauthorizedResponse(c)
			}

			return next(c)
		}
	}
}

func intersectRoles(role1 string, roles2 []string) bool {
	for _, role2 := range roles2 {
		if role1 == role2 {
			return true
		}
	}
	return false
}
