package pelanggan

import (
	pelanggan "kendrew/models/Pelanggan"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SavePelangganController(c echo.Context) error {
	nama_toko := c.FormValue("nama")
	no_telp := c.FormValue("no_telp")
	alamat := c.FormValue("alamat")
	kota := c.FormValue("kota")
	provinsi := c.FormValue("provinsi")
	nama_penanggungjawab := c.FormValue("nama_penanggungjawab")

	result, err := pelanggan.SavePelanggan(nama_toko, no_telp, alamat, kota,
		provinsi, nama_penanggungjawab)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ShowPelangganController(c echo.Context) error {

	result, err := pelanggan.ShowPelanggan()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
