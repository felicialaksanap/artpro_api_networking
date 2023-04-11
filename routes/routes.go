package routes

import (
	"artpro_api_networking/controllers"
	"github.com/labstack/echo"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Selamat Datang di Echo")
	})

	e.POST("/addakunuser", controllers.SimpanAkunUser)
	e.GET("/akunuser", controllers.DataAkunUser)

	e.POST("/addprofileuser", controllers.SimpanProfileUser)
	e.GET("/profileuser", controllers.DataProfileUser)

	e.POST("/addverifikasidata", controllers.SimpanDataVerifikasi)
	e.GET("/alldataverifikasi", controllers.DataAllVerifikasi)
	e.GET("/dataverifikasiuser", controllers.DataVerifikasiUser)
	e.PUT("/editdataverifikasi", controllers.UpdateDataVerifikasi)

	return e
}
