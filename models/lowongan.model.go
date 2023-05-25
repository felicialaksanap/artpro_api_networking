package models

import (
	"artpro_api_networking/db"
	"log"
	"net/http"
)

type LowonganKerja struct {
	IdLoker      int    `json:"idloker"`
	IdUser       int    `json:"iduser"`
	JudulLoker   string `json:"judulloker"`
	GajiAwal     int    `json:"gajiawal"`
	GajiAkhir    int    `json:"gajiakhir"`
	Informasi    string `json:"informasi"`
	Tugas        string `json:"tugas"`
	KPrt         int    `json:"kprt"`
	KBabysitter  int    `json:"kbabysitter"`
	KSeniorcare  int    `json:"kseniorcare"`
	KSupir       int    `json:"ksupir"`
	KOfficeboy   int    `json:"kofficeboy"`
	KTukangkebun int    `json:"ktukangkebun"`
	Hewan        int    `json:"hewan"`
	Masak        int    `json:"masak"`
	MabukJalan   int    `json:"mabukjalan"`
	SepedaMotor  int    `json:"sepedamotor"`
	Mobil        int    `json:"mobil"`
	TKMenginap   int    `json:"tkmenginap"`
	TKWarnen     int    `json:"tkwarnen"`
	SSingle      int    `json:"ssingle"`
	SMarried     int    `json:"smarried"`
	TglPost      string `json:"tglpost"`
	StatusLoker  int    `json:"statusloker"`
	NamaLengkap  string `json:"namalengkap"`
	JenisKelamin string `json:"jeniskelamin"`
	Kecamatan    string `json:"kecamatan"`
	Kota         string `json:"kota"`
}

func SimpanLowonganKerja(iduser int, judulloker string, gajiawal int, gajiakhir int,
	informasi string, tugas string, kprt int, kbabysitter int, kseniorcare int,
	ksupir int, kofficeboy int, ktukangkebun int, hewan int, masak int, mabukjalan int,
	sepedamotor int, mobil int, tkmenginap int, tkwarnen int, ssingle int, smarried int,
	tglpost string, statusloker int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO lowongankerja (iduser, judulloker, gajiawal, gajiakhir, informasi, tugas, kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, hewan, masak, mabukjalan, sepedamotor, mobil, tkmenginap, tkwarnen, ssingle, smarried, tglpost, statusloker) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? ,? ,?, ?, ? )"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(iduser, judulloker, gajiawal, gajiakhir, informasi, tugas, kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, hewan, masak, mabukjalan, sepedamotor, mobil, tkmenginap, tkwarnen, ssingle, smarried, tglpost, statusloker)
	if err != nil {
		return res, err
	}

	defer stmt.Close()

	getIdLast, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = map[string]int64{
		"getIdLast": getIdLast,
	}

	return res, nil
}

func DataAllLowonganKerja() (Response, error) { // menampilkan semua data loker untuk ART
	var obj LowonganKerja
	var arrobj []LowonganKerja
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT lk.idloker, lk.iduser, lk.judulloker, lk.gajiawal, lk.gajiakhir, lk.informasi, lk.tugas, lk.kprt, lk.kbabysitter, lk.kseniorcare, lk.ksupir, lk.kofficeboy, lk.ktukangkebun, lk.hewan, lk.masak, lk.mabukjalan, lk.sepedamotor, lk.mobil, lk.tkmenginap, lk.tkwarnen, lk.ssingle, lk.smarried, lk.tglpost, lk.statusloker, up.namalengkap, up.jeniskelamin, ud.kecamatan, ud.kota FROM lowongankerja lk JOIN userprofile up on lk.iduser = uP.iduser JOIN userdomisili ud on lk.iduser = ud.iduser ORDER BY lk.tglpost DESC"

	rows, err := con.Query(sqlStatemet)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdLoker, &obj.IdUser, &obj.JudulLoker, &obj.GajiAwal, &obj.GajiAkhir,
			&obj.Informasi, &obj.Tugas, &obj.KPrt, &obj.KBabysitter, &obj.KSeniorcare,
			&obj.KSupir, &obj.KOfficeboy, &obj.KTukangkebun, &obj.Hewan, &obj.Masak,
			&obj.MabukJalan, &obj.SepedaMotor, &obj.Mobil, &obj.TKMenginap, &obj.TKWarnen, &obj.SSingle, &obj.SMarried,
			&obj.TglPost, &obj.StatusLoker, &obj.NamaLengkap, &obj.JenisKelamin,
			&obj.Kecamatan, &obj.Kota)

		if err != nil {
			log.Printf(err.Error())
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	log.Printf("berhasil")
	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = arrobj

	return res, nil
}

