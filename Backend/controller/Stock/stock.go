package stock

import (
	stock "kendrew/models/Stock"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SaveStockController(c echo.Context) error {
	nama_barang := c.FormValue("nama_barang")
	harga_barang := c.FormValue("harga_barang")
	jenis_barang := c.FormValue("jenis_barang")
	id_ukuran := c.FormValue("id_ukuran")
	jumlah_stock := c.FormValue("jumlah_stock")

	hg_b, _ := strconv.ParseInt(harga_barang, 10, 64)

	result, err := stock.SaveStock(nama_barang, hg_b, jenis_barang, id_ukuran,
		jumlah_stock)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ShowStock(c echo.Context) error {

	result, err := stock.ShowStock()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
