package controller

import (
	"github.com/go-hai/utils"
	"github.com/labstack/echo"
)

//Index - first page for make sure this service is running well
func Index(c echo.Context) error {
	response := &struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}{
		Name:    utils.Config.Service.Name,
		Version: utils.Config.Service.Version,
	}
	return utils.SuccessResponse(c, response)
}
