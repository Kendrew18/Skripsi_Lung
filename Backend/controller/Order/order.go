package order

import (
	order "kendrew/models/Order"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SaveOrderController(c echo.Context) error {
	id_pelanggan := c.FormValue("id_pelanggan")
	tanggal_pemesanan := c.FormValue("tanggal_pemesanan")
	no_order := c.FormValue("no_order")
	pembayaran := c.FormValue("pembayaran")
	down_payment := c.FormValue("down_payment")
	tanggal_pembayaran := c.FormValue("tanggal_pembayaran")
	catatan := c.FormValue("catatan")
	total_barang := c.FormValue("total_barang")
	harga_jual := c.FormValue("harga_jual")
	sub_total := c.FormValue("sub_total")
	id_ukuran := c.FormValue("id_ukuran")
	id_stock := c.FormValue("id_stock")
	jumlah := c.FormValue("jumlah")
	satuan := c.FormValue("satuan")
	id_sales := c.FormValue("id_sales")

	id_p, _ := strconv.Atoi(id_pelanggan)
	id_s, _ := strconv.Atoi(id_sales)
	result, err := order.SaveOrder(id_p, id_s, tanggal_pemesanan, no_order,
		pembayaran, down_payment, tanggal_pembayaran, catatan,
		total_barang, harga_jual, sub_total, id_ukuran, id_stock, jumlah, satuan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ShowHeaderOrderController(c echo.Context) error {

	id_sales := c.FormValue("id_sales")

	id_s, _ := strconv.Atoi(id_sales)

	result, err := order.ShowOrderHeader(id_s)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ShowDetailTransaksiController(c echo.Context) error {

	id_order := c.FormValue("id_order")

	id_o, _ := strconv.Atoi(id_order)

	result, err := order.ShowDetailTransaksi(id_o)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func ShowOrderHeader_AdminController(c echo.Context) error {

	result, err := order.ShowOrderHeader_Admin()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
