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

	// ==== AKUN USER ====
	e.POST("/addakunuser", controllers.SimpanAkunUser)
	e.GET("/akunuser", controllers.DataAkunUser)
	// === END ====

	// ==== PROFILE USER ====
	e.POST("/addprofileuser", controllers.SimpanProfileUser)
	e.GET("/profileuser", controllers.DataProfileUser)
	// ==== END ====

	// ===== VERIFIKASI DATA KTP =====
	e.POST("/addverifikasidata", controllers.SimpanDataVerifikasi)
	e.GET("/alldataverifikasi", controllers.DataAllVerifikasi)
	e.GET("/dataverifikasiuser", controllers.DataVerifikasiUser)
	e.PUT("/editdataverifikasi", controllers.UpdateDataVerifikasi)
	// ==== END ====

	// ==== ALAMAT DOMISILI USER =====
	e.POST("/adduserdomisili", controllers.SimpanDomisiliUser)
	e.GET("/datauserdomisili", controllers.DataUserDomisili)
	// ===== END ======

	// ==== DETAIL PROFILE ART =====
	e.POST("/addprofileart", controllers.SimpanDetailProfileART)
	e.GET("/datauserdetailprofileart", controllers.DataUserDetailProfileART)
	// ==== END ====

	// ===== DETAIL KERJA ART ======
	e.POST("/addkerjaart", controllers.SimpanDetailKerjaART)
	e.GET("/alldatadetailkerjaart", controllers.DataAllDetailKerjaART)
	e.GET("/datakerjaperkategori", controllers.DataListKerjaPerKategori)
	// ==== END ====

	// ===== LOWONGAN KERJA =====
	e.POST("/addlowongankerja", controllers.SimpanLowonganKerja)
	// === END ===

	return e
}
