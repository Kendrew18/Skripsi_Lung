package routes

import (
	"kendrew/controller"
	order "kendrew/controller/Order"
	pelanggan "kendrew/controller/Pelanggan"
	retur "kendrew/controller/Retur"
	salesman "kendrew/controller/Salesman"
	stock "kendrew/controller/Stock"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Skripsi-LUNG")
	})
	sales := e.Group("/sales")
	stk := e.Group("/stk")
	plgn := e.Group("/plgn")
	ord := e.Group("/ord")
	rtr := e.Group("/rtr")

	//login_admin
	e.GET("/login", controller.LoginUM)
	e.GET("/login-admin", controller.LoginAdmin)

	//sales
	sales.POST("/savesales", salesman.SaveSalesController)
	sales.GET("/showSales", salesman.ShowSales)
	sales.PUT("/updateSalesLocation", salesman.UpdateLocation)

	//Stock
	stk.POST("/savestock", stock.SaveStockController)
	stk.GET("/showstock", stock.ShowStock)

	//pelanggan
	plgn.POST("/savepelanggan", pelanggan.SavePelangganController)
	plgn.GET("/showpelanggan", pelanggan.ShowPelangganController)

	//order
	ord.POST("/saveorder", order.SaveOrderController)
	ord.GET("/showheaderorder", order.ShowHeaderOrderController)
	ord.GET("/showorderdetail", order.ShowDetailTransaksiController)
	ord.GET("/showheaderorderadmin", order.ShowOrderHeader_AdminController)
	ord.PUT("/updatetanggalpengiriman", order.UpdateTanggalPengirimanController)

	//retur
	rtr.POST("/saveretur", retur.SaveReturController)
	rtr.GET("/showheaderretur", retur.ShowHeaderReturController)
	rtr.GET("/showreturdetail", retur.ShowDetailReturController)
	rtr.GET("/showheaderreturadmin", retur.ShowReturHeader_AdminController)

	return e
}
