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

	// ===== FOTO USER (PROFILE / VERIFIKASI) =====
	e.GET("/getimage", controllers.GetPhoto)
	e.POST("/uploadimage", controllers.UploadFoto)

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
	//e.GET("/datakerjaperkategori", controllers.DataListKerjaPerKategori)
	e.GET("/datauserdetailkerjaart", controllers.DataUserDetailKerjaART)
	// ==== END ====

	// ===== LOWONGAN KERJA =====
	e.POST("/addlowongankerja", controllers.SimpanLowonganKerja)
	e.GET("/alldatalowongankerja", controllers.DataAllLowonganKerja)
	e.GET("/datadetaillowongankerja", controllers.DataDetailLowonganKerja)
	e.POST("/addlokerselesai", controllers.SimpanLowonganKerjaSelesai)
	e.DELETE("/deletelowongankerja", controllers.DeleteDetailLowonganKerja)
	// === END ===

	// ===== KONTAK USER =====
	e.POST("/addkontakuser", controllers.SimpanKontakUser)
	e.GET("/datalistkontakbymajikan", controllers.DataListKontakByMajikan)
	e.GET("/datalistkontakbyart", controllers.DataListKontakByART)
	// ==== END ====

	// ===== PENILAIAN =====
	e.POST("/addpenilaian", controllers.SimpanPenilaian)
	e.GET("/datapenilaianart", controllers.DataPenilaianART)
	// === END ===

	// ===== SERTIFIKAT PELATIHAN =====
	e.POST("/addsertifpath", controllers.SimpanSertifikatPelatihan)
	e.GET("/datasertifpelatihanuser", controllers.DataSertifPelatihanUser)
	// ===== END =====

	return e
}
