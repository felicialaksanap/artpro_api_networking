package controllers

import (
	"artpro_api_networking/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func SimpanLowonganKerja(c echo.Context) error {
	iduser := c.FormValue("iduser")
	judulloker := c.FormValue("judulloker")
	gajiawal := c.FormValue("gajiawal")
	gajiakhir := c.FormValue("gajiakhir")
	informasi := c.FormValue("informasi")
	tugas := c.FormValue("tugas")
	kprt := c.FormValue("kprt")
	kbabysitter := c.FormValue("kbabysitter")
	kseniorcare := c.FormValue("kseniorcare")
	ksupir := c.FormValue("ksupir")
	kofficeboy := c.FormValue("kofficeboy")
	ktukangkebun := c.FormValue("ktukangkebun")
	hewan := c.FormValue("hewan")
	masak := c.FormValue("masak")
	mabukjalan := c.FormValue("mabukjalan")
	sepedamotor := c.FormValue("sepedamotor")
	mobil := c.FormValue("mobil")
	tkmenginap := c.FormValue("tkmenginap")
	tkwarnen := c.FormValue("tkwarnen")
	tglpost := c.FormValue("tglpost")
	statusloker := c.FormValue("statusloker")

	ii, _ := strconv.Atoi(iduser)

	result, err := models.SimpanLowonganKerja(ii, judulloker, gajiawal, gajiakhir, informasi, tugas, kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, hewan, masak, mabukjalan, sepedamotor, mobil, tkmenginap, tkwarnen, tglpost, statusloker)
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

func DataLowonganKerjaperUser(c echo.Context) error {
	iduser := c.FormValue("iduser")

	iu, _ := strconv.Atoi(iduser)

	result, err := models.DataLowonganKerjaperUser(iu)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataLowonganKerjaperIdLoker(c echo.Context) error {
	idloker := c.FormValue("idloker")

	il, _ := strconv.Atoi(idloker)

	result, err := models.DataLowonganKerjaperUser(il)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SimpanLowonganKerjaDetail(c echo.Context) error {
	idloker := c.FormValue("idloker")
	iduser := c.FormValue("iduser")
	judulloker := c.FormValue("judulloker")
	gajiawal := c.FormValue("gajiawal")
	gajiakhir := c.FormValue("gajiakhir")
	informasi := c.FormValue("informasi")
	tugas := c.FormValue("tugas")
	kprt := c.FormValue("kprt")
	kbabysitter := c.FormValue("kbabysitter")
	kseniorcare := c.FormValue("kseniorcare")
	ksupir := c.FormValue("ksupir")
	kofficeboy := c.FormValue("kofficeboy")
	ktukangkebun := c.FormValue("ktukangkebun")
	hewan := c.FormValue("hewan")
	masak := c.FormValue("masak")
	mabukjalan := c.FormValue("mabukjalan")
	sepedamotor := c.FormValue("sepedamotor")
	mobil := c.FormValue("mobil")
	tkmenginap := c.FormValue("tkmenginap")
	tkwarnen := c.FormValue("tkwarnen")
	tglmodif := c.FormValue("tglpost")
	statusloker := c.FormValue("statusloker")
	alasan := c.FormValue("alasan")

	ili, _ := strconv.Atoi(idloker)
	ii, _ := strconv.Atoi(iduser)

	result, err := models.SimpanLowonganKerjaDetail(ili, ii, judulloker, gajiawal, gajiakhir, informasi, tugas, kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, hewan, masak, mabukjalan, sepedamotor, mobil, tkmenginap, tkwarnen, tglmodif, statusloker, alasan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateStatusLoker(c echo.Context) error {
	idloker := c.FormValue("idloker")
	statusloker := c.FormValue("statusloker")
	tglpost := c.FormValue("tglpost")

	il, _ := strconv.Atoi(idloker)

	result, err := models.UpdateStatusLoker(il, statusloker, tglpost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func UpdateLowonganKerja(c echo.Context) error {
	idloker := c.FormValue("idloker")
	judulloker := c.FormValue("judulloker")
	gajiawal := c.FormValue("gajiawal")
	gajiakhir := c.FormValue("gajiakhir")
	informasi := c.FormValue("informasi")
	tugas := c.FormValue("tugas")
	kprt := c.FormValue("kprt")
	kbabysitter := c.FormValue("kbabysitter")
	kseniorcare := c.FormValue("kseniorcare")
	ksupir := c.FormValue("ksupir")
	kofficeboy := c.FormValue("kofficeboy")
	ktukangkebun := c.FormValue("ktukangkebun")
	hewan := c.FormValue("hewan")
	masak := c.FormValue("masak")
	mabukjalan := c.FormValue("mabukjalan")
	sepedamotor := c.FormValue("sepedamotor")
	mobil := c.FormValue("mobil")
	tkmenginap := c.FormValue("tkmenginap")
	tkwarnen := c.FormValue("tkwarnen")
	tglpost := c.FormValue("tglpost")

	ili, _ := strconv.Atoi(idloker)

	result, err := models.UpdateLowonganKerja(ili, judulloker, gajiawal, gajiakhir, informasi, tugas, kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, hewan, masak, mabukjalan, sepedamotor, mobil, tkmenginap, tkwarnen, tglpost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
