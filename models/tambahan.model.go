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
