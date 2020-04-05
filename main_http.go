package main

import (
	"fmt"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/sohainewbie/go-hai/controller"
	"github.com/sohainewbie/go-hai/utils"
)

// HTTPServeMain - function for handling http
func HTTPServeMain() {
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.CORS())

	e.HideBanner = true
	banner := `
	  ________ ________              ___ ___        .__ 
	 /  _____/ \_____  \            /   |   \_____  |__|
	/   \  ___  /   |   \   ______ /    ~    \__  \ |  |
	\    \_\  \/    |    \ /_____/ \    Y    // __ \|  |
	 \______  /\_______  /          \___|_  /(____  /__|
	        \/         \/                 \/      \/    
		Version : %s | Port : %s`
	fmt.Printf(banner+"\n", utils.Config.Service.Version, utils.Config.Service.Port)
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {

		startTime := time.Now()
		startDate := startTime.Format("2006-01-02 15:04:05")

		req := c.Request()
		res := c.Response()
		reqHeader := ""
		for headerName, headerValue := range req.Header {
			reqHeader += fmt.Sprintf("%s=%s \n", headerName, headerValue)
		}
		logString := "----------------------------------------------------------------------------\n"
		logString += fmt.Sprintf("DATE=%s\nREQUEST-METHOD=%s\nREQUEST-URL=%s\n", startDate, req.Method, req.RequestURI)
		logString += fmt.Sprintf("REQUEST-HEADER=\n%s\n", reqHeader)
		if len(reqBody) != 0 {
			logString += fmt.Sprintf("REQUEST-BODY=%s\n", reqBody)
		}
		logString += fmt.Sprintf("RESPONSE-STATUSCODE=%v\nRESPONSE-BODY=%s\n", res.Status, resBody)
		logString += "----------------------------------------------------------------------------\n"
		fmt.Println(logString)
	}))

	e.GET("/", controller.Index)
	e.POST("/register", controller.Register)
	e.POST("/login", controller.Login)

	v1 := e.Group("/v1")
	v1.GET("/user/:id", controller.DataUser, utils.Authorizer("admin"))
	v1.GET("/profile", controller.Profle, utils.Authorizer("user"))

	e.Logger.Debug(e.Start(":" + utils.Config.Service.Port))
}
