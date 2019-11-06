package utils

import (
	"errors"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type BoxToken struct {
	AccessToken string `json:"accessToken"`
}

type UserContext struct {
	Exp         uint64 `json:"exp"`
}

var EchoContext echo.Context

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
				// do something
				userContext := UserContext{
									Exp:    uint64(claims["exp"].(float64)),
								}
				c.Set("user", userContext)
			} else {
				return UnauthorizedResponse(c)
			}

			return next(c)
		}
	}
}
