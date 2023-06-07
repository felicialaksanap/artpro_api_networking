package controllers

import (
	"artpro_api_networking/models"
	"github.com/labstack/echo"
	"net/http"
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
