package salesman

import (
	"kendrew/db"
	"kendrew/list"
	"kendrew/tools"
	"net/http"
)

// Show All Salesman
func ShowSales() (tools.Response, error) {
	var salesman list.SalesNames
	var arr_salesman []list.SalesNames
	var res tools.Response

	con := db.CreateCon()

	sqlStatement := "SELECT id_user, name, alamat, nomor_hp, bank, no_rekening FROM user "

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&salesman.IdSales, &salesman.NamaSales, &salesman.Alamat,
			&salesman.Nomor_hp, &salesman.Bank, &salesman.No_rekening)
		if err != nil {
			return res, err
		}
		arr_salesman = append(arr_salesman, salesman)
	}

	if arr_salesman == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_salesman
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_salesman
	}

	return res, nil
}

func SaveSales(name string, alamat string, nomor_hp string, bank string,
	no_rekening string, username string, password string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlstatement := "INSERT INTO user (name, alamat, nomor_hp, bank, no_rekening, username, password) values(?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, alamat, nomor_hp, bank, no_rekening, username, password)

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
