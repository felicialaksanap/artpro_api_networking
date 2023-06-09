package controllers

import (
	"artpro_api_networking/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func SimpanBerita(c echo.Context) error {
	judul := c.FormValue("judul")
	isi := c.FormValue("isi")
	url := c.FormValue("url")
	tglpost := c.FormValue("tglpost")

	result, err := models.SimpanBerita(judul, isi, url, tglpost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataAllBerita(c echo.Context) error {
	result, err := models.DataAllBerita()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SimpanInfo(c echo.Context) error {
	judul := c.FormValue("judul")
	isi := c.FormValue("isi")
	url := c.FormValue("url")
	tglpost := c.FormValue("tglpost")

	result, err := models.SimpanInfo(judul, isi, url, tglpost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataAllInfo(c echo.Context) error {
	result, err := models.DataALLInfo()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SimpanPengaduan(c echo.Context) error {
	idmajikan := c.FormValue("idmajikan")
	idart := c.FormValue("idart")
	isipengaduan := c.FormValue("isipengaduan")
	penyelesaian := c.FormValue("penyelesaian")
	tglpengaduan := c.FormValue("tglpengaduan")

	imi, _ := strconv.Atoi(idmajikan)
	iai, _ := strconv.Atoi(idart)

	result, err := models.SimpanPengaduan(imi, iai, isipengaduan, penyelesaian, tglpengaduan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdatePenyelesaian(c echo.Context) error {
	idloker := c.FormValue("idloker")
	penyelesaian := c.FormValue("penyelesaian")

	ili, _ := strconv.Atoi(idloker)

	result, err := models.UpdatePenyelesaian(ili, penyelesaian)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataAllPengaduan(c echo.Context) error {
	result, err := models.DataALLPengaduan()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
