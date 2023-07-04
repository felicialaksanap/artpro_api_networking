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

type DataLokerFilter struct {
	IdLoker          int     `json:"idloker"`
	IdPencari        string  `json:"idpencari"`
	Innerproduct     float64 `json:"innerproduct"`
	X                float64 `json:"x"`
	Y                float64 `json:"y"`
	Cosinesimilarity float64 `json:"cosinesimilarity"`
	Jarak            float64 `json:"jarakawal"`
	Iduser           int     `json:"iduser"`
	Judul            string  `json:"judul"`
	Gajiawal         int     `json:"gajiawal"`
	Gajiakhir        int     `json:"gajiakhir"`
	Informasi        string  `json:"informasi"`
	Tugas            string  `json:"tugas"`
	Kprt             int     `json:"kprt"`
	Kbabysitter      int     `json:"kbabysitter"`
	Kseniorcare      int     `json:"kseniorcare"`
	Ksupir           int     `json:"ksupir"`
	Kofficeboy       int     `json:"kofficeboy"`
	Ktukangkebun     int     `json:"ktukangkebun"`
	Hewan            int     `json:"hewan"`
	Masak            int     `json:"masak"`
	Mabukjalan       int     `json:"mabukjalan"`
	Sepedamotor      int     `json:"sepedamotor"`
	Mobil            int     `json:"mobil"`
	Tkmenginap       int     `json:"tkmenginap"`
	Tkwarnen         int     `json:"tkwarnen"`
	Ssingle          int     `json:"ssingle"`
	Smarried         int     `json:"smarried"`
	Tglpost          string  `json:"tglpost"`
	Namalengkap      string  `json:"namalengkap"`
	Jeniskelamin     string  `json:"jeniskelamin"`
	Kecamatan        string  `json:"kecamatan"`
	Kota             string  `json:"kota"`
}

