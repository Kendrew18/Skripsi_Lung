package routes

import (
	"kendrew/controller"
	salesman "kendrew/controller/Salesman"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Project-NDL")
	})
	//for get data
	e.GET("/login", controller.LoginUM)
	e.GET("/showSales", salesman.ShowSales)

	return e
}
