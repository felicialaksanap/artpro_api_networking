package controllers

import (
	"artpro_api_networking/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func UploadFoto(c echo.Context) error {
	folder := c.FormValue("folder")
	id := c.FormValue("id")

	result, err := models.UploadFoto(c.Request(), id, folder)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func GetPhoto(c echo.Context) error {
	id := c.FormValue("id")
	folder := c.FormValue("folder")
	result := models.GetPhoto(folder, id)
	return c.File(result)
}

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

func UpdateEmailUser(c echo.Context) error {
	iduser := c.FormValue("iduser")
	email := c.FormValue("email")

	ii, _ := strconv.Atoi(iduser)

	result, err := models.UpdateEmailUser(ii, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdatePassUser(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	result, err := models.UpdatePassUser(email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SimpanProfileUser(c echo.Context) error {
	iduser := c.FormValue("iduser")
	namalengkap := c.FormValue("namalengkap")
	jeniskelamin := c.FormValue("jeniskelamin")
	tempatlahir := c.FormValue("tempatlahir")
	tanggallahir := c.FormValue("tanggallahir")
	telephone := c.FormValue("telephone")
	profilepicpath := c.FormValue("profilepicpath")

	ii, _ := strconv.Atoi(iduser)

	result, err := models.SimpanProfileUser(ii, namalengkap, jeniskelamin, tempatlahir, tanggallahir, telephone, profilepicpath)
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

func UpdateProfileUser(c echo.Context) error {
	iduser := c.FormValue("iduser")
	namalengkap := c.FormValue("namalengkap")
	jeniskelamin := c.FormValue("jeniskelamin")
	tempatlahir := c.FormValue("tempatlahir")
	tanggallahir := c.FormValue("tanggallahir")
	telephone := c.FormValue("telephone")

	ii, _ := strconv.Atoi(iduser)

	result, err := models.UpdateProfileUser(ii, namalengkap, jeniskelamin, tempatlahir, tanggallahir, telephone)
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

func UpdateDataVerifikasi(c echo.Context) error {
	statusverifikasi := c.FormValue("statusverifikasi")
	alasan := c.FormValue("alasan")
	iduser := c.FormValue("iduser")

	ii, _ := strconv.Atoi(iduser)

	result, err := models.UpdateDataVerifikasi(ii, statusverifikasi, alasan)
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

func UpdateUserDomisili(c echo.Context) error {
	iduser := c.FormValue("iduser")
	alamat := c.FormValue("alamat")
	kecamatan := c.FormValue("kecamatan")
	kelurahan := c.FormValue("kelurahan")
	provinsi := c.FormValue("provinsi")
	kota := c.FormValue("kota")
	longitude := c.FormValue("longitude")
	latitude := c.FormValue("latitude")

	ii, _ := strconv.Atoi(iduser)

	result, err := models.UpdateUserDomisili(ii, alamat, kecamatan, kelurahan, provinsi, kota, longitude, latitude)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataLongLat(c echo.Context) error {
	iduser := c.FormValue("iduser")

	ii, _ := strconv.Atoi(iduser)

	result, err := models.DataLongLatUser(ii)
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
	aislam := c.FormValue("aislam")
	akatolik := c.FormValue("akatolik")
	akristen := c.FormValue("akristen")
	ahindu := c.FormValue("ahindu")
	abuddha := c.FormValue("abuddha")
	akonghucu := c.FormValue("akonghucu")
	tkmenginap := c.FormValue("tkmenginap")
	tkwarnen := c.FormValue("tkwarnenr")
	hewan := c.FormValue("hewan")
	mabukjalan := c.FormValue("mabukjalan")
	sepedamotor := c.FormValue("sepedamotor")
	mobil := c.FormValue("mobil")
	masak := c.FormValue("masak")
	ssingle := c.FormValue("ssingle")
	smarried := c.FormValue("smarried")

	ii, _ := strconv.Atoi(iduser)
	bbi, _ := strconv.Atoi(beratbadan)
	tti, _ := strconv.Atoi(tinggibadan)
	aii, _ := strconv.Atoi(aislam)
	akti, _ := strconv.Atoi(akatolik)
	akri, _ := strconv.Atoi(akristen)
	ahi, _ := strconv.Atoi(ahindu)
	abi, _ := strconv.Atoi(abuddha)
	akhi, _ := strconv.Atoi(akonghucu)
	tmi, _ := strconv.Atoi(tkmenginap)
	twi, _ := strconv.Atoi(tkwarnen)
	hi, _ := strconv.Atoi(hewan)
	mji, _ := strconv.Atoi(mabukjalan)
	spmi, _ := strconv.Atoi(sepedamotor)
	mbi, _ := strconv.Atoi(mobil)
	mi, _ := strconv.Atoi(masak)
	si, _ := strconv.Atoi(ssingle)
	smi, _ := strconv.Atoi(smarried)

	result, err := models.SimpanDetailProfileART(ii, pendidikanterakhir, bbi, tti, aii, akti, akri, ahi, abi, akhi, tmi, twi, hi, mji, spmi, mbi, mi, si, smi)
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

func UpdateUserDetailProfileART(c echo.Context) error {
	iduser := c.FormValue("iduser")
	pendidikanterakhir := c.FormValue("pendidikanterakhir")
	beratbadan := c.FormValue("beratbadan")
	tinggibadan := c.FormValue("tinggibadan")
	aislam := c.FormValue("aislam")
	akatolik := c.FormValue("akatolik")
	akristen := c.FormValue("akristen")
	ahindu := c.FormValue("ahindu")
	abuddha := c.FormValue("abuddha")
	akonghucu := c.FormValue("akonghucu")
	tkmenginap := c.FormValue("tkmenginap")
	tkwarnen := c.FormValue("tkwarnen")
	hewan := c.FormValue("hewan")
	mabukjalan := c.FormValue("mabukjalan")
	sepedamotor := c.FormValue("sepedamotor")
	mobil := c.FormValue("mobil")
	masak := c.FormValue("masak")
	ssingle := c.FormValue("ssingle")
	smarried := c.FormValue("smarried")

	ii, _ := strconv.Atoi(iduser)
	bi, _ := strconv.Atoi(beratbadan)
	ti, _ := strconv.Atoi(tinggibadan)
	aii, _ := strconv.Atoi(aislam)
	akti, _ := strconv.Atoi(akatolik)
	akri, _ := strconv.Atoi(akristen)
	ahi, _ := strconv.Atoi(ahindu)
	abi, _ := strconv.Atoi(abuddha)
	akhi, _ := strconv.Atoi(akonghucu)
	tmi, _ := strconv.Atoi(tkmenginap)
	twi, _ := strconv.Atoi(tkwarnen)
	hi, _ := strconv.Atoi(hewan)
	mji, _ := strconv.Atoi(mabukjalan)
	spmi, _ := strconv.Atoi(sepedamotor)
	mbi, _ := strconv.Atoi(mobil)
	mi, _ := strconv.Atoi(masak)
	si, _ := strconv.Atoi(ssingle)
	smi, _ := strconv.Atoi(smarried)

	result, err := models.UpdateUserDetailProfileART(ii, pendidikanterakhir, bi, ti, aii, akti, akri, ahi, abi, akhi, tmi, twi, hi, mji, spmi, mbi, mi, si, smi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SimpanDetailKerjaART(c echo.Context) error {
	iduser := c.FormValue("iduser")
	kprt := c.FormValue("kprt")
	kbabysitter := c.FormValue("kbabysitter")
	kseniorcare := c.FormValue("kseniorcare")
	ksupir := c.FormValue("ksupir")
	kofficeboy := c.FormValue("kofficeboy")
	ktukangkebun := c.FormValue("ktukangkebun")
	pengalaman := c.FormValue("pengalaman")
	gajiawal := c.FormValue("gajiawal")
	gajiakhir := c.FormValue("gajiakhir")

	ii, _ := strconv.Atoi(iduser)
	kpi, _ := strconv.Atoi(kprt)
	kbi, _ := strconv.Atoi(kbabysitter)
	ksci, _ := strconv.Atoi(kseniorcare)
	ksi, _ := strconv.Atoi(ksupir)
	kobi, _ := strconv.Atoi(kofficeboy)
	ktki, _ := strconv.Atoi(ktukangkebun)
	gai, _ := strconv.Atoi(gajiawal)
	gaki, _ := strconv.Atoi(gajiakhir)

	result, err := models.SimpanDetailKerjaART(ii, kpi, kbi, ksci, ksi, kobi, ktki, pengalaman, gai, gaki)
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

func DataARTbyKategori(c echo.Context) error {
	kategori := c.FormValue("kategori")

	result, err := models.DataARTbyKategori(kategori)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataARTbyFK(c echo.Context) error {
	kategori := c.FormValue("kategori")
	idmajikan := c.FormValue("idmajikan")
	aislam := c.FormValue("aislam")
	akatolik := c.FormValue("akatolik")
	akristen := c.FormValue("akristen")
	ahindu := c.FormValue("ahindu")
	abuddha := c.FormValue("abuddha")
	akonghucu := c.FormValue("akonghucu")
	tkmenginap := c.FormValue("tkmenginap")
	tkwarnen := c.FormValue("tkwarnen")
	hewan := c.FormValue("hewan")
	mabukjalan := c.FormValue("mabukjalan")
	sepedamotor := c.FormValue("sepedamotor")
	mobil := c.FormValue("mobil")
	masak := c.FormValue("masak")
	ssingle := c.FormValue("ssingle")
	smarried := c.FormValue("smarried")
	kprt := c.FormValue("kprt")
	kbabysitter := c.FormValue("kbabysitter")
	kseniorcare := c.FormValue("kseniorcare")
	ksupir := c.FormValue("ksupir")
	kofficeboy := c.FormValue("kofficeboy")
	ktukangkebun := c.FormValue("ktukangkebun")
	gajiawal := c.FormValue("gajiawal")
	gajiakhir := c.FormValue("gajiakhir")
	jarak := c.FormValue("jarak")
	updatestatusjarak := c.FormValue("updatestatusjarak")

	imi, _ := strconv.Atoi(idmajikan)
	aii, _ := strconv.Atoi(aislam)
	akti, _ := strconv.Atoi(akatolik)
	akri, _ := strconv.Atoi(akristen)
	ahi, _ := strconv.Atoi(ahindu)
	abi, _ := strconv.Atoi(abuddha)
	akhi, _ := strconv.Atoi(akonghucu)
	tkmi, _ := strconv.Atoi(tkmenginap)
	tkwi, _ := strconv.Atoi(tkwarnen)
	hi, _ := strconv.Atoi(hewan)
	mji, _ := strconv.Atoi(mabukjalan)
	spdmi, _ := strconv.Atoi(sepedamotor)
	mbi, _ := strconv.Atoi(mobil)
	mi, _ := strconv.Atoi(masak)
	ssi, _ := strconv.Atoi(ssingle)
	smi, _ := strconv.Atoi(smarried)
	kpi, _ := strconv.Atoi(kprt)
	kbsi, _ := strconv.Atoi(kbabysitter)
	ksci, _ := strconv.Atoi(kseniorcare)
	ksi, _ := strconv.Atoi(ksupir)
	kobi, _ := strconv.Atoi(kofficeboy)
	ktki, _ := strconv.Atoi(ktukangkebun)
	gawi, _ := strconv.Atoi(gajiawal)
	gaki, _ := strconv.Atoi(gajiakhir)
	ji, _ := strconv.Atoi(jarak)

	result, err := models.DataARTbyFK(kategori, imi, aii, akti, akri, ahi, abi, akhi, tkmi, tkwi, hi, mji, spdmi, mbi, mi, ssi, smi, kpi, kbsi, ksci, ksi, kobi, ktki, gawi, gaki, ji, updatestatusjarak)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CreateAndInsertTableTemp(c echo.Context) error {
	kategori := c.FormValue("kategori")
	idmajikan := c.FormValue("idmajikan")

	ii, _ := strconv.Atoi(idmajikan)
	result, err := models.CreateAndInsertTableTemp(kategori, ii)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateJarak(c echo.Context) error {
	idmajikan := c.FormValue("idmajikan")
	idart := c.FormValue("idart")
	jarak := c.FormValue("jarak")

	imi, _ := strconv.Atoi(idmajikan)
	iai, _ := strconv.Atoi(idart)
	jf, _ := strconv.ParseFloat(jarak, 64)

	result, err := models.UpdateJarak(imi, iai, jf)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataUserDetailKerjaART(c echo.Context) error {
	iduser := c.FormValue("iduser")

	ii, _ := strconv.Atoi(iduser)
	result, err := models.DataUserDetailKerjaART(ii)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateUserDetailKerja(c echo.Context) error {
	iduser := c.FormValue("iduser")
	kprt := c.FormValue("kprt")
	kbabysitter := c.FormValue("kbabysitter")
	kseniorcare := c.FormValue("kseniorcare")
	ksupir := c.FormValue("ksupir")
	kofficeboy := c.FormValue("kofficeboy")
	ktukangkebun := c.FormValue("ktukangkebun")
	pengalaman := c.FormValue("pengalaman")
	gajiawal := c.FormValue("gajiawal")
	gajiakhir := c.FormValue("gajiakhir")

	ii, _ := strconv.Atoi(iduser)
	kpi, _ := strconv.Atoi(kprt)
	kbi, _ := strconv.Atoi(kbabysitter)
	ksci, _ := strconv.Atoi(kseniorcare)
	ksi, _ := strconv.Atoi(ksupir)
	kobi, _ := strconv.Atoi(kofficeboy)
	ktki, _ := strconv.Atoi(ktukangkebun)
	gai, _ := strconv.Atoi(gajiawal)
	gaki, _ := strconv.Atoi(gajiakhir)

	result, err := models.UpdateUserDetailKerja(ii, kpi, kbi, ksci, ksi, kobi, ktki, pengalaman, gai, gaki)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SimpanKontakUser(c echo.Context) error {
	idmajikan := c.FormValue("idmajikan")
	idart := c.FormValue("idart")
	waktukontak := c.FormValue("waktukontak")

	imi, _ := strconv.Atoi(idmajikan)
	iai, _ := strconv.Atoi(idart)

	result, err := models.SimpanKontakuser(imi, iai, waktukontak)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataKontakART(c echo.Context) error {
	idmajikan := c.FormValue("idmajikan")

	ii, _ := strconv.Atoi(idmajikan)

	result, err := models.DataKontakART(ii)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataKontakMajikan(c echo.Context) error {
	idart := c.FormValue("idart")

	ii, _ := strconv.Atoi(idart)

	result, err := models.DataKontakMajikan(ii)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func KontakART(c echo.Context) error {
	idart := c.FormValue("idart")

	ii, _ := strconv.Atoi(idart)

	result, err := models.KontakArt(ii)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SimpanPenilaian(c echo.Context) error {
	idart := c.FormValue("idart")
	idmajikan := c.FormValue("idmajikan")
	estetika := c.FormValue("estetika")
	etika := c.FormValue("etika")
	kebersihan := c.FormValue("kebersihan")
	kecepatan := c.FormValue("kecepatan")
	kerapian := c.FormValue("kerapian")
	rating := c.FormValue("rating")
	review := c.FormValue("review")
	tglpost := c.FormValue("tglpost")

	iai, _ := strconv.Atoi(idart)
	imi, _ := strconv.Atoi(idmajikan)
	ei, _ := strconv.Atoi(etika)
	esi, _ := strconv.Atoi(estetika)
	ki, _ := strconv.Atoi(kebersihan)
	kri, _ := strconv.Atoi(kerapian)
	kci, _ := strconv.Atoi(kecepatan)
	ri, _ := strconv.ParseFloat(rating, 64)

	result, err := models.SimpanPenilaian(iai, imi, ei, esi, ki, kri, kci, ri, review, tglpost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func RataPenilaianART(c echo.Context) error {
	idart := c.FormValue("idart")

	ii, _ := strconv.Atoi(idart)

	result, err := models.RataPenilaianART(ii)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataReviewMajikan(c echo.Context) error {
	idart := c.FormValue("idart")

	ii, _ := strconv.Atoi(idart)

	result, err := models.DataReviewMajikan(ii)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//func SimpanSertifikatPelatihan(c echo.Context) error {
//	iduser := c.FormValue("iduser")
//	sertifpath := c.FormValue("sertifpath")
//
//	ii, _ := strconv.Atoi(iduser)
//
//	result, err := models.SimpanSertifikatPelatihan(ii, sertifpath)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
//	}
//
//	return c.JSON(http.StatusOK, result)
//}
//
//func DataSertifPelatihanUser(c echo.Context) error {
//	iduser := c.FormValue("iduser")
//
//	ii, _ := strconv.Atoi(iduser)
//
//	result, err := models.DataSertifPelatihanUser(ii)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
//	}
//
//	return c.JSON(http.StatusOK, result)
//}
