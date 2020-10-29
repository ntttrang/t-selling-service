package controller

import (
	"fmt"

	"github.com/kataras/golog"
	"github.com/labstack/echo"
	"github.com/ntttrang/t-selling-service/service"
	"github.com/ntttrang/t-selling-service/startup/config"
)

type (
	//HTTPManager HTTPManager
	HTTPManager interface {
		RegisterRouter()
		Start()

		// TODO: Add controller name
		UserController
	}

	// UserController UserController
	UserController interface {
		GetUsers() interface{}
		InsertUser() interface{}
		UpdateUser() interface{}
		DeleteUser() interface{}
	}
)

// EchoController xxx
type EchoController struct {
	service service.Service
	config  config.ServerConfig
	e       *echo.Echo
}

//NewEchoHTTPManager new echo HTTP Manager
func NewEchoHTTPManager() HTTPManager {
	return &EchoController{
		service: service.NewBaseService(),
		config:  config.GetServerConf(),
		e:       echo.New(),
	}
}

// Start start server
func (c *EchoController) Start() {
	golog.Info(c.e.Start(fmt.Sprintf("%v:%v", c.config.Host, c.config.Port)))
}
