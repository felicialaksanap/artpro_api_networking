package routes

import (
	"artpro_api_networking/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Selamat Datang di Echo")
	})

	// ===== FOTO USER (PROFILE / VERIFIKASI) =====
	e.GET("/getimage", controllers.GetPhoto)
	e.POST("/uploadimage", controllers.UploadFoto)

	// ==== AKUN USER ====
	e.POST("/addakunuser", controllers.SimpanAkunUser)
	e.GET("/akunuser", controllers.DataAkunUser)
	e.PUT("/editemailuser", controllers.UpdateEmailUser)
	e.PUT("/editpassuser", controllers.UpdatePassUser)
	// === END ====

	// ==== PROFILE USER ====
	e.POST("/addprofileuser", controllers.SimpanProfileUser)
	e.GET("/profileuser", controllers.DataProfileUser)
	e.PUT("/editprofileuser", controllers.UpdateProfileUser)
	// ==== END ====

	// ===== VERIFIKASI DATA KTP =====
	e.POST("/addverifikasidata", controllers.SimpanDataVerifikasi)
	e.GET("/alldataverifikasi", controllers.DataAllVerifikasi)
	e.PUT("/editdataverifikasi", controllers.UpdateDataVerifikasi)
	// ==== END ====

	// ==== ALAMAT DOMISILI USER =====
	e.POST("/adduserdomisili", controllers.SimpanDomisiliUser)
	e.GET("/datauserdomisili", controllers.DataUserDomisili)
	e.PUT("/edituserdomisili", controllers.UpdateUserDomisili)
	e.GET("/getlonglatuser", controllers.DataLongLat)
	// ===== END ======

	// ==== DETAIL PROFILE ART =====
	e.POST("/addprofileart", controllers.SimpanDetailProfileART)
	e.GET("/datauserdetailprofileart", controllers.DataUserDetailProfileART)
	e.PUT("/edituserdetailprofileart", controllers.UpdateUserDetailProfileART)
	// ==== END ====

	// ===== DETAIL KERJA ART ======
	e.POST("/addkerjaart", controllers.SimpanDetailKerjaART)
	e.GET("/alldatadetailkerjaart", controllers.DataAllDetailKerjaART)
	e.GET("/dataartbykategori", controllers.DataARTbyKategori)
	e.GET("/dataartbyfk", controllers.DataARTbyFK)
	e.GET("/makeandcopytable", controllers.CreateAndInsertTableTemp)
	e.PUT("updatejarak", controllers.UpdateJarak)
	e.GET("/datauserdetailkerjaart", controllers.DataUserDetailKerjaART)
	e.PUT("/edituserdetailkerjaart", controllers.UpdateUserDetailKerja)
	// ==== END ====

	// ===== LOWONGAN KERJA =====
	e.POST("/addlowongankerja", controllers.SimpanLowonganKerja)
	e.GET("/alldatalowongankerja", controllers.DataAllLowonganKerja)
	e.GET("/datalokerperuser", controllers.DataLowonganKerjaperUser)
	e.PUT("/editstatusloker", controllers.UpdateStatusLoker)
	e.PUT("/editlowongankerja", controllers.UpdateLowonganKerja)
	e.GET("/getlokerfilter", controllers.DataLokerbyFilter)
	e.GET("/createcopytable", controllers.CreateAndCopyTable)
	e.PUT("/updatejarakloker", controllers.UpdateStatusJarakLoker)
	// === END ===

	// ===== KONTAK USER =====     // Majikan melakukan call ke ART
	e.POST("/addkontakuser", controllers.SimpanKontakUser)
	e.GET("/alldatakontakart", controllers.DataKontakART)
	e.GET("/alldatakontakmajikan", controllers.DataKontakMajikan)
	e.GET("/getinfokontak", controllers.KontakART)
	// ==== END ====

	// ===== PENILAIAN =====
	e.POST("/addpenilaian", controllers.SimpanPenilaian)
	e.GET("/dataratapenilaian", controllers.RataPenilaianART)
	e.GET("/datareviewmajikan", controllers.DataReviewMajikan)

	// === END ===

	// ===== SERTIFIKAT PELATIHAN =====
	//e.POST("/addsertifpath", controllers.SimpanSertifikatPelatihan)
	//e.GET("/datasertifpelatihanuser", controllers.DataSertifPelatihanUser)
	// ===== END =====

	// ===== BERITA ======
	e.POST("/addberitatips", controllers.SimpanBerita)
	e.GET("/getallberita", controllers.DataAllBerita)
	// ===== END OF BERITA =====

	// ===== INFO PELATIHAN =====
	e.POST("/addberitainfo", controllers.SimpanInfo)
	e.GET("/getallinfo", controllers.DataAllInfo)
	// ===== END OF INFO ======

	// ===== PENGADUAN ====
	e.POST("/addpengaduan", controllers.SimpanPengaduan)
	e.GET("/getallpengaduan", controllers.DataAllPengaduan)
	// ==== END OF PENGADUAN =====

	return e
}
