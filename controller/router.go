package controller

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// RegisterRouter to register the endpoints to `HttpManager`.
func (c *EchoController) RegisterRouter() {
	c.e.Use(middleware.Recover())
	c.e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		}))

	noNeedAuth := c.e.Group("/api/v1/admin")
	noNeedAuth.GET("/users", c.GetUsers().(echo.HandlerFunc))
	noNeedAuth.POST("/users", c.InsertUser().(echo.HandlerFunc))
	noNeedAuth.PUT("/users", c.UpdateUser().(echo.HandlerFunc))
	noNeedAuth.DELETE("/users", c.DeleteUser().(echo.HandlerFunc))

}
