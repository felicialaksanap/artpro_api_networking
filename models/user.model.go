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

func UpdateEmailUser(iduser int, email string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE userakun SET email=? WHERE iduser = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(email, iduser)
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

func UpdatePassUser(email string, password string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE userakun SET password=? WHERE email = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(password, email)
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

// UPDATE PROFILE
func UpdateProfileUser(iduser int, namalengkap string, jeniskelamin string, tempatlahir string, tanggallahir string, telephone string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE userprofile SET namalengkap = ?, jeniskelamin = ?, tempatlahir = ?, tanggallahir = ?, telephone = ? WHERE iduser = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(namalengkap, jeniskelamin, tempatlahir, tanggallahir, telephone, iduser)
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

type VerifikasiUser struct {
	IdUser           int    `json:"iduser"`
	NIK              string `json:"nik"`
	NamaLengkap      string `json:"namalengkap"`
	TempatLahir      string `json:"tempatlahir"`
	TanggalLahir     string `json:"tanggallahir"`
	JenisKelamin     string `json:"jeniskelamin"`
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

	sqlStatement := "SELECT uv.iduser, uv.nik, up.namalengkap, uv.tempatlahir, uv.tanggallahir, up.jeniskelamin, uv.alamat, uv.kecamatan, uv.kelurahan, uv.rt, uv.rw, uv.fotoktp, uv.selfiektp, uv.statusverifikasi FROM userverifikasi uv JOIN userprofile up on uv.iduser = up.iduser"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.IdUser, &obj.NIK, &obj.NamaLengkap, &obj.TempatLahir,
			&obj.TanggalLahir, &obj.TempatLahir, &obj.Alamat, &obj.Kecamatan,
			&obj.Kelurahan, &obj.RT, &obj.RW, &obj.FotoKTP, &obj.SelfieKTP,
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

func UpdateUserDomisili(iduser int, alamat string, kecamatan string, kelurahan string,
	provinsi string, kota string, longitude string, latitude string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE userdomisili SET alamat=?, kecamatan=?, kelurahan=?, provinsi=?, kota=?, longitude=?, latitude=? WHERE iduser=?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(alamat, kecamatan, kelurahan, provinsi, kota, longitude, latitude, iduser)
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

type DetailProfileART struct {
	IdUser             int    `json:"iduser"`
	PendidikanTerakhir string `json:"pendidikanterakhir"`
	BeratBadan         int    `json:"beratbadan"`
	TinggiBadan        int    `json:"tinggibadan"`
	Agama              string `json:"agama"`
	TKMenginap         int    `json:"tkmenginap"`
	TKWarnen           int    `json:"tkwarnen"`
	Hewan              int    `json:"hewan"`
	MabukJalan         int    `json:"mabukjalan"`
	SepedaMotor        int    `json:"sepedamotor"`
	Mobil              int    `json:"mobil"`
	Masak              int    `json:"masak"`
	SSingle            int    `json:"ssingle"`
	SMarried           int    `json:"smarried"`
}

func SimpanDetailProfileART(iduser int, pendidikanterakhir string, beratbadan int,
	tinggibadan int, agama string, tkmenginap int, tkwarnen int, hewan int,
	mabukjalan int, sepedamotor int, mobil int, masak int, ssingle int, smarried int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO detailprofileart (iduser, pendidikanterakhir, beratbadan, tinggibadan, agama, tkmenginap, tkwarnen, hewan, mabukjalan, sepedamotor, mobil, masak, ssingle, smarried) VALUES (?, ?, ?, ?, ?, ?, ? ,?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(iduser, pendidikanterakhir, beratbadan, tinggibadan, agama, tkmenginap, tkwarnen, hewan, mabukjalan, sepedamotor, mobil, masak, ssingle, smarried)
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
			&obj.Agama, &obj.TKMenginap, &obj.TKWarnen, &obj.Hewan, &obj.MabukJalan, &obj.SepedaMotor, &obj.Mobil,
			&obj.Masak, &obj.SSingle, &obj.SMarried)

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

func UpdateUserDetailProfileART(iduser int, pendidikanterakhir string, beratbadan int,
	tinggibadan int, agama string, tkmenginap int, tkwarnen int, hewan int,
	mabukjalan int, sepedamotor int, mobil int, masak int, ssingle int, smarried int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE detailprofileart SET pendidikanterakhir=?, beratbadan=?, tinggibadan=?, agama=?, tkmenginap=?, tkwarnen=?, hewan=?, mabukjalan=?, sepedamotor=?, mobil=?, masak=?, ssingle=?, smarried=? WHERE iduser=?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(pendidikanterakhir, beratbadan, tinggibadan, agama, tkmenginap, tkwarnen, hewan, mabukjalan, sepedamotor, mobil, masak, ssingle, smarried, iduser)
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

type DetailKerjaART struct {
	IdUser       int    `json:"iduser"`
	KPrt         int    `json:"kprt"`
	KBabysitter  int    `json:"kbabysitter"`
	KSeniorcare  int    `json:"kseniorcare"`
	KSupir       int    `json:"ksupir"`
	KOfficeboy   int    `json:"kofficeboy"`
	KTukangkebun int    `json:"ktukangkebun"`
	Pengalaman   string `json:"pengalaman"`
	GajiAwal     string `json:"gajiawal"`
	GajiAkhir    string `json:"gajiakhir"`
}

func SimpanDetailKerjaART(iduser int, kprt int, kbabysitter int, kseniorcare int,
	ksupir int, kofficeboy int, ktukangkebun int,
	pengalaman string, gajiawal string, gajiakhir string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO detailkerjaart (iduser, kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, pengalaman, gajiawal, gajiakhir) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(iduser, kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, pengalaman, gajiawal, gajiakhir)
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
		err = rows.Scan(&obj.IdUser, &obj.KPrt, &obj.KBabysitter,
			&obj.KSeniorcare, &obj.KSupir, &obj.KOfficeboy,
			&obj.KTukangkebun, &obj.Pengalaman,
			&obj.GajiAwal, &obj.GajiAkhir)

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

func DataListKerjaPerKategori(kategori string) (Response, error) {
	var obj DetailKerjaART
	var arrobj []DetailKerjaART
	var res Response

	con := db.CreateCon()

	if kategori == "prt" {
		sqlStatemet := "SELECT * FROM detailkerjaart WHERE kprt=1"

		rows, err := con.Query(sqlStatemet)
		defer rows.Close()
		if err != nil {
			log.Printf(err.Error())
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.IdUser, &obj.KPrt, &obj.KBabysitter,
				&obj.KOfficeboy, &obj.KSupir, &obj.KOfficeboy,
				&obj.KTukangkebun, &obj.Pengalaman,
				&obj.GajiAwal, &obj.GajiAkhir)

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
	} else if kategori == "babysitter" {
		sqlStatemet := "SELECT * FROM detailkerjaart WHERE kbabysitter=1"

		rows, err := con.Query(sqlStatemet)
		defer rows.Close()
		if err != nil {
			log.Printf(err.Error())
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.IdUser, &obj.KPrt, &obj.KBabysitter,
				&obj.KOfficeboy, &obj.KSupir, &obj.KOfficeboy,
				&obj.KTukangkebun, &obj.Pengalaman,
				&obj.GajiAwal, &obj.GajiAkhir)

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
	} else if kategori == "seniorcare" {
		sqlStatemet := "SELECT * FROM detailkerjaart WHERE kseniorcare=1"

		rows, err := con.Query(sqlStatemet)
		defer rows.Close()
		if err != nil {
			log.Printf(err.Error())
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.IdUser, &obj.KPrt, &obj.KBabysitter,
				&obj.KOfficeboy, &obj.KSupir, &obj.KOfficeboy,
				&obj.KTukangkebun, &obj.Pengalaman,
				&obj.GajiAwal, &obj.GajiAkhir)

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
	} else if kategori == "supir" {
		sqlStatemet := "SELECT * FROM detailkerjaart WHERE ksupir=1"

		rows, err := con.Query(sqlStatemet)
		defer rows.Close()
		if err != nil {
			log.Printf(err.Error())
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.IdUser, &obj.KPrt, &obj.KBabysitter,
				&obj.KOfficeboy, &obj.KSupir, &obj.KOfficeboy,
				&obj.KTukangkebun, &obj.Pengalaman,
				&obj.GajiAwal, &obj.GajiAkhir)

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
	} else if kategori == "officeboy" {
		sqlStatemet := "SELECT * FROM detailkerjaart WHERE kofficeboy=1"

		rows, err := con.Query(sqlStatemet)
		defer rows.Close()
		if err != nil {
			log.Printf(err.Error())
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.IdUser, &obj.KPrt, &obj.KBabysitter,
				&obj.KOfficeboy, &obj.KSupir, &obj.KOfficeboy,
				&obj.KTukangkebun, &obj.Pengalaman,
				&obj.GajiAwal, &obj.GajiAkhir)

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
	} else if kategori == "tukangkebun" {
		sqlStatemet := "SELECT * FROM detailkerjaart WHERE ktukangkebun=1"

		rows, err := con.Query(sqlStatemet)
		defer rows.Close()
		if err != nil {
			log.Printf(err.Error())
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.IdUser, &obj.KPrt, &obj.KBabysitter,
				&obj.KOfficeboy, &obj.KSupir, &obj.KOfficeboy,
				&obj.KTukangkebun, &obj.Pengalaman,
				&obj.GajiAwal, &obj.GajiAkhir)

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
	}

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
		err = rows.Scan(
			&obj.IdUser, &obj.KPrt, &obj.KBabysitter,
			&obj.KSeniorcare, &obj.KSupir, &obj.KOfficeboy,
			&obj.KTukangkebun, &obj.Pengalaman,
			&obj.GajiAwal, &obj.GajiAkhir)

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

func UpdateUserDetailKerja(iduser int, kprt int, kbabysitter int,
	kseniorcare int, ksupir int, kofficeboy int,
	ktukangkebun int, pengalaman string, gajiawal string, gajiakhir string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE detailkerjaart SET kprt=?, kbabysitter=?, kseniorcare=?, ksupir=?, kofficeboy=?, ktukangkebun=?, pengalaman=?, gajiawal=?, gajiakhir=? WHERE iduser=?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, pengalaman, gajiawal, gajiakhir, iduser)
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

//type KontakUser struct {
//	IdKontak    int    `json:"idkontak"`
//	IdMajikan   int    `json:"idmajikan"`
//	IdART       int    `json:"idart"`
//	WaktuKontak string `json:"waktukontak"`
//	Darimana    string `json:"darimana"`
//}

type TotalKontakART struct {
	IdART          int `json:"idart"`
	TotalKontakART int `json:"totalkontakart"`
}

func GetTotalKontakART(idart int) (Response, error) {
	var obj TotalKontakART
	var arrobj []TotalKontakART
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT idart, COUNT(idmajikan) as totalkontak FROM kontakart WHERE idart = ?"

	rows, err := con.Query(sqlStatemet, idart)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdART, &obj.TotalKontakART)

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

func SimpanKontakuser(idmajikan int, idart int, waktukontak string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO kontakart (idmajikan, idart, waktukontak) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(idmajikan, idart, waktukontak)
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

//func DataListKontakByMajikan(idmajikan int) (Response, error) {
//	var obj KontakUser
//	var arrobj []KontakUser
//	var res Response
//
//	con := db.CreateCon()
//
//	sqlStatemet := "SELECT * FROM kontakuser WHERE idmajikan=?"
//
//	rows, err := con.Query(sqlStatemet, idmajikan)
//
//	defer rows.Close()
//
//	if err != nil {
//		log.Printf(err.Error())
//		return res, err
//	}
//
//	for rows.Next() {
//		err = rows.Scan(&obj.IdKontak, &obj.IdMajikan, &obj.IdART, &obj.WaktuKontak, &obj.Darimana)
//
//		if err != nil {
//			log.Printf(err.Error())
//			return res, err
//		}
//
//		arrobj = append(arrobj, obj)
//	}
//	log.Printf("berhasil")
//	res.Status = http.StatusOK
//	res.Message = "Sukses"
//	res.Data = arrobj
//
//	return res, nil
//}
//
//func DataListKontakByART(idart int) (Response, error) {
//	var obj KontakUser
//	var arrobj []KontakUser
//	var res Response
//
//	con := db.CreateCon()
//
//	sqlStatemet := "SELECT * FROM kontakuser WHERE idart=?"
//
//	rows, err := con.Query(sqlStatemet, idart)
//
//	defer rows.Close()
//
//	if err != nil {
//		log.Printf(err.Error())
//		return res, err
//	}
//
//	for rows.Next() {
//		err = rows.Scan(&obj.IdKontak, &obj.IdMajikan, &obj.IdART, &obj.WaktuKontak, &obj.Darimana)
//
//		if err != nil {
//			log.Printf(err.Error())
//			return res, err
//		}
//
//		arrobj = append(arrobj, obj)
//	}
//	log.Printf("berhasil")
//	res.Status = http.StatusOK
//	res.Message = "Sukses"
//	res.Data = arrobj
//
//	return res, nil
//}

type Penilaian struct {
	IdART      int     `json:"idart"`
	Etika      float64 `json:"etika"`
	Estetika   float64 `json:"estetika"`
	Kebersihan float64 `json:"kebersihan"`
	Kerapian   float64 `json:"kerapian"`
	Kecepatan  float64 `json:"kecepatan"`
	Review     string  `json:"review"`
}

func SimpanPenilaian(idart int, idmajikan int, etika int, estetika int,
	kebersihan int, kerapian int, kecepatan int, avgnilai float64, review string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO penilaian (idart, idmajikan, estetika, etika, kebersihan, kecepatan, kerapian, avgnilai, review) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(idart, idmajikan, estetika, etika, kebersihan, kecepatan, kerapian, avgnilai, review)
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

	sqlStatemet := "SELECT idart, AVG(estetika) as estetika, AVG(etika) as etika, AVG(kebersihan) as kebersihan, AVG(kecepatan) as kecepatan, AVG(kerapian) as kerapian, review FROM penilaian WHERE idart = ?"

	rows, err := con.Query(sqlStatemet, idart)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdART, &obj.Etika, &obj.Estetika,
			&obj.Kebersihan, &obj.Kerapian, &obj.Kecepatan,
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

type AvgART struct {
	IdART   int     `json:"idart"`
	Average float64 `json:"average"`
}

func GetAvgNilaiART(idart int) (Response, error) {
	var obj AvgART
	var arrobj []AvgART
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT idart, AVG(avgnilai) as average FROM penilaian WHERE idart = ?"

	rows, err := con.Query(sqlStatemet, idart)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdART, &obj.Average)

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