func DataLowonganKerjaperUser(iduser int) (Response, error) { // menampilkan semua loker yang dibuat oleh majikan
	var obj LowonganKerja
	var arrobj []LowonganKerja
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT lk.idloker, lk.iduser, lk.judulloker, lk.gajiawal, lk.gajiakhir, lk.informasi, lk.tugas, lk.kprt, lk.kbabysitter, lk.kseniorcare, lk.ksupir, lk.kofficeboy, lk.ktukangkebun, lk.hewan, lk.masak, lk.mabukjalan, lk.sepedamotor, lk.mobil, lk.tkmenginap, lk.tkwarnen, lk.ssingle, lk.smarried, lk.tglpost, lk.statusloker, up.namalengkap, up.jeniskelamin, ud.kecamatan, ud.kota FROM lowongankerja lk JOIN userprofile up on lk.iduser = up.iduser JOIN userdomisili ud on lk.iduser = ud.iduser WHERE lk.iduser = ? ORDER BY lk.tglpost DESC"

	rows, err := con.Query(sqlStatemet, iduser)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdLoker, &obj.IdUser, &obj.JudulLoker, &obj.GajiAwal, &obj.GajiAkhir,
			&obj.Informasi, &obj.Tugas, &obj.KPrt, &obj.KBabysitter, &obj.KSeniorcare,
			&obj.KSupir, &obj.KOfficeboy, &obj.KTukangkebun, &obj.Hewan, &obj.Masak,
			&obj.MabukJalan, &obj.SepedaMotor, &obj.Mobil, &obj.TKMenginap, &obj.TKWarnen, &obj.SSingle, &obj.SMarried,
			&obj.TglPost, &obj.StatusLoker, &obj.NamaLengkap, &obj.JenisKelamin,
			&obj.Kecamatan, &obj.Kota)

		if err != nil {
			log.Printf(err.Error())
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	log.Printf("berhasil")
	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = arrobj

	return res, nil
}

func DataLowonganKerjaperIdLoker(iduser int) (Response, error) { // menampilkan semua loker yang dibuat oleh majikan
	var obj LowonganKerja
	var arrobj []LowonganKerja
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT lk.idloker, lk.iduser, lk.judulloker, lk.gajiawal, lk.gajiakhir, lk.informasi, lk.tugas, lk.kprt, lk.kbabysitter, lk.kseniorcare, lk.ksupir, lk.kofficeboy, lk.ktukangkebun, lk.hewan, lk.masak, lk.mabukjalan, lk.sepedamotor, lk.mobil, lk.tkmenginap, lk.tkwarnen, lk.tglpost, lk.statusloker, up.namalengkap, up.jeniskelamin, ud.kecamatan, ud.kota FROM lowongankerja lk JOIN userprofile up on lk.iduser = uP.iduser JOIN userdomisili ud on lk.iduser = ud.iduser WHERE lk.idloker = ?"

	rows, err := con.Query(sqlStatemet, iduser)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdLoker, &obj.IdUser, &obj.JudulLoker, &obj.GajiAwal, &obj.GajiAkhir,
			&obj.Informasi, &obj.Tugas, &obj.KPrt, &obj.KBabysitter, &obj.KSeniorcare,
			&obj.KSupir, &obj.KOfficeboy, &obj.KTukangkebun, &obj.Hewan, &obj.Masak,
			&obj.MabukJalan, &obj.SepedaMotor, &obj.Mobil, &obj.TKMenginap, &obj.TKWarnen,
			&obj.TglPost, &obj.StatusLoker, &obj.NamaLengkap, &obj.JenisKelamin,
			&obj.Kecamatan, &obj.Kota)

		if err != nil {
			log.Printf(err.Error())
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	log.Printf("berhasil")
	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = arrobj

	return res, nil
}

func UpdateStatusLoker(idloker int, statusloker int, tglpost string) (Response, error) { // untuk update status loker aktif atau tidak aktif
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE lowongankerja SET statusloker=?, tglpost=? WHERE idloker=?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(statusloker, tglpost, idloker)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = map[string]int64{
		"rows": rowsAffected,
	}

	return res, nil
}

func UpdateLowonganKerja(idloker int, judulloker string, gajiawal int, gajiakhir int,
	informasi string, tugas string, kprt int, kbabysitter int, kseniorcare int,
	ksupir int, kofficeboy int, ktukangkebun int, hewan int, masak int,
	mabukjalan int, sepedamotor int, mobil int, tkmenginap int, tkwarnen int,
	ssingle int, smarried int, tglpost string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE lowongankerja SET judulloker=?, gajiawal=?, gajiakhir=?, informasi=?, tugas=?, kprt=?, kbabysitter=?, kseniorcare=?, ksupir=?, kofficeboy=?, ktukangkebun=?, hewan=?, masak=?, mabukjalan=?, sepedamotor=?, mobil=?, tkmenginap=?, tkwarnen=?, ssingle=?, smarried=?, tglpost=? WHERE idloker=?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(judulloker, gajiawal, gajiakhir, informasi, tugas, kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, hewan, masak, mabukjalan, sepedamotor, mobil, tkmenginap, tkwarnen, ssingle, smarried, tglpost, idloker)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = map[string]int64{
		"rows": rowsAffected,
	}

	return res, nil
}
