package models

import (
	"artpro_api_networking/db"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func UploadFoto(request *http.Request, id string, folder string) (Response, error) {
	var res Response

	con := db.CreateCon()

	// LIMIT 10 MB
	request.ParseMultipartForm(10 * 1024 * 1024)

	file, handler, err := request.FormFile("photo")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fmt.Println("File Info")
	fmt.Println("File Name: ", handler.Filename)
	fmt.Println("File Size: ", handler.Size)
	fmt.Println("File Type: ", handler.Header.Get("Content-Type"))

	// Upload File
	tempFile, err2 := os.CreateTemp("uploads/"+folder, "upload-*.png")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer tempFile.Close()

	fileBytes, err3 := io.ReadAll(file)
	if err3 != nil {
		fmt.Println(err3)
	}
	tempFile.Write(fileBytes)
	tempFile.Close()
	fmt.Println("Selesai")

	newPath := "uploads/" + folder + "/" + folder + "-" + id + ".png"
	fmt.Println("newpath: " + newPath)

	renameFile := os.Rename(tempFile.Name(), newPath)
	if renameFile != nil {
		log.Fatal(renameFile)
	}

	if folder == "fotoktp" {
		sqlStatement := "UPDATE userverifikasi SET fotoktp = ? WHERE iduser = ?"

		stmt, err4 := con.Prepare(sqlStatement)
		if err4 != nil {
			return res, err
		}

		result, err5 := stmt.Exec(newPath, id)
		if err5 != nil {
			return res, err5
		}

		rowsAffected, err6 := result.RowsAffected()
		if err6 != nil {
			return res, err6
		}

		res.Status = http.StatusOK
		res.Message = "Sukses Upload Foto"
		res.Data = map[string]int64{
			"rows": rowsAffected,
		}
	} else if folder == "selfiektp" {
		sqlStatement := "UPDATE userverifikasi SET selfiektp = ? WHERE iduser = ?"

		stmt, err4 := con.Prepare(sqlStatement)
		if err4 != nil {
			return res, err
		}

		result, err5 := stmt.Exec(newPath, id)
		if err5 != nil {
			return res, err5
		}

		rowsAffected, err6 := result.RowsAffected()
		if err6 != nil {
			return res, err6
		}

		res.Status = http.StatusOK
		res.Message = "Sukses Upload Foto"
		res.Data = map[string]int64{
			"rows": rowsAffected,
		}
	} else if folder == "profpic" {
		sqlStatement := "UPDATE userprofile SET profilepicpath = ? WHERE iduser = ?"

		stmt, err4 := con.Prepare(sqlStatement)
		if err4 != nil {
			return res, err
		}

		result, err5 := stmt.Exec(newPath, id)
		if err5 != nil {
			return res, err5
		}

		rowsAffected, err6 := result.RowsAffected()
		if err6 != nil {
			return res, err6
		}

		res.Status = http.StatusOK
		res.Message = "Sukses Upload Foto"
		res.Data = map[string]int64{
			"rows": rowsAffected,
		}
	}

	return res, err
}

func GetPhoto(path string, id string) string {
	result := "uploads/" + path + "/" + path + "-" + id + ".png"
	return result
}

type AkunUser struct {
	IdUser     int    `json:"iduser"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	StatusUser string `json:"statususer"`
}

func SimpanAkunUser(email string, password string, statususer string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO userakun (email, password, statususer) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(email, password, statususer)
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

func DataAkunUser(email string, password string) (Response, error) {
	var obj AkunUser
	var arrobj []AkunUser
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT * FROM userakun WHERE email = ? AND password = ?"

	rows, err := con.Query(sqlStatemet, email, password)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdUser, &obj.Email, &obj.Password, &obj.StatusUser)

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

type ProfileUser struct {
	IdUser         int    `json:"iduser"`
	NamaLengkap    string `json:"namalengkap"`
	JenisKelamin   string `json:"jeniskelamin"`
	TempatLahir    string `json:"tempatlahir"`
	TanggalLahir   string `json:"tanggallahir"`
	Telephone      string `json:"telephone"`
	ProfilePicPath string `json:"profilepicpath"`
}

func SimpanProfileUser(iduser int, namalengkap string, jeniskelamin string, tempatlahir string, tanggallahir string, telephone string, profilepicpath string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO userprofile (iduser, namalengkap, jeniskelamin, tempatlahir, tanggallahir, telephone, profilepicpath) VALUES (?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(iduser, namalengkap, jeniskelamin, tempatlahir, tanggallahir, telephone, profilepicpath)
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

func DataProfileUser(iduser int) (Response, error) {
	var obj ProfileUser
	var arrobj []ProfileUser
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT * FROM userprofile WHERE iduser=?"

	rows, err := con.Query(sqlStatemet, iduser)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdUser, &obj.NamaLengkap, &obj.JenisKelamin, &obj.TempatLahir,
			&obj.TanggalLahir, &obj.Telephone, &obj.ProfilePicPath)

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

type VerifikasiUser struct {
	IdUser           int    `json:"iduser"`
	NIK              string `json:"nik"`
	TempatLahir      string `json:"tempatlahir"`
	TanggalLahir     string `json:"tanggallahir"`
	Alamat           string `json:"alamat"`
	Kecamatan        string `json:"kecamatan"`
	Kelurahan        string `json:"kelurahan"`
	RT               int    `json:"rt"`
	RW               int    `json:"rw"`
	FotoKTP          string `json:"fotoktp"`
	SelfieKTP        string `json:"selfiektp"`
	StatusVerifikasi string `json:"statusverifikasi"`
}

func SimpanDataVerifikasi(iduser int, nik string, tempatlahir string, tanggallahir string,
	alamat string, kecamatan string, kelurahan string, rt int, rw int,
	fotoktp string, selfiektp string, statusverifikasi string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO userverifikasi (iduser, nik, tempatlahir, tanggallahir, alamat, kecamatan, kelurahan, rt, rw, fotoktp, selfiektp, statusverifikasi) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(iduser, nik, tempatlahir, tanggallahir, alamat, kecamatan, kelurahan, rt, rw, fotoktp, selfiektp, statusverifikasi)
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

func DataAllVerifikasi() (Response, error) {
	var obj VerifikasiUser
	var arrobj []VerifikasiUser
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM userverifikasi"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.IdUser, &obj.NIK, &obj.TempatLahir, &obj.TanggalLahir, &obj.Alamat,
			&obj.Kecamatan, &obj.Kelurahan, &obj.RT, &obj.RW, &obj.FotoKTP, &obj.SelfieKTP,
			&obj.StatusVerifikasi)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = arrobj

	return res, nil
}

func DataVerifikasiUser(iduser int) (Response, error) {
	var obj VerifikasiUser
	var arrobj []VerifikasiUser
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM userverifikasi WHERE iduser=?"

	rows, err := con.Query(sqlStatement, iduser)
	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.IdUser, &obj.NIK, &obj.TempatLahir, &obj.TanggalLahir, &obj.Alamat,
			&obj.Kecamatan, &obj.Kelurahan, &obj.RT, &obj.RW, &obj.FotoKTP, &obj.SelfieKTP,
			&obj.StatusVerifikasi)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = arrobj

	return res, nil
}

func UpdateDataVerifikasi(iduser int, statusverifikasi string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE userverifikasi SET statusverifikasi = ? WHERE iduser = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(statusverifikasi, iduser)
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

type DomisiliUser struct {
	IdUser    int    `json:"iduser"`
	Alamat    string `json:"alamat"`
	Kecamatan string `json:"kecamatan"`
	Kelurahan string `json:"kelurahan"`
	Provinsi  string `json:"provinsi"`
	Kota      string `json:"kota"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

func SimpanDomisiliUser(iduser int, alamat string, kecamatan string, kelurahan string,
	provinsi string, kota string, longitude string, latitude string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO userdomisili (iduser, alamat, kecamatan, kelurahan, provinsi, kota, longitude, latitude) VALUES (?, ?, ?, ?, ?, ?, ? ,?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(iduser, alamat, kecamatan, kelurahan, provinsi, kota, longitude, latitude)
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

func DataUserDomisili(iduser int) (Response, error) {
	var obj DomisiliUser
	var arrobj []DomisiliUser
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT * FROM userdomisili WHERE iduser=?"

	rows, err := con.Query(sqlStatemet, iduser)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdUser, &obj.Alamat, &obj.Kecamatan, &obj.Kelurahan,
			&obj.Provinsi, &obj.Kota, &obj.Longitude, &obj.Latitude)

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

type DetailProfileART struct {
	IdUser             int    `json:"iduser"`
	PendidikanTerakhir string `json:"pendidikanterakhir"`
	BeratBadan         int    `json:"beratbadan"`
	TinggiBadan        int    `json:"tinggibadan"`
	Agama              string `json:"agama"`
	TipeKerja          string `json:"tipekerja"`
	Hewan              string `json:"hewan"`
	MabukJalan         string `json:"mabukjalan"`
	SepedaMotor        string `json:"sepedamotor"`
	Mobil              string `json:"mobil"`
	Masak              string `json:"masak"`
}

func SimpanDetailProfileART(iduser int, pendidikanterakhir string, beratbadan int,
	tinggibadan int, agama string, tipekerja string, hewan string,
	mabukjalan string, sepedamotor string, mobil string, masak string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO detailprofileart (iduser, pendidikanterakhir, beratbadan, tinggibadan, agama, tipekerja, hewan, mabukjalan, sepedamotor, mobil, masak) VALUES (?, ?, ?, ?, ?, ?, ? ,?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(iduser, pendidikanterakhir, beratbadan, tinggibadan, agama, tipekerja, hewan, mabukjalan, sepedamotor, mobil, masak)
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

func DataUserDetailProfileART(iduser int) (Response, error) {
	var obj DetailProfileART
	var arrobj []DetailProfileART
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT * FROM detailprofileart WHERE iduser=?"

	rows, err := con.Query(sqlStatemet, iduser)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdUser, &obj.PendidikanTerakhir, &obj.BeratBadan, &obj.TinggiBadan,
			&obj.Agama, &obj.TipeKerja, &obj.Hewan, &obj.MabukJalan, &obj.SepedaMotor, &obj.Mobil,
			&obj.Masak)

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

type DetailKerjaART struct {
	IdUser     int    `json:"iduser"`
	Kategori   string `json:"kategori"`
	Pengalaman string `json:"pengalaman"`
	GajiAwal   string `json:"gajiawal"`
	GajiAkhir  string `json:"gajiakhir"`
}

func SimpanDetailKerjaART(iduser int, kategori string, pengalaman string, gajiawal string, gajiakhir string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO detailkerjaart (iduser, kategori, pengalaman, gajiawal, gajiakhir) VALUES (?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(iduser, kategori, pengalaman, gajiawal, gajiakhir)
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

func DataListKerjaPerKategori(kategori string) (Response, error) {
	var obj DetailKerjaART
	var arrobj []DetailKerjaART
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT * FROM detailkerjaart WHERE kategori=?"

	rows, err := con.Query(sqlStatemet, kategori)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdUser, &obj.Kategori, &obj.Pengalaman, &obj.GajiAwal, &obj.GajiAkhir)

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

func DataAllDetailKerjaART() (Response, error) {
	var obj DetailKerjaART
	var arrobj []DetailKerjaART
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT * FROM detailkerjaart"

	rows, err := con.Query(sqlStatemet)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdUser, &obj.Kategori, &obj.Pengalaman, &obj.GajiAwal, &obj.GajiAkhir)

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

func DataUserDetailKerjaART(iduser int) (Response, error) {
	var obj DetailKerjaART
	var arrobj []DetailKerjaART
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT * FROM detailkerjaart WHERE iduser=?"

	rows, err := con.Query(sqlStatemet, iduser)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdUser, &obj.Kategori, &obj.Pengalaman, &obj.GajiAwal, &obj.GajiAkhir)

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

type KontakUser struct {
	IdKontak    int    `json:"idkontak"`
	IdMajikan   int    `json:"idmajikan"`
	IdART       int    `json:"idart"`
	WaktuKontak string `json:"waktukontak"`
	Darimana    string `json:"darimana"`
}

func SimpanKontakuser(idmajikan int, idart int, waktukontak string, darimana string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO kontakuser (idmajikan, idart, waktukontak, darimana) VALUES (?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(idmajikan, idart, waktukontak, darimana)
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

func DataListKontakByMajikan(idmajikan int) (Response, error) {
	var obj KontakUser
	var arrobj []KontakUser
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT * FROM kontakuser WHERE idmajikan=?"

	rows, err := con.Query(sqlStatemet, idmajikan)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdKontak, &obj.IdMajikan, &obj.IdART, &obj.WaktuKontak, &obj.Darimana)

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

func DataListKontakByART(idart int) (Response, error) {
	var obj KontakUser
	var arrobj []KontakUser
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT * FROM kontakuser WHERE idart=?"

	rows, err := con.Query(sqlStatemet, idart)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdKontak, &obj.IdMajikan, &obj.IdART, &obj.WaktuKontak, &obj.Darimana)

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

type Penilaian struct {
	IdNilai    int    `json:"idnilai"`
	IdART      int    `json:"idart"`
	IdMajikan  int    `json:"idmajikan"`
	Etika      int    `json:"etika"`
	Estetika   int    `json:"estetika"`
	Kebersihan int    `json:"kebersihan"`
	Kerapian   int    `json:"kerapian"`
	Kecepatan  int    `json:"kecepatan"`
	Review     string `json:"review"`
}

func SimpanPenilaian(idart int, idmajikan int, etika int, estetika int,
	kebersihan int, kerapian int, kecepatan int, review string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO penilaian (idart, idmajikan, etika, estetika, kebersihan, kerapian, kecepatan, review) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(idart, idmajikan, etika, estetika, kebersihan, kerapian, kecepatan, review)
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

func DataPenilaianART(idart int) (Response, error) {
	var obj Penilaian
	var arrobj []Penilaian
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT * FROM penilaian WHERE idart=?"

	rows, err := con.Query(sqlStatemet, idart)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdNilai, &obj.IdART, &obj.IdMajikan, &obj.Etika,
			&obj.Estetika, &obj.Kebersihan, &obj.Kerapian, &obj.Kecepatan,
			&obj.Review)

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

type SertifikatPelatihan struct {
	IdUser     int    `json:"iduser"`
	SertifPath string `json:"sertifpath"`
}

func SimpanSertifikatPelatihan(iduser int, sertifpath string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO sertifikatpelatihan (iduser, sertifpath) VALUES (?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(iduser, sertifpath)
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

func DataSertifPelatihanUser(iduser int) (Response, error) {
	var obj SertifikatPelatihan
	var arrobj []SertifikatPelatihan
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT * FROM sertifikatpelatihan WHERE iduser=?"

	rows, err := con.Query(sqlStatemet, iduser)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdUser, &obj.SertifPath)

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
