package retur

import (
	retur "kendrew/models/Retur"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SaveReturController(c echo.Context) error {
	tanggal_retur := c.FormValue("tanggal_retur")
	no_order := c.FormValue("no_order")
	total_barang := c.FormValue("total_barang")
	harga_jual := c.FormValue("harga_jual")
	sub_total := c.FormValue("sub_total")
	id_ukuran := c.FormValue("id_ukuran")
	id_stock := c.FormValue("id_stock")
	jumlah := c.FormValue("jumlah")
	satuan := c.FormValue("satuan")
	id_sales := c.FormValue("id_sales")

	id_s, _ := strconv.Atoi(id_sales)
	result, err := retur.SaveRetur(no_order,
		tanggal_retur, total_barang, harga_jual,
		sub_total, id_stock, id_ukuran, jumlah, satuan,
		id_s)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ShowHeaderReturController(c echo.Context) error {

	id_sales := c.FormValue("id_sales")

	id_s, _ := strconv.Atoi(id_sales)

	result, err := retur.ShowReturHeader(id_s)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ShowDetailReturController(c echo.Context) error {

	id_retur := c.FormValue("id_retur")

	id_R, _ := strconv.Atoi(id_retur)

	result, err := retur.ShowDetailRetur(id_R)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