func DataLokerbyFilter(idart string, kprt int, kbabysitter int, kseniorcare int,
	ksupir int, kofficeboy int, ktukangkebun int, hewan int,
	masak int, mabukjalan int, sepedamotor int, mobil int,
	tkmenginap int, tkwarnen int, ssingle int, smarried int,
	gajiawal int, gajiakhir int, jarak float64, updatejarak string) (Response, error) {
	var obj DataLokerFilter
	var arrobj []DataLokerFilter
	var res Response

	con := db.CreateCon()

	name := "tabletemplk"
	tablename := name + idart

	sqlStatement := ""

	if updatejarak == "false" {
		sqlStatement = "CREATE TABLE IF NOT EXISTS " + tablename +
			" (idloker varchar(5), kprt int, kbabysitter int, kseniorcare int," +
			" ksupir int, kofficeboy int, ktukangkebun int, hewan int, masak int," +
			" mabukjalan int, sepedamotor int, mobil int, tkmenginap int, tkwarnen int," +
			" ssingle int, smarried int, gajiawal double, gajiakhir double)"

		_, err := con.Exec(sqlStatement)
		if err != nil {
			log.Printf(err.Error())
			return res, nil
		}

		sqlStatement = "INSERT INTO " + tablename +
			" SELECT idloker, kprt, kbabysitter, kseniorcare, ksupir," +
			" kofficeboy, ktukangkebun, hewan, masak, mabukjalan," +
			" sepedamotor, mobil, tkmenginap, tkwarnen, ssingle," +
			" smarried, gajiawal, gajiakhir FROM lowongankerja" +
			" WHERE statusloker = 1"

		_, err = con.Exec(sqlStatement)
		if err != nil {
			log.Printf(err.Error())
			return res, nil
		}

		sqlStatement = "ALTER TABLE " + tablename + " ADD jarak double NOT NULL"
		_, err = con.Exec(sqlStatement)
		if err != nil {
			log.Printf(err.Error())
			return res, nil
		}
	}

	sqlStatement = "INSERT INTO " + tablename +
		" (idloker, kprt, kbabysitter, kseniorcare, ksupir," +
		" kofficeboy, ktukangkebun, hewan, masak, mabukjalan," +
		" sepedamotor, mobil, tkmenginap, tkwarnen, ssingle," +
		" smarried, gajiawal, gajiakhir, jarak) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := con.Exec(sqlStatement, idart, kprt, kbabysitter, kseniorcare, ksupir,
		kofficeboy, ktukangkebun, hewan, masak, mabukjalan,
		sepedamotor, mobil, tkmenginap, tkwarnen, ssingle, smarried,
		gajiawal, gajiakhir, jarak)
	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	sqlStatement = "UPDATE " + tablename +
		" SET gajiawal = (gajiawal-0)/(4000000-0)," +
		" gajiakhir = (gajiakhir-0)/(4000000-0)," +
		" jarak = (jarak-0)/(10-0)"
	_, err = con.Exec(sqlStatement)
	if err != nil {
		log.Printf(err.Error())
		return res, nil
	}

	sqlStatement = "SELECT a.idloker as idloker, b.idloker as idpencari," +
		" ((a.kprt * b.kprt) + (a.kbabysitter * b.kbabysitter) + (a.kseniorcare * b.kseniorcare) +" +
		" (a.ksupir * b.ksupir) + (a.kofficeboy * b.kofficeboy) * (a.ktukangkebun * b.ktukangkebun) +" +
		" (a.hewan * b.hewan) + (a.masak * b.masak) + (a.mabukjalan * b.mabukjalan) + (a.sepedamotor * b.sepedamotor) +" +
		" (a.mobil * b.mobil) + (a.tkmenginap * b.tkmenginap) + (a.tkwarnen * b.tkwarnen) + (a.ssingle * b.ssingle) +" +
		" (a.smarried * b.smarried) + (a.gajiawal * b.gajiawal) + (a.gajiakhir * b.gajiakhir)) as innerproduct," +
		" ((a.kprt * a.kprt) + (a.kbabysitter * a.kbabysitter) + (a.kseniorcare * a.kseniorcare) +" +
		" (a.ksupir * a.ksupir) + (a.kofficeboy * a.kofficeboy) * (a.ktukangkebun * a.ktukangkebun) +" +
		" (a.hewan * a.hewan) + (a.masak * a.masak) + (a.mabukjalan * a.mabukjalan) + (a.sepedamotor * a.sepedamotor) +" +
		" (a.mobil * a.mobil) + (a.tkmenginap * a.tkmenginap) + (a.tkwarnen * a.tkwarnen) + (a.ssingle * a.ssingle) +" +
		" (a.smarried * a.smarried) + (a.gajiawal * a.gajiawal) + (a.gajiakhir * a.gajiakhir)) as x," +
		" ((b.kprt * b.kprt) + (b.kbabysitter * b.kbabysitter) + (b.kseniorcare * b.kseniorcare) + (b.ksupir * b.ksupir) +" +
		" (b.kofficeboy * b.kofficeboy) + (b.ktukangkebun * b.ktukangkebun) + (b.hewan * b.hewan) + (b.masak * b.masak) +" +
		" (b.mabukjalan * b.mabukjalan) + (b.sepedamotor * b.sepedamotor) + (b.mobil * b.mobil) +" +
		" (b.tkmenginap * b.tkmenginap) + (b.tkwarnen * b.tkwarnen) + (b.ssingle * b.ssingle) + (b.smarried * b.smarried) +" +
		" (b.gajiawal * b.gajiawal) + (b.gajiakhir * b.gajiakhir)) as y," +
		" ((a.kprt * b.kprt) + (a.kbabysitter * b.kbabysitter) + (a.kseniorcare * b.kseniorcare) + (a.ksupir * b.ksupir) +" +
		" (a.kofficeboy * b.kofficeboy) * (a.ktukangkebun * b.ktukangkebun) + (a.hewan * b.hewan) + (a.masak * b.masak) +" +
		" (a.mabukjalan * b.mabukjalan) + (a.sepedamotor * b.sepedamotor) + (a.mobil * b.mobil) + (a.tkmenginap * b.tkmenginap) +" +
		" (a.tkwarnen * b.tkwarnen) + (a.ssingle * b.ssingle) + (a.smarried * b.smarried) + (a.gajiawal * b.gajiawal) +" +
		" (a.gajiakhir * b.gajiakhir)) /" +
		" sqrt(((a.kprt * a.kprt) + (a.kbabysitter * a.kbabysitter) + (a.kseniorcare * a.kseniorcare) + (a.ksupir * a.ksupir) +" +
		" (a.kofficeboy * a.kofficeboy) * (a.ktukangkebun * a.ktukangkebun) + (a.hewan * a.hewan) + (a.masak * a.masak) +" +
		" (a.mabukjalan * a.mabukjalan) + (a.sepedamotor * a.sepedamotor) + (a.mobil * a.mobil) + (a.tkmenginap * a.tkmenginap) +" +
		" (a.tkwarnen * a.tkwarnen) + (a.ssingle * a.ssingle) + (a.smarried * a.smarried) + (a.gajiawal * a.gajiawal) +" +
		" (a.gajiakhir * a.gajiakhir)) * ((b.kprt * b.kprt) + (b.kbabysitter * b.kbabysitter) + (b.kseniorcare * b.kseniorcare) +" +
		" (b.ksupir * b.ksupir) + (b.kofficeboy * b.kofficeboy) + (b.ktukangkebun * b.ktukangkebun) + (b.hewan * b.hewan) +" +
		" (b.masak * b.masak) + (b.mabukjalan * b.mabukjalan) + (b.sepedamotor * b.sepedamotor) + (b.mobil * b.mobil) +" +
		" (b.tkmenginap * b.tkmenginap) + (b.tkwarnen * b.tkwarnen) + (b.ssingle * b.ssingle) + (b.smarried * b.smarried) +" +
		" (b.gajiawal * b.gajiawal) + (b.gajiakhir * b.gajiakhir))) as cosinesimilarity," +
		" (a.jarak * 10) as jarakawal," +
		" lk.iduser, judulloker, lk.gajiawal, lk.gajiakhir, informasi, tugas, lk.kprt, lk.kbabysitter, lk.kseniorcare," +
		" lk.ksupir, lk.kofficeboy, lk.ktukangkebun, lk.hewan, lk.masak, lk.mabukjalan, lk.sepedamotor, lk.mobil, lk.tkmenginap," +
		" lk.tkwarnen, lk.ssingle, lk.smarried, lk.tglpost, up.namalengkap, jeniskelamin, ud.kecamatan, kota" +
		" FROM " + tablename + " a" +
		" JOIN " + tablename + " b" +
		" JOIN lowongankerja lk ON a.idloker = lk.idloker" +
		" JOIN userprofile up ON lk.iduser = up.iduser" +
		" JOIN userdomisili ud ON lk.iduser = ud.iduser" +
		" WHERE b.idloker = ? AND a.idloker != ?" +
		" GROUP BY a.idloker" +
		" ORDER BY cosinesimilarity desc"
	rows, err := con.Query(sqlStatement, idart, idart)
	defer rows.Close()
	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdLoker, &obj.IdPencari, &obj.Innerproduct, &obj.X, &obj.Y, &obj.Cosinesimilarity, &obj.Jarak,
			&obj.Iduser, &obj.Judul, &obj.Gajiawal, &obj.Gajiakhir, &obj.Informasi,
			&obj.Tugas, &obj.Kprt, &obj.Kbabysitter, &obj.Kseniorcare, &obj.Ksupir,
			&obj.Kofficeboy, &obj.Ktukangkebun, &obj.Hewan, &obj.Masak, &obj.Mabukjalan, &obj.Sepedamotor,
			&obj.Mobil, &obj.Tkmenginap, &obj.Tkwarnen, &obj.Ssingle, &obj.Smarried, &obj.Tglpost,
			&obj.Namalengkap, &obj.Jeniskelamin, &obj.Kecamatan, &obj.Kota)

		if err != nil {
			log.Printf(err.Error())
			return res, err
		}

		arrobj = append(arrobj, obj)
	}
	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = arrobj

	sqlStatement = "DROP TABLE " + tablename
	_, err = con.Exec(sqlStatement)
	if err != nil {
		log.Printf(err.Error())
		return res, nil
	}

	return res, err
}

