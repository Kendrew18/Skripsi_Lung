package models

import (
	"fmt"
	"kendrew/db"
	"kendrew/list"
	"kendrew/tools"
	"net/http"
)

// Login
func Login(username string, password string) (tools.Response, error) {
	var arr list.LoginUser
	var res tools.Response

	con := db.CreateCon()

	sqlStatement := "SELECT id_user FROM user where username=? && password=?"

	err := con.QueryRow(sqlStatement, username, password).Scan(&arr.IdUser)

	if err != nil || arr.IdUser == "" {
		arr.IdUser = ""
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr

	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr
	}

	fmt.Println(arr)
	return res, nil
}

func Login_admin(username string, password string) (tools.Response, error) {
	var arr list.LoginUser
	var res tools.Response

	con := db.CreateCon()

	sqlStatement := "SELECT id_admin FROM admin where username=? && password=?"

	err := con.QueryRow(sqlStatement, username, password).Scan(&arr.IdUser)

	if err != nil || arr.IdUser == "" {
		arr.IdUser = ""
		res.Status = http.StatusNotFound
		res.Message = "Status Not Found"
		res.Data = arr

	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr
	}

	fmt.Println(arr)
	return res, nil
}
