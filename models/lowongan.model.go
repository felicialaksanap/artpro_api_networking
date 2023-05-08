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
	GajiAwal     string `json:"gajiawal"`
	GajiAkhir    string `json:"gajiakhir"`
	Informasi    string `json:"informasi"`
	Tugas        string `json:"tugas"`
	Kriteria     string `json:"kriteria"`
	KPrt         string `json:"kprt"`
	KBabysitter  string `json:"kbabysitter"`
	KSeniorcare  string `json:"kseniorcare"`
	KSupir       string `json:"ksupir"`
	KOfficeboy   string `json:"kofficeboy"`
	KTukangkebun string `json:"ktukangkebun"`
	TglPost      string `json:"tglpost"`
	NamaLengkap  string `json:"namalengkap"`
	JenisKelamin string `json:"jeniskelamin"`
	Kecamatan    string `json:"kecamatan"`
	Kota         string `json:"kota"`
}

func SimpanLowonganKerja(iduser int, judulloker string, gajiawal string, gajiakhir string,
	informasi string, tugas string, kriteria string, kprt string, kbabysitter string,
	kseniorcare string, ksupir string, kofficeboy string, ktukangkebun string, tglpost string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO lowongankerja (iduser, judulloker, gajiawal, gajiakhir, informasi, tugas, kriteria, kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, tglpost) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(iduser, judulloker, gajiawal, gajiakhir, informasi, tugas, kriteria, kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, tglpost)
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

func DataAllLowonganKerja() (Response, error) {
	var obj LowonganKerja
	var arrobj []LowonganKerja
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT lk.idloker, lk.iduser, lk.judulloker, lk.gajiawal, lk.gajiakhir, lk.informasi, lk.tugas, lk.kriteria, lk.kprt, lk.kbabysitter, lk.kseniorcare, lk.ksupir, lk.kofficeboy, lk.ktukangkebun, lk.tglpost, up.namalengkap, up.jeniskelamin, ud.kecamatan, ud.kota FROM lowongankerja lk JOIN userprofile up on lk.iduser = uP.iduser JOIN userdomisili ud on lk.iduser = ud.iduser"

	rows, err := con.Query(sqlStatemet)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdLoker, &obj.IdUser, &obj.JudulLoker, &obj.GajiAwal, &obj.GajiAkhir,
			&obj.Informasi, &obj.Tugas, &obj.Kriteria, &obj.KPrt, &obj.KBabysitter,
			&obj.KSeniorcare, &obj.KSupir, &obj.KOfficeboy, &obj.KTukangkebun, &obj.TglPost,
			&obj.NamaLengkap, &obj.JenisKelamin, &obj.Kecamatan, &obj.Kota)

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

func DataLowonganKerjaperUser(iduser int) (Response, error) {
	var obj LowonganKerja
	var arrobj []LowonganKerja
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT lk.idloker, lk.iduser, lk.judulloker, lk.gajiawal, lk.gajiakhir, lk.informasi, lk.tugas, lk.kriteria, lk.kprt, lk.kbabysitter, lk.kseniorcare, lk.ksupir, lk.kofficeboy, lk.ktukangkebun, lk.tglpost, up.namalengkap, up.jeniskelamin, ud.kecamatan, ud.kota FROM lowongankerja lk JOIN userprofile up on lk.iduser = uP.iduser JOIN userdomisili ud on lk.iduser = ud.iduser WHERE lk.iduser = ?"

	rows, err := con.Query(sqlStatemet, iduser)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdLoker, &obj.IdUser, &obj.JudulLoker, &obj.GajiAwal, &obj.GajiAkhir,
			&obj.Informasi, &obj.Tugas, &obj.Kriteria, &obj.KPrt, &obj.KBabysitter,
			&obj.KSeniorcare, &obj.KSupir, &obj.KOfficeboy, &obj.KTukangkebun, &obj.TglPost,
			&obj.NamaLengkap, &obj.JenisKelamin, &obj.Kecamatan, &obj.Kota)

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

func UpdateLowonganKerja(idloker int, judulloker string, gajiawal string, gajiakhir string, informasi string,
	tugas string, kriteria string, kprt string, kbabysitter string, kseniorcare string, ksupir string,
	kofficeboy string, ktukangkebun string, tglpost string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE lowongankerja SET judulloker=?, gajiawal=?, gajiakhir=?, informasi=?, tugas=?, kriteria=?,  kprt=?, kbabysitter=?, kseniorcare=?, ksupir=?, kofficeboy=?, ktukangkebun=?, tglpost=? WHERE idloker=?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(judulloker, gajiawal, gajiakhir, informasi, tugas, kriteria, kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, tglpost, idloker)
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

func SimpanLowonganKerjaSelesai(idloker int, iduser int, judulloker string, gajiawal string, gajiakhir string,
	informasi string, tugas string, kriteria string, kprt string, kbabysitter string,
	kseniorcare string, ksupir string, kofficeboy string, ktukangkebun string,
	tglpost string, alasan string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO lowongankerjaselesai (idloker, iduser, judulloker, gajiawal, gajiakhir, informasi, tugas, kriteria, kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, tglpost, alasan) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(idloker, iduser, judulloker, gajiawal, gajiakhir, informasi, tugas, kriteria, kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, tglpost, alasan)
	if err != nil {
		return res, err
	}

	defer stmt.Close()

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = map[string]int64{
		"rowsAffected": rowsAffected,
	}

	return res, nil
}

func DeleteDetailLowonganKerja(idloker int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE  FROM lowongankerja WHERE idloker=?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(idloker)
	if err != nil {
		return res, err
	}

	defer stmt.Close()
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
