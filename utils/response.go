package utils

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

// ServerHeader - function for serving custom header
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "get-go/1.0 (arch64)")
		return next(c)
	}
}

func SuccessResponseMap(ctx echo.Context, data map[string]interface{}) error {

	responseData := map[string]interface{}{
		"success": true,
		"data":    data,
	}

	return ctx.JSON(http.StatusOK, responseData)
}

func SuccessResponseWithMessage(ctx echo.Context, message string, data interface{}) error {

	responseData := map[string]interface{}{
		"success": true,
		"message": message,
		"data":    data,
	}

	return ctx.JSON(http.StatusOK, responseData)
}

func SuccessResponse(ctx echo.Context, data interface{}) error {

	responseData := map[string]interface{}{
		"success": true,
		"data":    data,
	}

	return ctx.JSON(http.StatusOK, responseData)
}

func ErrorResponse(ctx echo.Context, err error) error {

	responseData := map[string]interface{}{
		"success": false,
		"error":   err.Error(),
	}

	return ctx.JSON(http.StatusBadRequest, responseData)
}

func UnauthorizedResponse(ctx echo.Context) error {
	err := errors.New("Unauthorized")

	responseData := map[string]interface{}{
		"success": false,
		"message": err.Error(),
	}

	return ctx.JSON(http.StatusUnauthorized, responseData)
}

func RestrictedAccessResponse(ctx echo.Context, err error) error {

	responseData := map[string]interface{}{
		"success": false,
		"message": err.Error(),
	}

	return ctx.JSON(http.StatusUnauthorized, responseData)
}
