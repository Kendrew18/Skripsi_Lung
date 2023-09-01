package controller

import (
	"kendrew/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// login_user
func LoginUM(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	result, err := models.Login(username, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

// login admin
func LoginAdmin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	result, err := models.Login_admin(username, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
