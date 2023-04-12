package controllers

import (
	"artpro_api_networking/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func SimpanAkunUser(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	statususer := c.FormValue("statususer")

	result, err := models.SimpanAkunUser(email, password, statususer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataAkunUser(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	result, err := models.DataAkunUser(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SimpanProfileUser(c echo.Context) error {
	iduser := c.FormValue("iduser")
	namalengkap := c.FormValue("namalengkap")
	jeniskelamin := c.FormValue("jeniskelamin")
	tanggallahir := c.FormValue("tanggallahir")
	telephone := c.FormValue("telephone")
	profilepicpath := c.FormValue("profilepicpath")

	ii, _ := strconv.Atoi(iduser)

	result, err := models.SimpanProfileUser(ii, namalengkap, jeniskelamin, tanggallahir, telephone, profilepicpath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataProfileUser(c echo.Context) error {
	iduser := c.FormValue("iduser")

	ii, _ := strconv.Atoi(iduser)

	result, err := models.DataProfileUser(ii)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SimpanDataVerifikasi(c echo.Context) error {
	iduser := c.FormValue("iduser")
	nik := c.FormValue("nik")
	tempatlahir := c.FormValue("tempatlahir")
	tanggallahir := c.FormValue("tanggallahir")
	alamat := c.FormValue("alamat")
	kecamatan := c.FormValue("kecamatan")
	kelurahan := c.FormValue("kelurahan")
	rt := c.FormValue("rt")
	rw := c.FormValue("rw")
	fotoktp := c.FormValue("fotoktp")
	selfiektp := c.FormValue("selfiektp")
	statusverifikasi := c.FormValue("statusverifikasi")

	ii, _ := strconv.Atoi(iduser)
	rti, _ := strconv.Atoi(rt)
	rwi, _ := strconv.Atoi(rw)

	result, err := models.SimpanDataVerifikasi(ii, nik, tempatlahir, tanggallahir, alamat, kecamatan, kelurahan, rti, rwi, fotoktp, selfiektp, statusverifikasi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataAllVerifikasi(c echo.Context) error {
	result, err := models.DataAllVerifikasi()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataVerifikasiUser(c echo.Context) error {
	iduser := c.FormValue("iduser")

	ii, _ := strconv.Atoi(iduser)

	result, err := models.DataVerifikasiUser(ii)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateDataVerifikasi(c echo.Context) error {
	statusverifikasi := c.FormValue("statusverifikasi")
	iduser := c.FormValue("iduser")

	ii, _ := strconv.Atoi(iduser)

	result, err := models.UpdateDataVerifikasi(ii, statusverifikasi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SimpanDomisiliUser(c echo.Context) error {
	iduser := c.FormValue("iduser")
	alamat := c.FormValue("alamat")
	kecamatan := c.FormValue("kecamatan")
	kelurahan := c.FormValue("kelurahan")
	provinsi := c.FormValue("provinsi")
	kota := c.FormValue("kota")
	longitude := c.FormValue("longitude")
	latitude := c.FormValue("latitude")

	ii, _ := strconv.Atoi(iduser)

	result, err := models.SimpanDomisiliUser(ii, alamat, kecamatan, kelurahan, provinsi, kota, longitude, latitude)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataUserDomisili(c echo.Context) error {
	iduser := c.FormValue("iduser")

	ii, _ := strconv.Atoi(iduser)

	result, err := models.DataUserDomisili(ii)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SimpanDetailProfileART(c echo.Context) error {
	iduser := c.FormValue("iduser")
	pendidikanterakhir := c.FormValue("pendidikanterakhir")
	beratbadan := c.FormValue("beratbadan")
	tinggibadan := c.FormValue("tinggibadan")
	agama := c.FormValue("agama")
	tipekerja := c.FormValue("tipekerja")
	hewan := c.FormValue("hewan")
	mabukjalan := c.FormValue("mabukjalan")
	sepedamotor := c.FormValue("sepedamotor")
	mobil := c.FormValue("mobil")
	masak := c.FormValue("masak")

	ii, _ := strconv.Atoi(iduser)
	bbi, _ := strconv.Atoi(beratbadan)
	tti, _ := strconv.Atoi(tinggibadan)

	result, err := models.SimpanDetailProfileART(ii, pendidikanterakhir, bbi, tti, agama, tipekerja, hewan, mabukjalan, sepedamotor, mobil, masak)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataUserDetailProfileART(c echo.Context) error {
	iduser := c.FormValue("iduser")

	ii, _ := strconv.Atoi(iduser)

	result, err := models.DataUserDetailProfileART(ii)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SimpanDetailKerjaART(c echo.Context) error {
	iduser := c.FormValue("iduser")
	kategori := c.FormValue("kategori")
	pengalaman := c.FormValue("pengalaman")
	gajiawal := c.FormValue("gajiawal")
	gajiakhir := c.FormValue("gajiakhir")

	ii, _ := strconv.Atoi(iduser)

	result, err := models.SimpanDetailKerjaART(ii, kategori, pengalaman, gajiawal, gajiakhir)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataAllDetailKerjaART(c echo.Context) error {
	result, err := models.DataAllDetailKerjaART()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataListKerjaPerKategori(c echo.Context) error {
	kategori := c.FormValue("kategori")

	result, err := models.DataListKerjaPerKategori(kategori)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SimpanLowonganKerja(c echo.Context) error {
	iduser := c.FormValue("iduser")
	kategori := c.FormValue("kategori")
	informasi := c.FormValue("informasi")
	uraiantugas := c.FormValue("uraiantugas")
	keahlian := c.FormValue("keahlian")
	tipekerja := c.FormValue("tipekerja")
	gajiawal := c.FormValue("gajiawal")
	gajiakhir := c.FormValue("gajiakhir")

	ii, _ := strconv.Atoi(iduser)

	result, err := models.SimpanLowonganKerja(ii, kategori, informasi, uraiantugas, keahlian, tipekerja, gajiawal, gajiakhir)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
