package models

import (
	"artpro_api_networking/db"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
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
	Alasan           string `json:"alasan"`
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

	sqlStatement := "SELECT uv.iduser, uv.nik, up.namalengkap, uv.tempatlahir, uv.tanggallahir, up.jeniskelamin, uv.alamat, uv.kecamatan, uv.kelurahan, uv.rt, uv.rw, uv.fotoktp, uv.selfiektp, uv.statusverifikasi, uv.alasan FROM userverifikasi uv JOIN userprofile up on uv.iduser = up.iduser"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.IdUser, &obj.NIK, &obj.NamaLengkap, &obj.TempatLahir,
			&obj.TanggalLahir, &obj.TempatLahir, &obj.Alamat, &obj.Kecamatan,
			&obj.Kelurahan, &obj.RT, &obj.RW, &obj.FotoKTP, &obj.SelfieKTP,
			&obj.StatusVerifikasi, &obj.Alasan)

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

func UpdateDataVerifikasi(iduser int, statusverifikasi string, alasan string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE userverifikasi SET statusverifikasi = ?, alasan = ? WHERE iduser = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(statusverifikasi, alasan, iduser)
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

type LongLatObj struct {
	Iduser    int    `json:"iduser"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

func DataLongLatUser(iduser int) (Response, error) {
	var obj LongLatObj
	var arrobj []LongLatObj
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT iduser, longitude, latitude FROM userdomisili WHERE iduser=?"

	rows, err := con.Query(sqlStatemet, iduser)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Iduser, &obj.Longitude, &obj.Latitude)

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
	AIslam             int    `json:"aislam"`
	AKatolik           int    `json:"akatolik"`
	AKristen           int    `json:"akristen"`
	AHindu             int    `json:"ahindu"`
	ABuddha            int    `json:"abuddha"`
	AKonghucu          int    `json:"akonghucu"`
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
	tinggibadan int, aislam int, akatolik int, akristen int, ahindu int, abuddha int, akonghucu int,
	tkmenginap int, tkwarnen int, hewan int, mabukjalan int, sepedamotor int, mobil int, masak int,
	ssingle int, smarried int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO detailprofileart (iduser, pendidikanterakhir, beratbadan, tinggibadan, aislam, akatolik, akristen, ahindu, abuddha, akonghucu, tkmenginap, tkwarnen, hewan, mabukjalan, sepedamotor, mobil, masak, ssingle, smarried) VALUES (?, ?, ?, ?, ?, ?, ? ,?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(iduser, pendidikanterakhir, beratbadan, tinggibadan, aislam, akatolik, akristen, ahindu, abuddha, akonghucu, tkmenginap, tkwarnen, hewan, mabukjalan, sepedamotor, mobil, masak, ssingle, smarried)
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
			&obj.AIslam, &obj.AKatolik, &obj.AKristen, &obj.AHindu, &obj.ABuddha, &obj.AKonghucu,
			&obj.TKMenginap, &obj.TKWarnen, &obj.Hewan, &obj.MabukJalan, &obj.SepedaMotor, &obj.Mobil,
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
	tinggibadan int, aislam int, akatolik int, akristen int, ahindu int, abuddha int, akonghucu int, tkmenginap int, tkwarnen int, hewan int,
	mabukjalan int, sepedamotor int, mobil int, masak int, ssingle int, smarried int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE detailprofileart SET pendidikanterakhir=?, beratbadan=?, tinggibadan=?, aislam=?, akatolik=?, akristen=?, ahindu=?, abuddha=?, akonghucu=?, tkmenginap=?, tkwarnen=?, hewan=?, mabukjalan=?, sepedamotor=?, mobil=?, masak=?, ssingle=?, smarried=? WHERE iduser=?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(pendidikanterakhir, beratbadan, tinggibadan, aislam, akatolik, akristen, ahindu, abuddha, akonghucu, tkmenginap, tkwarnen, hewan, mabukjalan, sepedamotor, mobil, masak, ssingle, smarried, iduser)
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
	GajiAwal     int    `json:"gajiawal"`
	GajiAkhir    int    `json:"gajiakhir"`
}

func SimpanDetailKerjaART(iduser int, kprt int, kbabysitter int, kseniorcare int,
	ksupir int, kofficeboy int, ktukangkebun int,
	pengalaman string, gajiawal int, gajiakhir int) (Response, error) {
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

type DatabyKategori struct {
	IdART        int     `json:"idart"`
	NamaLengkap  string  `json:"namalengkap"`
	JenisKelamin string  `json:"jeniskelamin"`
	TempatLahir  string  `json:"tempatlahir"`
	TanggalLahir string  `json:"tanggallahir"`
	Telephone    string  `json:"telephone"`
	ProfPicPath  string  `json:"profpicpath"`
	Pendidikan   string  `json:"pendidikan"`
	BeratBadan   int     `json:"beratbadan"`
	TinggiBadan  int     `json:"tinggibadan"`
	AIslam       int     `json:"aislam"`
	AKatolik     int     `json:"akatolik"`
	AKristen     int     `json:"akristen"`
	AHindu       int     `json:"ahindu"`
	ABuddha      int     `json:"abuddha"`
	AKonghucu    int     `json:"akonghucu"`
	TkMenginap   int     `json:"tkmenginap"`
	TkWarnen     int     `json:"tkwarnen"`
	Hewan        int     `json:"hewan"`
	MabukJalan   int     `json:"mabukjalan"`
	SepedaMotor  int     `json:"sepedamotor"`
	Mobil        int     `json:"mobil"`
	Masak        int     `json:"masak"`
	SSingle      int     `json:"ssingle"`
	SMarried     int     `json:"smarried"`
	KPrt         int     `json:"kprt"`
	KBabySitter  int     `json:"kbabysitter"`
	KSeniorCare  int     `json:"kseniorcare"`
	KSupir       int     `json:"ksupir"`
	KOfficeBoy   int     `json:"kofficeboy"`
	KTukangKebun int     `json:"ktukangkebun"`
	Pengalaman   string  `json:"pengalaman"`
	GajiAwal     int     `json:"gajiawal"`
	GajiAkhir    int     `json:"gajiakhir"`
	Latitude     string  `json:"latitude"`
	Longitude    string  `json:"longitude"`
	Rating       float64 `json:"rating"`
}

func DataARTbyKategori(kategori string) (Response, error) {
	var obj DatabyKategori
	var arrobj []DatabyKategori
	var res Response

	con := db.CreateCon()

	if kategori == "prt" {
		sqlStatement := "SELECT up.iduser as idart, namalengkap, jeniskelamin, tempatlahir, " +
			"tanggallahir, telephone, profilepicpath, dpa.pendidikanterakhir, beratbadan, " +
			"tinggibadan, aislam, akatolik, akristen, ahindu, abuddha, akonghucu, tkmenginap, " +
			"tkwarnen, hewan, mabukjalan, sepedamotor, mobil, masak, ssingle, smarried, dka.kprt, " +
			"kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, pengalaman, gajiawal, " +
			"gajiakhir, ud.latitude, longitude, AVG(p.rating) as rating " +
			"FROM userprofile up " +
			"JOIN detailprofileart dpa on up.iduser = dpa.iduser " +
			"JOIN detailkerjaart dka on up.iduser = dka.iduser " +
			"JOIN userdomisili ud on up.iduser = ud.iduser " +
			"JOIN penilaian p on up.iduser = p.idart " +
			"WHERE kprt = 1 " +
			"GROUP BY idart " +
			"ORDER BY rating DESC"

		rows, err := con.Query(sqlStatement)
		defer rows.Close()
		if err != nil {
			log.Printf(err.Error())
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.IdART, &obj.NamaLengkap, &obj.JenisKelamin, &obj.TempatLahir,
				&obj.TanggalLahir, &obj.Telephone, &obj.ProfPicPath, &obj.Pendidikan,
				&obj.BeratBadan, &obj.TinggiBadan, &obj.AIslam, &obj.AKatolik, &obj.AKristen,
				&obj.AHindu, &obj.ABuddha, &obj.AKonghucu, &obj.TkMenginap,
				&obj.TkWarnen, &obj.Hewan, &obj.MabukJalan, &obj.SepedaMotor, &obj.Mobil,
				&obj.Masak, &obj.SSingle, &obj.SMarried, &obj.KPrt, &obj.KBabySitter,
				&obj.KSeniorCare, &obj.KSupir, &obj.KOfficeBoy, &obj.KTukangKebun,
				&obj.Pengalaman, &obj.GajiAwal, &obj.GajiAkhir, &obj.Latitude, &obj.Longitude, &obj.Rating)

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
		sqlStatement := "SELECT up.iduser as idart, namalengkap, jeniskelamin, tempatlahir, " +
			"tanggallahir, telephone, profilepicpath, dpa.pendidikanterakhir, beratbadan, " +
			"tinggibadan, aislam, akatolik, akristen, ahindu, abuddha, akonghucu, tkmenginap, " +
			"tkwarnen, hewan, mabukjalan, sepedamotor, mobil, masak, ssingle, smarried, " +
			"dka.kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, pengalaman, " +
			"gajiawal, gajiakhir, ud.latitude, longitude, AVG(p.rating) as rating " +
			"FROM userprofile up " +
			"JOIN detailprofileart dpa on up.iduser = dpa.iduser " +
			"JOIN detailkerjaart dka on up.iduser = dka.iduser " +
			"JOIN userdomisili ud on up.iduser = ud.iduser " +
			"JOIN penilaian p on up.iduser = p.idart " +
			"WHERE kbabysitter = 1 " +
			"GROUP BY idart " +
			"ORDER BY rating DESC"

		rows, err := con.Query(sqlStatement)
		defer rows.Close()
		if err != nil {
			log.Printf(err.Error())
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.IdART, &obj.NamaLengkap, &obj.JenisKelamin, &obj.TempatLahir,
				&obj.TanggalLahir, &obj.Telephone, &obj.ProfPicPath, &obj.Pendidikan,
				&obj.BeratBadan, &obj.TinggiBadan, &obj.AIslam, &obj.AKatolik, &obj.AKristen,
				&obj.AHindu, &obj.ABuddha, &obj.AKonghucu, &obj.TkMenginap,
				&obj.TkWarnen, &obj.Hewan, &obj.MabukJalan, &obj.SepedaMotor, &obj.Mobil,
				&obj.Masak, &obj.SSingle, &obj.SMarried, &obj.KPrt, &obj.KBabySitter,
				&obj.KSeniorCare, &obj.KSupir, &obj.KOfficeBoy, &obj.KTukangKebun,
				&obj.Pengalaman, &obj.GajiAwal, &obj.GajiAkhir, &obj.Latitude, &obj.Longitude, &obj.Rating)

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
		sqlStatement := "SELECT up.iduser as idart, namalengkap, jeniskelamin, tempatlahir, " +
			"tanggallahir, telephone, profilepicpath, dpa.pendidikanterakhir, beratbadan, " +
			"tinggibadan, aislam, akatolik, akristen, ahindu, abuddha, akonghucu, tkmenginap, " +
			"tkwarnen, hewan, mabukjalan, sepedamotor, mobil, masak, ssingle, smarried, dka.kprt, " +
			"kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, pengalaman, gajiawal, " +
			"gajiakhir, ud.latitude, longitude, AVG(p.rating) as rating " +
			"FROM userprofile up " +
			"JOIN detailprofileart dpa on up.iduser = dpa.iduser " +
			"JOIN detailkerjaart dka on up.iduser = dka.iduser " +
			"JOIN userdomisili ud on up.iduser = ud.iduser " +
			"JOIN penilaian p on up.iduser = p.idart " +
			"WHERE kseniorcare = 1 " +
			"GROUP BY idart " +
			"ORDER BY rating DESC"

		rows, err := con.Query(sqlStatement)
		defer rows.Close()
		if err != nil {
			log.Printf(err.Error())
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.IdART, &obj.NamaLengkap, &obj.JenisKelamin, &obj.TempatLahir,
				&obj.TanggalLahir, &obj.Telephone, &obj.ProfPicPath, &obj.Pendidikan,
				&obj.BeratBadan, &obj.TinggiBadan, &obj.AIslam, &obj.AKatolik, &obj.AKristen,
				&obj.AHindu, &obj.ABuddha, &obj.AKonghucu, &obj.TkMenginap,
				&obj.TkWarnen, &obj.Hewan, &obj.MabukJalan, &obj.SepedaMotor, &obj.Mobil,
				&obj.Masak, &obj.SSingle, &obj.SMarried, &obj.KPrt, &obj.KBabySitter,
				&obj.KSeniorCare, &obj.KSupir, &obj.KOfficeBoy, &obj.KTukangKebun,
				&obj.Pengalaman, &obj.GajiAwal, &obj.GajiAkhir, &obj.Latitude, &obj.Longitude, &obj.Rating)

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
		sqlStatement := "SELECT up.iduser as idart, namalengkap, jeniskelamin, tempatlahir, " +
			"tanggallahir, telephone, profilepicpath, dpa.pendidikanterakhir, beratbadan, " +
			"tinggibadan, aislam, akatolik, akristen, ahindu, abuddha, akonghucu, " +
			"tkmenginap, tkwarnen, hewan, mabukjalan, sepedamotor, mobil, masak, ssingle, " +
			"smarried, dka.kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, " +
			"pengalaman, gajiawal, gajiakhir, ud.latitude, longitude, AVG(p.rating) as rating " +
			"FROM userprofile up " +
			"JOIN detailprofileart dpa on up.iduser = dpa.iduser " +
			"JOIN detailkerjaart dka on up.iduser = dka.iduser " +
			"JOIN userdomisili ud on up.iduser = ud.iduser " +
			"JOIN penilaian p on up.iduser = p.idart " +
			"WHERE ksupir = 1 G" +
			"ROUP BY idart " +
			"ORDER BY rating DESC"

		rows, err := con.Query(sqlStatement)
		defer rows.Close()
		if err != nil {
			log.Printf(err.Error())
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.IdART, &obj.NamaLengkap, &obj.JenisKelamin, &obj.TempatLahir,
				&obj.TanggalLahir, &obj.Telephone, &obj.ProfPicPath, &obj.Pendidikan,
				&obj.BeratBadan, &obj.TinggiBadan, &obj.AIslam, &obj.AKatolik, &obj.AKristen,
				&obj.AHindu, &obj.ABuddha, &obj.AKonghucu, &obj.TkMenginap,
				&obj.TkWarnen, &obj.Hewan, &obj.MabukJalan, &obj.SepedaMotor, &obj.Mobil,
				&obj.Masak, &obj.SSingle, &obj.SMarried, &obj.KPrt, &obj.KBabySitter,
				&obj.KSeniorCare, &obj.KSupir, &obj.KOfficeBoy, &obj.KTukangKebun,
				&obj.Pengalaman, &obj.GajiAwal, &obj.GajiAkhir, &obj.Latitude, &obj.Longitude, &obj.Rating)

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
		sqlStatement := "SELECT up.iduser as idart, namalengkap, jeniskelamin, tempatlahir, " +
			"tanggallahir, telephone, profilepicpath, dpa.pendidikanterakhir, beratbadan, " +
			"tinggibadan, aislam, akatolik, akristen, ahindu, abuddha, akonghucu, " +
			"tkmenginap, tkwarnen, hewan, mabukjalan, sepedamotor, mobil, masak, ssingle, " +
			"smarried, dka.kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, " +
			"pengalaman, gajiawal, gajiakhir, ud.latitude, longitude, AVG(p.rating) as rating " +
			"FROM userprofile up " +
			"JOIN detailprofileart dpa on up.iduser = dpa.iduser " +
			"JOIN detailkerjaart dka on up.iduser = dka.iduser " +
			"JOIN userdomisili ud on up.iduser = ud.iduser " +
			"JOIN penilaian p on up.iduser = p.idart " +
			"WHERE kofficeboy = 1 " +
			"GROUP BY idart " +
			"ORDER BY rating DESC"

		rows, err := con.Query(sqlStatement)
		defer rows.Close()
		if err != nil {
			log.Printf(err.Error())
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.IdART, &obj.NamaLengkap, &obj.JenisKelamin, &obj.TempatLahir,
				&obj.TanggalLahir, &obj.Telephone, &obj.ProfPicPath, &obj.Pendidikan,
				&obj.BeratBadan, &obj.TinggiBadan, &obj.AIslam, &obj.AKatolik, &obj.AKristen,
				&obj.AHindu, &obj.ABuddha, &obj.AKonghucu, &obj.TkMenginap,
				&obj.TkWarnen, &obj.Hewan, &obj.MabukJalan, &obj.SepedaMotor, &obj.Mobil,
				&obj.Masak, &obj.SSingle, &obj.SMarried, &obj.KPrt, &obj.KBabySitter,
				&obj.KSeniorCare, &obj.KSupir, &obj.KOfficeBoy, &obj.KTukangKebun,
				&obj.Pengalaman, &obj.GajiAwal, &obj.GajiAkhir, &obj.Latitude, &obj.Longitude, &obj.Rating)

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
		sqlStatement := "SELECT up.iduser as idart, namalengkap, jeniskelamin, tempatlahir, " +
			"tanggallahir, telephone, profilepicpath, dpa.pendidikanterakhir, beratbadan, " +
			"tinggibadan, aislam, akatolik, akristen, ahindu, abuddha, akonghucu, " +
			"tkmenginap, tkwarnen, hewan, mabukjalan, sepedamotor, mobil, masak, ssingle, " +
			"smarried, dka.kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, " +
			"pengalaman, gajiawal, gajiakhir, ud.latitude, longitude, AVG(p.rating) as rating " +
			"FROM userprofile up " +
			"JOIN detailprofileart dpa on up.iduser = dpa.iduser " +
			"JOIN detailkerjaart dka on up.iduser = dka.iduser " +
			"JOIN userdomisili ud on up.iduser = ud.iduser " +
			"JOIN penilaian p on up.iduser = p.idart " +
			"WHERE ktukangkebun = 1 " +
			"GROUP BY idart " +
			"ORDER BY rating DESC"

		rows, err := con.Query(sqlStatement)
		defer rows.Close()
		if err != nil {
			log.Printf(err.Error())
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.IdART, &obj.NamaLengkap, &obj.JenisKelamin, &obj.TempatLahir,
				&obj.TanggalLahir, &obj.Telephone, &obj.ProfPicPath, &obj.Pendidikan,
				&obj.BeratBadan, &obj.TinggiBadan, &obj.AIslam, &obj.AKatolik, &obj.AKristen,
				&obj.AHindu, &obj.ABuddha, &obj.AKonghucu, &obj.TkMenginap,
				&obj.TkWarnen, &obj.Hewan, &obj.MabukJalan, &obj.SepedaMotor, &obj.Mobil,
				&obj.Masak, &obj.SSingle, &obj.SMarried, &obj.KPrt, &obj.KBabySitter,
				&obj.KSeniorCare, &obj.KSupir, &obj.KOfficeBoy, &obj.KTukangKebun,
				&obj.Pengalaman, &obj.GajiAwal, &obj.GajiAkhir, &obj.Latitude, &obj.Longitude, &obj.Rating)

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

func DataARTbyId(idart int) (Response, error) {
	var obj DatabyKategori
	var arrobj []DatabyKategori
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT up.iduser as idart, namalengkap, jeniskelamin, tempatlahir, " +
		"tanggallahir, telephone, profilepicpath, dpa.pendidikanterakhir, beratbadan, " +
		"tinggibadan, aislam, akatolik, akristen, ahindu, abuddha, akonghucu, " +
		"tkmenginap, tkwarnen, hewan, mabukjalan, sepedamotor, mobil, masak, ssingle, " +
		"smarried, dka.kprt, kbabysitter, kseniorcare, ksupir, kofficeboy, ktukangkebun, " +
		"pengalaman, gajiawal, gajiakhir, ud.latitude, longitude, AVG(p.rating) as rating " +
		"FROM userprofile up " +
		"JOIN detailprofileart dpa on up.iduser = dpa.iduser " +
		"JOIN detailkerjaart dka on up.iduser = dka.iduser " +
		"JOIN userdomisili ud on up.iduser = ud.iduser " +
		"JOIN penilaian p on up.iduser = p.idart " +
		"WHERE idart = ? " +
		"GROUP BY idart " +
		"ORDER BY rating DESC"

	rows, err := con.Query(sqlStatement, idart)
	defer rows.Close()
	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdART, &obj.NamaLengkap, &obj.JenisKelamin, &obj.TempatLahir,
			&obj.TanggalLahir, &obj.Telephone, &obj.ProfPicPath, &obj.Pendidikan,
			&obj.BeratBadan, &obj.TinggiBadan, &obj.AIslam, &obj.AKatolik, &obj.AKristen,
			&obj.AHindu, &obj.ABuddha, &obj.AKonghucu, &obj.TkMenginap,
			&obj.TkWarnen, &obj.Hewan, &obj.MabukJalan, &obj.SepedaMotor, &obj.Mobil,
			&obj.Masak, &obj.SSingle, &obj.SMarried, &obj.KPrt, &obj.KBabySitter,
			&obj.KSeniorCare, &obj.KSupir, &obj.KOfficeBoy, &obj.KTukangKebun,
			&obj.Pengalaman, &obj.GajiAwal, &obj.GajiAkhir, &obj.Latitude, &obj.Longitude, &obj.Rating)

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

type DatabyFilter struct {
	IdART            int     `json:"idart"`
	InnerProduct     float64 `json:"innerproduct"`
	X                float64 `json:"x"`
	Y                float64 `json:"y"`
	CosineSimilarity float64 `json:"cosinesimilarity"`
	NamaLengkap      string  `json:"namalengkap"`
	JenisKelamin     string  `json:"jeniskelamin"`
	TempatLahir      string  `json:"tempatlahir"`
	TanggalLahir     string  `json:"tanggallahir"`
	Telephone        string  `json:"telephone"`
	Pendidikan       string  `json:"pendidikan"`
	BeratBadan       int     `json:"beratbadan"`
	TinggiBadan      int     `json:"tinggibadan"`
	AIslam           int     `json:"aislam"`
	AKatolik         int     `json:"akatolik"`
	AKristen         int     `json:"akristen"`
	AHindu           int     `json:"ahindu"`
	ABuddha          int     `json:"abuddha"`
	AKonghucu        int     `json:"akonghucu"`
	TKMenginap       int     `json:"tkmenginap"`
	TKWarnen         int     `json:"tkwarnen"`
	Hewan            int     `json:"hewan"`
	MabukJalan       int     `json:"mabukjalan"`
	SepedaMotor      int     `json:"sepedamotor"`
	Mobil            int     `json:"mobil"`
	Masak            int     `json:"masak"`
	SSingle          int     `json:"ssingle"`
	SMarried         int     `json:"smarried"`
	KPrt             int     `json:"kprt"`
	KBabysitter      int     `json:"kbabysitter"`
	KSeniorcare      int     `json:"kseniorcare"`
	KSupir           int     `json:"ksupir"`
	KOfficeboy       int     `json:"kofficeboy"`
	KTukangkebun     int     `json:"ktukangkebun"`
	Pengalaman       string  `json:"pengalaman"`
	Gajiawal         int     `json:"gajiawal"`
	Gajiakhir        int     `json:"gajiakhir"`
	Jarak            float64 `json:"jarak"`
	Rating           float64 `json:"rating"`
}

func PrepareTableDataFK(kategori string, idmajikan int) (Response, error) {
	var res Response

	con := db.CreateCon()

	name := "tabletemp"
	id := strconv.Itoa(idmajikan)
	tablename := name + id

	// create tabletemp
	sqlStatement := "CREATE TABLE IF NOT EXISTS " + tablename +
		" (iduser int, aislam int, akatolik int, akristen int," +
		" ahindu int, abuddha int, akonghucu int, hewan int," +
		" mabukjalan int, sepedamotor int, mobil int, masak int," +
		" ssingle int, smarried int, gajiawal double, gajiakhir double)"
	_, err := con.Exec(sqlStatement)
	if err != nil {
		log.Printf(err.Error())
		return res, nil
	}

	// copy data by kategori
	if kategori == "prt" {
		sqlStatement = "INSERT INTO " + tablename +
			" SELECT dp.iduser, aislam, akatolik, akristen, " +
			"ahindu, abuddha, akonghucu, hewan, mabukjalan, " +
			"sepedamotor, mobil, masak, ssingle, smarried, " +
			"dk.gajiawal, dk.gajiakhir " +
			"FROM detailprofileart dp " +
			"JOIN detailkerjaart dk ON dp.iduser = dk.iduser " +
			"WHERE kprt = 1"
		_, err = con.Exec(sqlStatement)
		if err != nil {
			log.Printf(err.Error())
			return res, nil
		}
	} else if kategori == "babysitter" {
		sqlStatement = "INSERT INTO " + tablename +
			" SELECT dp.iduser, aislam, akatolik, akristen, " +
			"ahindu, abuddha, akonghucu, hewan, mabukjalan, " +
			"sepedamotor, mobil, masak, ssingle, smarried, " +
			"dk.gajiawal, dk.gajiakhir " +
			"FROM detailprofileart dp " +
			"JOIN detailkerjaart dk ON dp.iduser = dk.iduser " +
			"WHERE kbabysitter = 1"
		_, err = con.Exec(sqlStatement)
		if err != nil {
			log.Printf(err.Error())
			return res, nil
		}
	} else if kategori == "seniorcare" {
		sqlStatement = "INSERT INTO " + tablename +
			" SELECT dp.iduser, aislam, akatolik, akristen, " +
			"ahindu, abuddha, akonghucu, hewan, mabukjalan, " +
			"sepedamotor, mobil, masak, ssingle, smarried, " +
			"dk.gajiawal, dk.gajiakhir " +
			"FROM detailprofileart dp " +
			"JOIN detailkerjaart dk ON dp.iduser = dk.iduser " +
			"WHERE kseniorcare = 1"
		_, err = con.Exec(sqlStatement)
		if err != nil {
			log.Printf(err.Error())
			return res, nil
		}
	} else if kategori == "supir" {
		sqlStatement = "INSERT INTO " + tablename +
			" SELECT dp.iduser, aislam, akatolik, akristen, " +
			"ahindu, abuddha, akonghucu, hewan, mabukjalan, " +
			"sepedamotor, mobil, masak, ssingle, smarried, " +
			"dk.gajiawal, dk.gajiakhir " +
			"FROM detailprofileart dp " +
			"JOIN detailkerjaart dk ON dp.iduser = dk.iduser " +
			"WHERE ksupir = 1"
		_, err = con.Exec(sqlStatement)
		if err != nil {
			log.Printf(err.Error())
			return res, nil
		}
	} else if kategori == "officeboy" {
		sqlStatement = "INSERT INTO " + tablename +
			" SELECT dp.iduser, aislam, akatolik, akristen, " +
			"ahindu, abuddha, akonghucu, hewan, mabukjalan, " +
			"sepedamotor, mobil, masak, ssingle, smarried, " +
			"dk.gajiawal, dk.gajiakhir " +
			"FROM detailprofileart dp " +
			"JOIN detailkerjaart dk ON dp.iduser = dk.iduser " +
			"WHERE kofficeboy = 1"
		_, err = con.Exec(sqlStatement)
		if err != nil {
			log.Printf(err.Error())
			return res, nil
		}
	} else if kategori == "tukangkebun" {
		sqlStatement = "INSERT INTO " + tablename +
			" SELECT dp.iduser, aislam, akatolik, akristen, " +
			"ahindu, abuddha, akonghucu, hewan, mabukjalan, " +
			"sepedamotor, mobil, masak, ssingle, smarried, " +
			"dk.gajiawal, dk.gajiakhir " +
			"FROM detailprofileart dp " +
			"JOIN detailkerjaart dk ON dp.iduser = dk.iduser " +
			"WHERE ktukangkebun = 1"
		_, err = con.Exec(sqlStatement)
		if err != nil {
			log.Printf(err.Error())
			return res, nil
		}
	}

	sqlStatement = "ALTER TABLE " + tablename + " ADD jarak double NOT NULL"

	_, err = con.Exec(sqlStatement)
	if err != nil {
		log.Printf(err.Error())
		return res, nil
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = "Berhasil Prepare Data Table"

	return res, nil
}

func UpdateJarakTableDataFK(idmajikan int, idart int, jarak float64) (Response, error) {
	var res Response

	con := db.CreateCon()

	name := "tabletemp"
	id := strconv.Itoa(idmajikan)
	tablename := name + id

	sqlStatement := "UPDATE " + tablename + " SET jarak=? WHERE iduser=?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(jarak, idart)
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

func SearchDataARTbyFK(idmajikan int, aislam int, akatolik int, akristen int,
	ahindu int, abuddha int, akonghucu int, hewan int, mabukjalan int,
	sepedamotor int, mobil int, masak int, ssingle int, smarried int,
	gajiawal int, gajiakhir int, jarak int, mandatory string) (Response, error) {
	var obj DatabyFilter
	var arrobj []DatabyFilter
	var res Response

	con := db.CreateCon()

	name := "tabletemp"
	id := strconv.Itoa(idmajikan)
	tablename := name + id

	sqlStatement := ""

	sqlStatement = "INSERT INTO " + tablename +
		" (iduser, aislam, akatolik, akristen, ahindu, abuddha, akonghucu," +
		" hewan, mabukjalan, sepedamotor, mobil, masak, ssingle, smarried," +
		" gajiawal, gajiakhir, jarak)" +
		" VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := con.Exec(sqlStatement, idmajikan, aislam, akatolik, akristen, ahindu, abuddha, akonghucu, hewan, mabukjalan, sepedamotor, mobil, masak, ssingle, smarried, gajiawal, gajiakhir, jarak)
	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	if mandatory == "menginap" {
		sqlStatement = "UPDATE " + tablename +
			" SET gajiawal = (gajiawal-0)/(4000000-0)," +
			" gajiakhir = (gajiakhir-0)/(4000000-0)"
		_, err = con.Exec(sqlStatement)
		if err != nil {
			log.Printf(err.Error())
			return res, err
		}

		sqlStatement = "SELECT a.iduser as idart," +
			" ((a.aislam*b.aislam) + (a.akatolik*b.akatolik) + (a.akristen*b.akristen) +" +
			" (a.ahindu*b.ahindu) + (a.abuddha*b.abuddha) + (a.akonghucu*b.akonghucu) +" +
			" (a.hewan*b.hewan) + (a.mabukjalan*b.mabukjalan) + (a.sepedamotor*b.sepedamotor) +" +
			" (a.mobil*b.mobil) + (a.masak*b.masak) + (a.ssingle*b.ssingle) + (a.smarried*b.smarried) +" +
			" (a.gajiawal*b.gajiawal) + (a.gajiakhir*b.gajiakhir)) as innerproduct," +
			" ((a.aislam*a.aislam) + (a.akatolik*a.akatolik) + (a.akristen*a.akristen) +" +
			" (a.ahindu*a.ahindu) + (a.abuddha*a.abuddha) + (a.akonghucu*a.akonghucu) +" +
			" (a.hewan*a.hewan) + (a.mabukjalan*a.mabukjalan) + (a.sepedamotor*a.sepedamotor) +" +
			" (a.mobil*a.mobil) + (a.masak*a.masak) + (a.ssingle*a.ssingle) + (a.smarried*a.smarried) +" +
			" (a.gajiawal*a.gajiawal) + (a.gajiakhir*a.gajiakhir)) as x," +
			" ((b.aislam*b.aislam) + (b.akatolik*b.akatolik) + (b.akristen*b.akristen) +" +
			" (b.ahindu*b.ahindu) + (b.abuddha*b.abuddha) + (b.akonghucu*b.akonghucu) +" +
			" (b.hewan*b.hewan) + (b.mabukjalan*b.mabukjalan) + (b.sepedamotor*b.sepedamotor) +" +
			" (b.mobil*b.mobil) + (b.masak*b.masak) + (b.ssingle*b.ssingle) + (b.smarried*b.smarried) +" +
			" (b.gajiawal*b.gajiawal) + (b.gajiakhir*b.gajiakhir)) as y," +
			" ((a.aislam*b.aislam) + (a.akatolik*b.akatolik) + (a.akristen*b.akristen) +" +
			" (a.ahindu*b.ahindu) + (a.abuddha*b.abuddha) + (a.akonghucu*b.akonghucu) +" +
			" (a.hewan*b.hewan) + (a.mabukjalan*b.mabukjalan) + (a.sepedamotor*b.sepedamotor) +" +
			" (a.mobil*b.mobil) + (a.masak*b.masak) + (a.ssingle*b.ssingle) + (a.smarried*b.smarried) +" +
			" (a.gajiawal*b.gajiawal) + (a.gajiakhir*b.gajiakhir)) /" +
			" sqrt(((a.aislam*a.aislam) + (a.akatolik*a.akatolik) + (a.akristen*a.akristen) +" +
			" (a.ahindu*a.ahindu) + (a.abuddha*a.abuddha) + (a.akonghucu*a.akonghucu) +" +
			" (a.hewan*a.hewan) + (a.mabukjalan*a.mabukjalan) + (a.sepedamotor*a.sepedamotor) +" +
			" (a.mobil*a.mobil) + (a.masak*a.masak) + (a.ssingle*a.ssingle) + (a.smarried*a.smarried) +" +
			" (a.gajiawal*a.gajiawal) + (a.gajiakhir*a.gajiakhir)) *" +
			" ((b.aislam*b.aislam) + (b.akatolik*b.akatolik) + (b.akristen*b.akristen) +" +
			" (b.ahindu*b.ahindu) + (b.abuddha*b.abuddha) + (b.akonghucu*b.akonghucu) +" +
			" (b.hewan*b.hewan) + (b.mabukjalan*b.mabukjalan) + (b.sepedamotor*b.sepedamotor) +" +
			" (b.mobil*b.mobil) + (b.masak*b.masak) + (b.ssingle*b.ssingle) + (b.smarried*b.smarried) +" +
			" (b.gajiawal*b.gajiawal) + (b.gajiakhir*b.gajiakhir))) as cosinesimilarity," +
			" up.namalengkap, jeniskelamin, tempatlahir, tanggallahir, telephone," +
			" dp.pendidikanterakhir, beratbadan, tinggibadan, dp.aislam, dp.akatolik," +
			" dp.akristen, dp.ahindu, dp.abuddha, dp.akonghucu, dp.tkmenginap, dp.tkwarnen," +
			" dp.hewan, dp.mabukjalan, dp.sepedamotor, dp.mobil, dp.masak, dp.ssingle, dp.smarried," +
			" dk.kprt, dk.kbabysitter, dk.kseniorcare, dk.ksupir, dk.kofficeboy, dk.ktukangkebun," +
			" dk.pengalaman, dk.gajiawal, dk.gajiakhir, a.jarak," +
			" AVG(p.rating) as rating" +
			" FROM " + tablename + " a" +
			" JOIN " + tablename + " b" +
			" JOIN userprofile up ON up.iduser = a.iduser" +
			" JOIN detailprofileart dp ON dp.iduser = a.iduser" +
			" JOIN detailkerjaart dk ON dk.iduser = a.iduser" +
			" JOIN penilaian p ON p.idart = a.iduser" +
			" WHERE b.iduser = " + id + " AND a.iduser != " + id +
			" GROUP BY a.iduser" +
			" ORDER BY cosinesimilarity DESC"

		rows, err := con.Query(sqlStatement)
		defer rows.Close()
		if err != nil {
			log.Printf(err.Error())
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.IdART, &obj.InnerProduct, &obj.X, &obj.Y, &obj.CosineSimilarity,
				&obj.NamaLengkap, &obj.JenisKelamin, &obj.TempatLahir, &obj.TanggalLahir, &obj.Telephone,
				&obj.Pendidikan, &obj.BeratBadan, &obj.TinggiBadan, &obj.AIslam, &obj.AKatolik,
				&obj.AKristen, &obj.AHindu, &obj.ABuddha, &obj.AKonghucu, &obj.TKMenginap, &obj.TKWarnen,
				&obj.Hewan, &obj.MabukJalan, &obj.SepedaMotor, &obj.Mobil, &obj.Masak, &obj.SSingle,
				&obj.SMarried, &obj.KPrt, &obj.KBabysitter, &obj.KSeniorcare, &obj.KSupir, &obj.KOfficeboy,
				&obj.KTukangkebun, &obj.Pengalaman, &obj.Gajiawal, &obj.Gajiakhir, &obj.Jarak, &obj.Rating)

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
	} else if mandatory == "warnen" {
		sqlStatement = "UPDATE " + tablename +
			" SET gajiawal = (gajiawal-0)/(4000000-0)," +
			" gajiakhir = (gajiakhir-0)/(4000000-0)," +
			" jarak = (jarak-0)/(10-0)"
		_, err = con.Exec(sqlStatement)
		if err != nil {
			log.Printf(err.Error())
			return res, err
		}

		sqlStatement = "SELECT a.iduser as idart," +
			" ((a.aislam*b.aislam) + (a.akatolik*b.akatolik) + (a.akristen*b.akristen) +" +
			" (a.ahindu*b.ahindu) + (a.abuddha*b.abuddha) + (a.akonghucu*b.akonghucu) +" +
			" (a.hewan*b.hewan) + (a.mabukjalan*b.mabukjalan) + (a.sepedamotor*b.sepedamotor) +" +
			" (a.mobil*b.mobil) + (a.masak*b.masak) + (a.ssingle*b.ssingle) + (a.smarried*b.smarried) +" +
			" (a.gajiawal*b.gajiawal) + (a.gajiakhir*b.gajiakhir) + (a.jarak*b.jarak)) as innerproduct," +
			" ((a.aislam*a.aislam) + (a.akatolik*a.akatolik) + (a.akristen*a.akristen) +" +
			" (a.ahindu*a.ahindu) + (a.abuddha*a.abuddha) + (a.akonghucu*a.akonghucu) +" +
			" (a.hewan*a.hewan) + (a.mabukjalan*a.mabukjalan) + (a.sepedamotor*a.sepedamotor) +" +
			" (a.mobil*a.mobil) + (a.masak*a.masak) + (a.ssingle*a.ssingle) + (a.smarried*a.smarried) +" +
			" (a.gajiawal*a.gajiawal) + (a.gajiakhir*a.gajiakhir) + (a.jarak*a.jarak)) as x," +
			" ((b.aislam*b.aislam) + (b.akatolik*b.akatolik) + (b.akristen*b.akristen) +" +
			" (b.ahindu*b.ahindu) + (b.abuddha*b.abuddha) + (b.akonghucu*b.akonghucu) +" +
			" (b.hewan*b.hewan) + (b.mabukjalan*b.mabukjalan) + (b.sepedamotor*b.sepedamotor) +" +
			" (b.mobil*b.mobil) + (b.masak*b.masak) + (b.ssingle*b.ssingle) + (b.smarried*b.smarried) +" +
			" (b.gajiawal*b.gajiawal) + (b.gajiakhir*b.gajiakhir) + (b.jarak*b.jarak)) as y," +
			" ((a.aislam*b.aislam) + (a.akatolik*b.akatolik) + (a.akristen*b.akristen) +" +
			" (a.ahindu*b.ahindu) + (a.abuddha*b.abuddha) + (a.akonghucu*b.akonghucu) +" +
			" (a.hewan*b.hewan) + (a.mabukjalan*b.mabukjalan) + (a.sepedamotor*b.sepedamotor) +" +
			" (a.mobil*b.mobil) + (a.masak*b.masak) + (a.ssingle*b.ssingle) + (a.smarried*b.smarried) +" +
			" (a.gajiawal*b.gajiawal) + (a.gajiakhir*b.gajiakhir) + (a.jarak*b.jarak)) /" +
			" sqrt(((a.aislam*a.aislam) + (a.akatolik*a.akatolik) + (a.akristen*a.akristen) +" +
			" (a.ahindu*a.ahindu) + (a.abuddha*a.abuddha) + (a.akonghucu*a.akonghucu) +" +
			" (a.hewan*a.hewan) + (a.mabukjalan*a.mabukjalan) + (a.sepedamotor*a.sepedamotor) +" +
			" (a.mobil*a.mobil) + (a.masak*a.masak) + (a.ssingle*a.ssingle) + (a.smarried*a.smarried) +" +
			" (a.gajiawal*a.gajiawal) + (a.gajiakhir*a.gajiakhir) + (a.jarak*a.jarak)) *" +
			" ((b.aislam*b.aislam) + (b.akatolik*b.akatolik) + (b.akristen*b.akristen) +" +
			" (b.ahindu*b.ahindu) + (b.abuddha*b.abuddha) + (b.akonghucu*b.akonghucu) +" +
			" (b.hewan*b.hewan) + (b.mabukjalan*b.mabukjalan) + (b.sepedamotor*b.sepedamotor) +" +
			" (b.mobil*b.mobil) + (b.masak*b.masak) + (b.ssingle*b.ssingle) + (b.smarried*b.smarried) +" +
			" (b.gajiawal*b.gajiawal) + (b.gajiakhir*b.gajiakhir) + (b.jarak*b.jarak))) as cosinesimilarity," +
			" up.namalengkap, jeniskelamin, tempatlahir, tanggallahir, telephone," +
			" dp.pendidikanterakhir, beratbadan, tinggibadan, dp.aislam, dp.akatolik," +
			" dp.akristen, dp.ahindu, dp.abuddha, dp.akonghucu, dp.tkmenginap, dp.tkwarnen," +
			" dp.hewan, dp.mabukjalan, dp.sepedamotor, dp.mobil, dp.masak, dp.ssingle, dp.smarried," +
			" dk.kprt, dk.kbabysitter, dk.kseniorcare, dk.ksupir, dk.kofficeboy, dk.ktukangkebun," +
			" dk.pengalaman, dk.gajiawal, dk.gajiakhir, ((a.jarak * (10-0)) - 0) as jarak," +
			" AVG(p.rating) as rating" +
			" FROM " + tablename + " a" +
			" JOIN " + tablename + " b" +
			" JOIN userprofile up ON up.iduser = a.iduser" +
			" JOIN detailprofileart dp ON dp.iduser = a.iduser" +
			" JOIN detailkerjaart dk ON dk.iduser = a.iduser" +
			" JOIN penilaian p ON p.idart = a.iduser" +
			" WHERE b.iduser = " + id + " AND a.iduser != " + id +
			" GROUP BY a.iduser" +
			" ORDER BY cosinesimilarity DESC"

		rows, err := con.Query(sqlStatement)
		defer rows.Close()
		if err != nil {
			log.Printf(err.Error())
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.IdART, &obj.InnerProduct, &obj.X, &obj.Y, &obj.CosineSimilarity,
				&obj.NamaLengkap, &obj.JenisKelamin, &obj.TempatLahir, &obj.TanggalLahir, &obj.Telephone,
				&obj.Pendidikan, &obj.BeratBadan, &obj.TinggiBadan, &obj.AIslam, &obj.AKatolik,
				&obj.AKristen, &obj.AHindu, &obj.ABuddha, &obj.AKonghucu, &obj.TKMenginap, &obj.TKWarnen,
				&obj.Hewan, &obj.MabukJalan, &obj.SepedaMotor, &obj.Mobil, &obj.Masak, &obj.SSingle,
				&obj.SMarried, &obj.KPrt, &obj.KBabysitter, &obj.KSeniorcare, &obj.KSupir, &obj.KOfficeboy,
				&obj.KTukangkebun, &obj.Pengalaman, &obj.Gajiawal, &obj.Gajiakhir, &obj.Jarak, &obj.Rating)

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
	ktukangkebun int, pengalaman string, gajiawal int, gajiakhir int) (Response, error) {
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

type DataKontakARTObj struct {
	Idart        int     `json:"idart"`
	Namalengkap  string  `json:"namalengkap"`
	Tanggallahir string  `json:"tanggallahir"`
	Telephone    string  `json:"telephone"`
	Kprt         int     `json:"kprt"`
	Kbabysitter  int     `json:"kbabysitter"`
	Kseniorcare  int     `json:"kseniorcare"`
	Ksupir       int     `json:"ksupir"`
	Kofficeboy   int     `json:"kofficeboy"`
	Ktukangkebun int     `json:"ktukangkebun"`
	Rating       float64 `json:"rating"`
}

func DataKontakART(idmajikan int) (Response, error) {
	var obj DataKontakARTObj
	var arrobj []DataKontakARTObj
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT ka.idart as idart, up.namalengkap, up.tanggallahir, up.telephone, dk.kprt," +
		" dk.kbabysitter, dk.kseniorcare, dk.ksupir, dk.kofficeboy, dk.ktukangkebun, AVG(p.rating)" +
		" FROM kontakart ka" +
		" JOIN userprofile up ON ka.idart = up.iduser" +
		" JOIN detailkerjaart dk on ka.idart = dk.iduser" +
		" JOIN penilaian p ON ka.idart = p.idart" +
		" WHERE ka.idmajikan = ?" +
		" GROUP BY ka.idart"

	rows, err := con.Query(sqlStatement, idmajikan)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Idart, &obj.Namalengkap, &obj.Tanggallahir, &obj.Telephone,
			&obj.Kprt, &obj.Kbabysitter, &obj.Kseniorcare, &obj.Ksupir,
			&obj.Kofficeboy, &obj.Ktukangkebun, &obj.Rating)

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

type DataPelamarObject struct {
	Idloker      int     `json:"idloker"`
	Judulloker   string  `json:"judulloker"`
	Idart        int     `json:"idart"`
	Namalengkap  string  `json:"namalengkap"`
	Tanggallahir string  `json:"tanggallahir"`
	Telephone    string  `json:"telephone"`
	Kprt         int     `json:"kprt"`
	Kbabysitter  int     `json:"kbabysitter"`
	Kseniorcare  int     `json:"kseniorcare"`
	Ksupir       int     `json:"ksupir"`
	Kofficeboy   int     `json:"kofficeboy"`
	Ktukangkebun int     `json:"ktukangkebun"`
	Rating       float64 `json:"rating"`
}

func DataPelamar(idmajikan int) (Response, error) {
	var obj DataPelamarObject
	var arrobj []DataPelamarObject
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT ll.idloker as idloker, lk.judulloker," +
		" ll.idart as idart, up.namalengkap, up.tanggallahir, up.telephone," +
		" dk.kprt, dk.kbabysitter, dk.kseniorcare, dk.ksupir, dk.kofficeboy," +
		" dk.ktukangkebun, AVG(p.rating)" +
		" FROM lamarloker ll" +
		" JOIN userprofile up ON ll.idart = up.iduser" +
		" JOIN detailkerjaart dk ON ll.idart = dk.iduser" +
		" JOIN penilaian p ON ll.idart = p.idart" +
		" JOIN lowongankerja lk ON ll.idloker = lk.idloker" +
		" WHERE lk.iduser = ?" +
		" GROUP BY ll.idlamar"

	rows, err := con.Query(sqlStatement, idmajikan)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Idloker, &obj.Judulloker, &obj.Idart, &obj.Namalengkap, &obj.Tanggallahir, &obj.Telephone,
			&obj.Kprt, &obj.Kbabysitter, &obj.Kseniorcare, &obj.Ksupir,
			&obj.Kofficeboy, &obj.Ktukangkebun, &obj.Rating)

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

func SaveLamaran(idloker int, idart int, waktulamar string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO lamarloker (idloker, idart, waktulamar) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(idloker, idart, waktulamar)
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

type DataLamaranObj struct {
	Idlamar     int    `json:"idlamar"`
	Idloker     int    `json:"idloker"`
	Judulloker  string `json:"judulloker"`
	Iduser      int    `json:"iduser"`
	Namalengkap string `json:"namalengkap"`
	Kecamatan   string `json:"kecamatan"`
	Kota        string `json:"kota"`
	Waktulamar  string `json:"waktulamar"`
}

func DataLamaran(idart int) (Response, error) {
	var obj DataLamaranObj
	var arrobj []DataLamaranObj
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT ll.idlamar, ll.idloker, lk.judulloker, lk.iduser, up.namalengkap," +
		" ud.kecamatan, ud.kota, ll.waktulamar" +
		" FROM lamarloker ll" +
		" JOIN lowongankerja lk ON ll.idloker = lk.idloker" +
		" JOIN userprofile up ON lk.iduser = up.iduser" +
		" JOIN userdomisili ud ON lk.iduser = ud.iduser" +
		" WHERE ll.idart = ?" +
		" ORDER BY ll.waktulamar DESC"

	rows, err := con.Query(sqlStatement, idart)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Idlamar, &obj.Idloker, &obj.Judulloker, &obj.Iduser,
			&obj.Namalengkap, &obj.Kecamatan, &obj.Kota, &obj.Waktulamar)

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

type DataKontakMajikanObj struct {
	Idmajikan   int    `json:"idmajikan"`
	Namalengkap string `json:"namalengkap"`
	Telephone   string `json:"telephone"`
	Kecamatan   string `json:"kecamatan"`
	Kota        string `json:"kota"`
}

func DataKontakMajikan(idart int) (Response, error) {
	var obj DataKontakMajikanObj
	var arrobj []DataKontakMajikanObj
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT ka.idmajikan as idmajikan, up.namalengkap, up.telephone," +
		" ud.kecamatan, ud.kota " +
		"FROM kontakart ka " +
		"JOIN userprofile up ON ka.idmajikan = up.iduser " +
		"JOIN userdomisili ud on ka.idmajikan = ud.iduser " +
		"WHERE ka.idart = ?"

	rows, err := con.Query(sqlStatement, idart)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Idmajikan, &obj.Namalengkap, &obj.Telephone, &obj.Kecamatan, &obj.Kota)

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

type KontakArtObj struct {
	Idkontak    int    `json:"idkontak"`
	Idmajikan   int    `json:"idmajikan"`
	Idart       int    `json:"idart"`
	Waktukontak string `json:"waktukontak"`
}

func KontakArt(idart int, idmajikan int) (Response, error) {
	var obj KontakArtObj
	var arrobj []KontakArtObj
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM kontakart WHERE idart = ? AND idmajikan = ?"

	rows, err := con.Query(sqlStatement, idart, idmajikan)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Idkontak, &obj.Idmajikan, &obj.Idart, &obj.Waktukontak)

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

func SimpanLamaran(idloker int, idart int, waktulamar string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO kontakart (idloker, idart, waktulamar) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(idloker, idart, waktulamar)
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

type LamarObj struct {
	Idlamar    int    `json:"idlamar"`
	Idloker    int    `json:"idloker"`
	Idart      int    `json:"idart"`
	Waktulamar string `json:"waktulamar"`
}

func LamarLoker(idloker int, idart int) (Response, error) {
	var obj LamarObj
	var arrobj []LamarObj
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM lamarloker WHERE idloker = ? AND idart = ?"

	rows, err := con.Query(sqlStatement, idloker, idart)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Idlamar, &obj.Idloker, &obj.Idart, &obj.Waktulamar)

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

type IdLamarObj struct {
	Idlamar int `json:"idlamar"`
	Idloker int `json:"idloker"`
}

func GetIdLamar(idloker int) (Response, error) {
	var obj IdLamarObj
	var arrobj []IdLamarObj
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT idlamar, idloker FROM lamarloker WHERE idloker = ?"

	rows, err := con.Query(sqlStatement, idloker)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Idlamar, &obj.Idloker)

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

func DeleteLamarLoker(idlamar int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM lamarloker WHERE idlamar = ?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(idlamar)
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

type RataPenilaian struct {
	Estetika   float64 `json:"estetika"`
	Etika      float64 `json:"etika"`
	Kebersihan float64 `json:"kebersihan"`
	Kerapian   float64 `json:"kerapian"`
	Kecepatan  float64 `json:"kecepatan"`
}

func SimpanPenilaian(idart int, idmajikan int, etika int, estetika int,
	kebersihan int, kerapian int, kecepatan int, rating float64, review string, tglpost string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO penilaian (idart, idmajikan, estetika, etika, kebersihan, kecepatan, kerapian, rating, review, tglpost) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(idart, idmajikan, estetika, etika, kebersihan, kecepatan, kerapian, rating, review, tglpost)
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

func RataPenilaianART(idart int) (Response, error) {
	var obj RataPenilaian
	var arrobj []RataPenilaian
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT AVG(estetika) as estetika, AVG(etika) as etika ,AVG(kebersihan) as kebersihan, AVG(kecepatan) as kecepatan, AVG(kerapian) as kerapian FROM penilaian WHERE idart = ?"

	rows, err := con.Query(sqlStatemet, idart)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Estetika, &obj.Etika,
			&obj.Kebersihan, &obj.Kerapian, &obj.Kecepatan)

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

type ReviewMajikan struct {
	IdMajikan   int    `json:"idmajikan"`
	NamaLengkap string `json:"namalengkap"`
	Review      string `json:"review"`
	TglPost     string `json:"tglpost"`
}

func DataReviewMajikan(idart int) (Response, error) {
	var obj ReviewMajikan
	var arrobj []ReviewMajikan
	var res Response

	con := db.CreateCon()

	sqlStatemet := "SELECT p.idmajikan, up.namalengkap, p.review, p.tglpost FROM penilaian p JOIN userprofile up ON p.idmajikan = up.iduser WHERE idart = ?"

	rows, err := con.Query(sqlStatemet, idart)

	defer rows.Close()

	if err != nil {
		log.Printf(err.Error())
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdMajikan, &obj.NamaLengkap, &obj.Review, &obj.TglPost)

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
