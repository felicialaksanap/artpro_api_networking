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
	KPrt         string `json:"kprt"`
	KBabysitter  string `json:"kbabysitter"`
	KSeniorcare  string `json:"kseniorcare"`
	KSupir       string `json:"ksupir"`
	KOfficeboy   string `json:"kofficeboy"`
	KTukangkebun string `json:"ktukangkebun"`
	Hewan        string `json:"hewan"`
	Masak        string `json:"masak"`
	MabukJalan   string `json:"mabukjalan"`
	SepedaMotor  string `json:"sepedamotor"`
	Mobil        string `json:"mobil"`
	TKMenginap   string `json:"tkmenginap"`
	TKWarnen     string `json:"tkwarnen"`
	TglPost      string `json:"tglpost"`
	StatusLoker  string `json:"statusloker"`
	NamaLengkap  string `json:"namalengkap"`
	JenisKelamin string `json:"jeniskelamin"`
	Kecamatan    string `json:"kecamatan"`
	Kota         string `json:"kota"`
}

func SimpanLowonganKerja(iduser int, judulloker string, gajiawal string, gajiakhir string,
	informasi string, tugas string, kprt string, kbabysitter string,
	kseniorcare string, ksupir string, kofficeboy string,
	ktukangkebun string, hewan string, masak string, mabukjalan string,
	sepedamotor string, mobil string, tkmenginap string, tkwarnen string,
	tglpost string, statusloker string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO lowongankerja (iduser, judulloker, gajiawal, gajiakhir, informasi, tugas, kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, hewan, masak, mabukjalan, sepedamotor, mobil, tkmenginap, tkwarnen, tglpost, statusloker) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? ,? ,?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(iduser, judulloker, gajiawal, gajiakhir, informasi, tugas, kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, hewan, masak, mabukjalan, sepedamotor, mobil, tkmenginap, tkwarnen, tglpost, statusloker)
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

	sqlStatemet := "SELECT lk.idloker, lk.iduser, lk.judulloker, lk.gajiawal, lk.gajiakhir, lk.informasi, lk.tugas, lk.kprt, lk.kbabysitter, lk.kseniorcare, lk.ksupir, lk.kofficeboy, lk.ktukangkebun, lk.hewan, lk.masak, lk.mabukjalan, lk.sepedamotor, lk.mobil, lk.tkmenginap, lk.tkwarnen, lk.tglpost, lk.statusloker, up.namalengkap, up.jeniskelamin, ud.kecamatan, ud.kota FROM lowongankerja lk JOIN userprofile up on lk.iduser = uP.iduser JOIN userdomisili ud on lk.iduser = ud.iduser ORDER BY lk.tglpost DESC"

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

func DataLowonganKerjaperUser(iduser int) (Response, error) { // menampilkan semua loker yang dibuat oleh majikan
	var obj LowonganKerja
	var arrobj []LowonganKerja
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT lk.idloker, lk.iduser, lk.judulloker, lk.gajiawal, lk.gajiakhir, lk.informasi, lk.tugas, lk.kprt, lk.kbabysitter, lk.kseniorcare, lk.ksupir, lk.kofficeboy, lk.ktukangkebun, lk.hewan, lk.masak, lk.mabukjalan, lk.sepedamotor, lk.mobil, lk.tkmenginap, lk.tkwarnen, lk.tglpost, lk.statusloker, up.namalengkap, up.jeniskelamin, ud.kecamatan, ud.kota FROM lowongankerja lk JOIN userprofile up on lk.iduser = uP.iduser JOIN userdomisili ud on lk.iduser = ud.iduser WHERE lk.iduser = ? ORDER BY lk.idloker DESC"

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

func UpdateStatusLoker(idloker int, statusloker string, tglpost string) (Response, error) { // untuk update status loker aktif atau tidak aktif
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

func UpdateLowonganKerja(idloker int, judulloker string, gajiawal string, gajiakhir string,
	informasi string, tugas string, kprt string, kbabysitter string,
	kseniorcare string, ksupir string, kofficeboy string, ktukangkebun string,
	hewan string, masak string, mabukjalan string, sepedamotor string,
	mobil string, tkmenginap string, tkwarnen string, tglpost string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE lowongankerja SET judulloker=?, gajiawal=?, gajiakhir=?, informasi=?, tugas=?, kprt=?, kbabysitter=?, kseniorcare=?, ksupir=?, kofficeboy=?, ktukangkebun=?, hewan=?, masak=?, mabukjalan=?, sepedamotor=?, mobil=?, tkmenginap=?, tkwarnen=?, tglpost=? WHERE idloker=?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(judulloker, gajiawal, gajiakhir, informasi, tugas, kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, hewan, masak, mabukjalan, sepedamotor, mobil, tkmenginap, tkwarnen, tglpost, idloker)
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

func SimpanLowonganKerjaDetail(idloker int, iduser int, judulloker string, gajiawal string, gajiakhir string,
	informasi string, tugas string, kprt string, kbabysitter string,
	kseniorcare string, ksupir string, kofficeboy string, ktukangkebun string,
	hewan string, masak string, mabukjalan string, sepedamotor string,
	mobil string, tkmenginap string, tkwarnen string, tglmodif string,
	statusloker string, alasan string) (Response, error) { // Simpan data saat user majikan melakukan perubahan status loker aktif atau tidak aktif
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO lowongankerjadetail (idloker, iduser, judulloker, gajiawal, gajiakhir, informasi, tugas, kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, hewan, masak, mabukjalan, sepedamotor, mobil, tkmenginap, tkwarnen, tglmodif, statusloker, alasan) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(idloker, iduser, judulloker, gajiawal, gajiakhir, informasi, tugas, kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, hewan, masak, mabukjalan, sepedamotor, mobil, tkmenginap, tkwarnen, tglmodif, statusloker, alasan)
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
