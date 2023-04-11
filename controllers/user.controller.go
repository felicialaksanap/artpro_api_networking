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
