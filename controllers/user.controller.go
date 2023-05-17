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

func SimpanDetailProfileART(c echo.Context) error {
	iduser := c.FormValue("iduser")
	pendidikanterakhir := c.FormValue("pendidikanterakhir")
	beratbadan := c.FormValue("beratbadan")
	tinggibadan := c.FormValue("tinggibadan")
	agama := c.FormValue("agama")
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
	tmi, _ := strconv.Atoi(tkmenginap)
	twi, _ := strconv.Atoi(tkwarnen)
	hi, _ := strconv.Atoi(hewan)
	mji, _ := strconv.Atoi(mabukjalan)
	spmi, _ := strconv.Atoi(sepedamotor)
	mbi, _ := strconv.Atoi(mobil)
	mi, _ := strconv.Atoi(masak)
	si, _ := strconv.Atoi(ssingle)
	smi, _ := strconv.Atoi(smarried)

	result, err := models.SimpanDetailProfileART(ii, pendidikanterakhir, bbi, tti, agama, tmi, twi, hi, mji, spmi, mbi, mi, si, smi)
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
	agama := c.FormValue("agama")
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
	tmi, _ := strconv.Atoi(tkmenginap)
	twi, _ := strconv.Atoi(tkwarnen)
	hi, _ := strconv.Atoi(hewan)
	mji, _ := strconv.Atoi(mabukjalan)
	spmi, _ := strconv.Atoi(sepedamotor)
	mbi, _ := strconv.Atoi(mobil)
	mi, _ := strconv.Atoi(masak)
	si, _ := strconv.Atoi(ssingle)
	smi, _ := strconv.Atoi(smarried)

	result, err := models.UpdateUserDetailProfileART(ii, pendidikanterakhir, bi, ti, agama, tmi, twi, hi, mji, spmi, mbi, mi, si, smi)
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

	result, err := models.SimpanDetailKerjaART(ii, kpi, kbi, ksci, ksi, kobi, ktki, pengalaman, gajiawal, gajiakhir)
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

	result, err := models.UpdateUserDetailKerja(ii, kpi, kbi, ksci, ksi, kobi, ktki, pengalaman, gajiawal, gajiakhir)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

//func GetTotalKontakART(c echo.Context) error {
//	idart := c.FormValue("idart")
//
//	ii, _ := strconv.Atoi(idart)
//
//	result, err := models.GetTotalKontakART(ii)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
//	}
//
//	return c.JSON(http.StatusOK, result)
//}

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

//func DataListKontakByMajikan(c echo.Context) error {
//	idmajikan := c.FormValue("idmajikan")
//
//	ii, _ := strconv.Atoi(idmajikan)
//
//	result, err := models.DataListKontakByMajikan(ii)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
//	}
//
//	return c.JSON(http.StatusOK, result)
//}
//
//func DataListKontakByART(c echo.Context) error {
//	idart := c.FormValue("idart")
//
//	ii, _ := strconv.Atoi(idart)
//
//	result, err := models.DataListKontakByART(ii)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
//	}
//
//	return c.JSON(http.StatusOK, result)
//}

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

	iai, _ := strconv.Atoi(idart)
	imi, _ := strconv.Atoi(idmajikan)
	ei, _ := strconv.Atoi(etika)
	esi, _ := strconv.Atoi(estetika)
	ki, _ := strconv.Atoi(kebersihan)
	kri, _ := strconv.Atoi(kerapian)
	kci, _ := strconv.Atoi(kecepatan)
	ri, _ := strconv.ParseFloat(rating, 64)

	result, err := models.SimpanPenilaian(iai, imi, ei, esi, ki, kri, kci, ri, review)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataPenilaianART(c echo.Context) error {
	idart := c.FormValue("idart")

	ii, _ := strconv.Atoi(idart)

	result, err := models.DataPenilaianART(ii)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetRatingART(c echo.Context) error {
	idart := c.FormValue("idart")

	ii, _ := strconv.Atoi(idart)

	result, err := models.GetRatingART(ii)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SimpanSertifikatPelatihan(c echo.Context) error {
	iduser := c.FormValue("iduser")
	sertifpath := c.FormValue("sertifpath")

	ii, _ := strconv.Atoi(iduser)

	result, err := models.SimpanSertifikatPelatihan(ii, sertifpath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DataSertifPelatihanUser(c echo.Context) error {
	iduser := c.FormValue("iduser")

	ii, _ := strconv.Atoi(iduser)

	result, err := models.DataSertifPelatihanUser(ii)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
