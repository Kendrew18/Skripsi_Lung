package controller

import (
	salesman "kendrew/models/Salesman"
	"net/http"

	"github.com/labstack/echo/v4"
)

// save_sales
func SaveSalesController(c echo.Context) error {
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	nomor_hp := c.FormValue("nomor_hp")
	bank := c.FormValue("bank")
	no_rekening := c.FormValue("no_rekening")
	username := c.FormValue("username")
	password := c.FormValue("password")

	result, err := salesman.SaveSales(nama, alamat, nomor_hp, bank,
		no_rekening, username, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

// Show_sales
func ShowSales(c echo.Context) error {

	result, err := salesman.ShowSales()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateLocation(c echo.Context) error {
	longitude := c.FormValue("longitude")
	latitude := c.FormValue("latitude")
	idUser := c.FormValue("id_user")

	result, err := salesman.UpdateUserLocation(longitude, latitude, idUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
