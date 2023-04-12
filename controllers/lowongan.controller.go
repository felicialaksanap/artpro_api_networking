package controllers

import (
	"artpro_api_networking/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

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

func DataAllLowonganKerja(c echo.Context) error {
	result, err := models.DataAllLowonganKerja()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataDetailLowonganKerja(c echo.Context) error {
	idloker := c.FormValue("idloker")

	ii, _ := strconv.Atoi(idloker)

	result, err := models.DataDetailLowonganKerja(ii)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SimpanLowonganKerjaSelesai(c echo.Context) error {
	idloker := c.FormValue("idloker")
	iduser := c.FormValue("iduser")
	kategori := c.FormValue("kategori")
	informasi := c.FormValue("informasi")
	uraiantugas := c.FormValue("uraiantugas")
	keahlian := c.FormValue("keahlian")
	tipekerja := c.FormValue("tipekerja")
	gajiawal := c.FormValue("gajiawal")
	gajiakhir := c.FormValue("gajiakhir")
	alasan := c.FormValue("alasan")

	ili, _ := strconv.Atoi(idloker)
	ii, _ := strconv.Atoi(iduser)

	result, err := models.SimpanLowonganKerjaSelesai(ili, ii, kategori, informasi, uraiantugas, keahlian, tipekerja, gajiawal, gajiakhir, alasan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteDetailLowonganKerja(c echo.Context) error {
	idloker := c.FormValue("idloker")

	ii, _ := strconv.Atoi(idloker)

	result, err := models.DeleteDetailLowonganKerja(ii)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
