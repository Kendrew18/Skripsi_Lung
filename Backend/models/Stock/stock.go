package stock

import (
	"kendrew/db"
	"kendrew/list"
	"kendrew/tools"
	"net/http"
)

func ShowStock() (tools.Response, error) {
	var stock list.ShowAllProduct
	var arr_stock []list.ShowAllProduct
	var res tools.Response

	con := db.CreateCon()

	sqlStatement := "SELECT stock.id_stock, stock.nama_barang, stock.harga, stock.jenis_barang FROM stock"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		var ukuran list.Ukuran
		var arr_ukuran []list.Ukuran

		err = rows.Scan(&stock.IdProduct, &stock.NamaProduct,
			&stock.Harga, &stock.JenisProduct)

		if err != nil {
			return res, err
		}

		sqlStatement2 := "SELECT jenis_ukuran,stock FROM stock_to_ukuran JOIN ukuran ON stock_to_ukuran.id_ukuran=ukuran.id_ukuran WHERE id_stock=?"

		rows2, err := con.Query(sqlStatement2)

		defer rows2.Close()

		if err != nil {
			return res, err
		}

		for rows2.Next() {

			err = rows.Scan(&ukuran.UkuranProduct, &ukuran.Stock)
			if err != nil {
				return res, err
			}

			arr_ukuran = append(arr_ukuran, ukuran)
		}

		stock.Ukuran = arr_ukuran

		arr_stock = append(arr_stock, stock)
	}

	if arr_stock == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_stock
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_stock
	}

	return res, nil
}

func SaveStock(nama_barang string, harga_barang int64, jenis_barang string,
	id_ukuran string, jumlah_stock string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlstatement := "INSERT INTO stock (nama_barang,harga,jenis_barang) values(?,?,?)"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama_barang, harga_barang, jenis_barang)

	if err != nil {
		return res, err
	}

	getIdLast, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	id_u := tools.String_Separator_To_Int(id_ukuran)
	jumlah_stok_int := tools.String_Separator_To_Int(jumlah_stock)

	id_stock := int(getIdLast)

	for i := 0; i < len(id_u); i++ {

		sqlstatement = "INSERT INTO stock_to_ukuran (id_stock,id_ukuran,stock) values(?,?,?)"

		stmt, err = con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(id_stock, id_u[i], jumlah_stok_int[i])

		if err != nil {
			return res, err
		}
	}

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = map[string]int64{
		"getIdLast": getIdLast,
	}

	return res, nil
}
