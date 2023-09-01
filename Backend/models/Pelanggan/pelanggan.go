package pelanggan

import (
	"kendrew/db"
	"kendrew/list"
	"kendrew/tools"
	"net/http"
)

func ShowPelanggan() (tools.Response, error) {
	var pelanggan list.ShowPelanggan
	var arr_pelanggan []list.ShowPelanggan
	var res tools.Response

	con := db.CreateCon()

	sqlStatement := "SELECT id_pelanggan, nama_toko, no_telp, alamat, kota, provinsi, nama_penanggungjawab FROM pelanggan"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&pelanggan.Id_pelanggan, &pelanggan.Nama_toko,
			&pelanggan.No_telp, &pelanggan.Alamat, &pelanggan.Kota,
			&pelanggan.Provinsi, &pelanggan.Nama_penanggungjawab)

		if err != nil {
			return res, err
		}

		arr_pelanggan = append(arr_pelanggan, pelanggan)
	}

	if arr_pelanggan == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_pelanggan
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_pelanggan
	}

	return res, nil
}

func SavePelanggan(nama_toko string, no_hp string, alamat string, kota string,
	provinsi string, nama_penanggungjawab string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlstatement := "INSERT INTO pelanggan (nama_toko, no_telp, alamat, kota, provinsi, nama_penanggungjawab) values(?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama_toko, no_hp, alamat,
		kota, provinsi, nama_penanggungjawab)

	if err != nil {
		return res, err
	}

	getIdLast, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = map[string]int64{
		"getIdLast": getIdLast,
	}

	return res, nil
}
