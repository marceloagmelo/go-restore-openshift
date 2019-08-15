package router

import (
	"github.com/labstack/echo"
	"gitlab.produbanbr.corp/paas-brasil/go-restore-openshift/controller"
)

//App é uma instancia de App
var App *echo.Echo

func init() {
	App = echo.New()

	// A página inicial da aplicação
	App.GET("/", controller.Home)

	api := App.Group("/v1")
	api.POST("/restore-openshift", controller.Restore)

}
