package models

import (
	"artpro_api_networking/db"
	"log"
	"net/http"
)

type LowonganKerja struct {
	IdLoker     int    `json:"idloker"`
	IdUser      int    `json:"iduser"`
	Kategori    string `json:"kategori"`
	Informasi   string `json:"informasi"`
	UraianTugas string `json:"uraiantugas"`
	Keahlian    string `json:"keahlian"`
	TipeKerja   string `json:"tipekerja"`
	GajiAwal    string `json:"gajiawal"`
	GajiAkhir   string `json:"gajiakhir"`
}

func SimpanLowonganKerja(iduser int, kategori string, informasi string, uraiantugas string,
	keahlian string, tipekerja string, gajiawal string, gajiakhir string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO lowongankerja (iduser, kategori, informasi, uraiantugas, keahlian, tipekerja, gajiawal, gajiakhir) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(iduser, kategori, informasi, uraiantugas, keahlian, tipekerja, gajiawal, gajiakhir)
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

	sqlStatemet := "SELECT * FROM lowongankerja"

	rows, err := con.Query(sqlStatemet)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdLoker, &obj.IdUser, &obj.Kategori, &obj.Informasi, &obj.UraianTugas, &obj.Keahlian, &obj.TipeKerja, &obj.GajiAwal, &obj.GajiAkhir)

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

func DataDetailLowonganKerja(idloker int) (Response, error) {
	var obj LowonganKerja
	var arrobj []LowonganKerja
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT * FROM lowongankerja WHERE idloker=?"

	rows, err := con.Query(sqlStatemet, idloker)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdLoker, &obj.IdUser, &obj.Kategori, &obj.Informasi, &obj.UraianTugas, &obj.Keahlian, &obj.TipeKerja, &obj.GajiAwal, &obj.GajiAkhir)

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

func SimpanLowonganKerjaSelesai(idloker int, iduser int, kategori string, informasi string, uraiantugas string,
	keahlian string, tipekerja string, gajiawal string, gajiakhir string, alasan string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO lowongankerjaselesai (idloker, iduser, kategori, informasi, uraiantugas, keahlian, tipekerja, gajiawal, gajiakhir, alasan) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(idloker, iduser, kategori, informasi, uraiantugas, keahlian, tipekerja, gajiawal, gajiakhir, alasan)
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

	sqlStatement := "DELETE FROM lowongankerja WHERE idloker=?"
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