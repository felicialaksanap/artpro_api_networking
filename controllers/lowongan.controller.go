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
	ssingle := c.FormValue("ssingle")
	smarried := c.FormValue("smarried")
	tglpost := c.FormValue("tglpost")
	statusloker := c.FormValue("statusloker")

	ii, _ := strconv.Atoi(iduser)
	gai, _ := strconv.Atoi(gajiawal)
	gaki, _ := strconv.Atoi(gajiakhir)
	kpi, _ := strconv.Atoi(kprt)
	kbi, _ := strconv.Atoi(kbabysitter)
	ksci, _ := strconv.Atoi(kseniorcare)
	ksi, _ := strconv.Atoi(ksupir)
	kobi, _ := strconv.Atoi(kofficeboy)
	ktki, _ := strconv.Atoi(ktukangkebun)
	hi, _ := strconv.Atoi(hewan)
	mi, _ := strconv.Atoi(masak)
	mji, _ := strconv.Atoi(mabukjalan)
	spdi, _ := strconv.Atoi(sepedamotor)
	mbi, _ := strconv.Atoi(mobil)
	tmi, _ := strconv.Atoi(tkmenginap)
	twi, _ := strconv.Atoi(tkwarnen)
	ssi, _ := strconv.Atoi(ssingle)
	smi, _ := strconv.Atoi(smarried)
	sti, _ := strconv.Atoi(statusloker)

	result, err := models.SimpanLowonganKerja(ii, judulloker, gai, gaki, informasi, tugas, kpi, kbi, ksci, ksi, kobi, ktki, hi, mi, mji, spdi, mbi, tmi, twi, ssi, smi, tglpost, sti)
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

func UpdateStatusLoker(c echo.Context) error {
	idloker := c.FormValue("idloker")
	statusloker := c.FormValue("statusloker")
	tglpost := c.FormValue("tglpost")

	il, _ := strconv.Atoi(idloker)
	sti, _ := strconv.Atoi(statusloker)

	result, err := models.UpdateStatusLoker(il, sti, tglpost)
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
	ssingle := c.FormValue("ssingle")
	smarried := c.FormValue("smarried")
	tglpost := c.FormValue("tglpost")

	ili, _ := strconv.Atoi(idloker)
	gai, _ := strconv.Atoi(gajiawal)
	gaki, _ := strconv.Atoi(gajiakhir)
	kpi, _ := strconv.Atoi(kprt)
	kbi, _ := strconv.Atoi(kbabysitter)
	ksci, _ := strconv.Atoi(kseniorcare)
	ksi, _ := strconv.Atoi(ksupir)
	kobi, _ := strconv.Atoi(kofficeboy)
	ktki, _ := strconv.Atoi(ktukangkebun)
	hi, _ := strconv.Atoi(hewan)
	mi, _ := strconv.Atoi(masak)
	mji, _ := strconv.Atoi(mabukjalan)
	spdi, _ := strconv.Atoi(sepedamotor)
	mbi, _ := strconv.Atoi(mobil)
	tmi, _ := strconv.Atoi(tkmenginap)
	twi, _ := strconv.Atoi(tkwarnen)
	ssi, _ := strconv.Atoi(ssingle)
	smi, _ := strconv.Atoi(smarried)

	result, err := models.UpdateLowonganKerja(ili, judulloker, gai, gaki, informasi, tugas, kpi, kbi, ksci, ksi, kobi, ktki, hi, mi, mji, spdi, mbi, tmi, twi, ssi, smi, tglpost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