func CreateAndCopyTable(idart string) (Response, error) {
	var res Response

	con := db.CreateCon()

	name := "tabletemplk"
	tablename := name + idart

	sqlStatement := "CREATE TABLE IF NOT EXISTS " + tablename +
		" (idloker varchar(5), kprt int, kbabysitter int,  kseniorcare int," +
		" ksupir int, kofficeboy int, ktukangkebun int, hewan int, masak int," +
		" mabukjalan int, sepedamotor int, mobil int, tkmenginap int, tkwarnen int," +
		" ssingle int, smarried int, gajiawal double, gajiakhir double)"

	_, err := con.Exec(sqlStatement)
	if err != nil {
		log.Printf(err.Error())
		return res, nil
	}

	sqlStatement = "INSERT INTO " + tablename +
		" SELECT idloker, kprt, kbabysitter, kseniorcare, ksupir," +
		" kofficeboy, ktukangkebun, hewan, masak, mabukjalan," +
		" sepedamotor, mobil, tkmenginap, tkwarnen, ssingle," +
		" smarried, gajiawal, gajiakhir FROM lowongankerja" +
		" WHERE statusloker = 1"

	_, err = con.Exec(sqlStatement)
	if err != nil {
		log.Printf(err.Error())
		return res, nil
	}

	sqlStatement = "ALTER TABLE " + tablename + " ADD jarak double NOT NULL"
	_, err = con.Exec(sqlStatement)
	if err != nil {
		log.Printf(err.Error())
		return res, nil
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = "Berhasil buat dan copy table"

	return res, err
}

func UpdateJarakLoker(idart string, idloker int, jarak float64) (Response, error) {
	var res Response

	con := db.CreateCon()

	name := "tabletemplk"
	tablename := name + idart

	sqlStatement := "UPDATE " + tablename + " SET jarak=? WHERE idloker=?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(jarak, idloker)
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

	return res, err
}
