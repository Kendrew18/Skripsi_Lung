package order

import (
	"kendrew/db"
	"kendrew/list"
	"kendrew/tools"
	"net/http"
	"time"
)

func SaveOrder(id_pelanggan string, tanggal_pemesanan string, no_order string,
	pembayaran string, down_payment string, tanggal_pembayaran string,
	catatan string, total_barang int, harga_jual int64, sub_total int64,
	id_ukuran string, id_stock string, jumlah string, satuan string) (tools.Response, error) {

	var res tools.Response

	con := db.CreateCon()

	sqlstatement := "INSERT INTO order (id_pelanggan, tanggal_pemesanan, no_order, pembayaran, down_payment, tanggal_pembayaran, total_barang, catatan, harga_jual, sub_total) values(?,?,?,?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	date_pemesanan, _ := time.Parse("02-01-2006", tanggal_pemesanan)
	date_pemesanan_sql := date_pemesanan.Format("2006-01-02")

	date_pembayaran, _ := time.Parse("02-01-2006", tanggal_pembayaran)
	date_pembayaran_sql := date_pembayaran.Format("2006-01-02")

	result, err := stmt.Exec(id_pelanggan, date_pemesanan_sql, no_order,
		pembayaran, down_payment, date_pembayaran_sql, total_barang,
		catatan, harga_jual, sub_total)

	if err != nil {
		return res, err
	}

	getIdLast, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	id_s := tools.String_Separator_To_String(id_stock)
	id_u := tools.String_Separator_To_Int(id_ukuran)
	jumlah_order := tools.String_Separator_To_Int(jumlah)
	satuan_order := tools.String_Separator_To_Int(satuan)

	id_order := int(getIdLast)

	for i := 0; i < len(id_u); i++ {

		sqlstatement = "INSERT INTO detail_order (id_ukuran, jumlah, id_order, satuan, id_stock) values(?,?,?,?,?)"

		stmt, err = con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(id_u[i], jumlah_order[i], id_order, satuan_order[i], id_s[i])

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

func ShowOrderHeader(id_sales int) (tools.Response, error) {
	var orderheader list.ShowOrderHeader
	var arr_orderheader []list.ShowOrderHeader
	var res tools.Response

	con := db.CreateCon()

	sqlStatement := "SELECT id_order, no_order, sub_total, name, nama_toko, kota, total_barang FROM `order` JOIN user ON `order`.id_sales = user.id_user JOIN pelanggan ON pelanggan.id_pelanggan = `order`.id_pelanggan WHERE id_sales=?"

	rows, err := con.Query(sqlStatement, id_sales)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&orderheader.Id_order, &orderheader.No_order,
			&orderheader.Sub_total, &orderheader.Nama_sales,
			&orderheader.Nama_toko, &orderheader.Kota, &orderheader.Total_barang)

		if err != nil {
			return res, err
		}

		arr_orderheader = append(arr_orderheader, orderheader)
	}

	if arr_orderheader == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_orderheader
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_orderheader
	}

	return res, nil
}
