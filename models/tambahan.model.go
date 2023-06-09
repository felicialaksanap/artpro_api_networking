package models

import (
	"artpro_api_networking/db"
	"log"
	"net/http"
)

func SimpanBerita(judul string, isi string, url string, tglpost string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO beritatips (judul, isi, url, tglpost) VALUES (?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(judul, isi, url, tglpost)
	if err != nil {
		return res, err
	}

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

type BeritaObject struct {
	IdBerita int    `json:"idberita"`
	Judul    string `json:"judul"`
	Isi      string `json:"isi"`
	URL      string `json:"url"`
	Tglpost  string `json:"tglpost"`
}

func DataAllBerita() (Response, error) {
	var obj BeritaObject
	var arrobj []BeritaObject
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM beritatips ORDER BY tglpost DESC"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdBerita, &obj.Judul, &obj.Isi, &obj.URL, &obj.Tglpost)

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

func SimpanInfo(judul string, isi string, url string, tglpost string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO infopelatihan (judul, isi, url, tglpost) VALUES (?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(judul, isi, url, tglpost)
	if err != nil {
		return res, err
	}

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

type InfoObject struct {
	Idinfo  int    `json:"idinfo"`
	Judul   string `json:"judul"`
	Isi     string `json:"isi"`
	URL     string `json:"url"`
	Tglpost string `json:"tglpost"`
}

func DataALLInfo() (Response, error) {
	var obj InfoObject
	var arrobj []InfoObject
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM infopelatihan ORDER BY tglpost DESC"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Idinfo, &obj.Judul, &obj.Isi, &obj.URL, &obj.Tglpost)

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

func SimpanPengaduan(idmajikan int, idart int, isipengaduan string, penyelesaian string, tglpengaduan string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO pengaduan (idmajikan, idart, isipengaduan, penyelesaian, tglpengaduan) VALUES (?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(idmajikan, idart, isipengaduan, penyelesaian, tglpengaduan)
	if err != nil {
		return res, err
	}

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

func UpdatePenyelesaian(idloker int, penyelesaian string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE pengaduan SET penyelesaian = ? WHERE idpengaduan = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(penyelesaian, idloker)
	if err != nil {
		return res, err
	}

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

type PengaduanObject struct {
	Idpengaduan  int    `json:"idpengaduan"`
	IdMajikan    int    `json:"idmajikan"`
	NamaMajikan  string `json:"namamajikan"`
	IdArt        int    `json:"idart"`
	NamaART      string `json:"namaart"`
	IsiPengaduan string `json:"isipengaduan"`
	Penyelesaian string `json:"penyelesaian"`
	Tglpengaduan string `json:"tglpengaduan"`
}

func DataALLPengaduan() (Response, error) {
	var obj PengaduanObject
	var arrobj []PengaduanObject
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT p.idpengaduan, p.idmajikan, um.namalengkap as namamajikan," +
		" p.idart, ua.namalengkap as namaart, p.isipengaduan, p.penyelesaian, p.tglpengaduan" +
		" FROM pengaduan p" +
		" JOIN userprofile um ON p.idmajikan = um.iduser" +
		" JOIN userprofile ua ON p.idart = ua.iduser"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Idpengaduan, &obj.IdMajikan, &obj.NamaMajikan, &obj.IdArt, &obj.NamaART,
			&obj.IsiPengaduan, &obj.Penyelesaian, &obj.Tglpengaduan)

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
