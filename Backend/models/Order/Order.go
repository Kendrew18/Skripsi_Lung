package order

import (
	"kendrew/db"
	odr "kendrew/list/Order"
	"kendrew/tools"
	"net/http"
	"time"
)

// save order
func SaveOrder(id_pelanggan int, id_sales int, tanggal_pemesanan string, no_order string,
	pembayaran string, down_payment string, tanggal_pembayaran string,
	catatan string, total_barang string, harga_jual string, sub_total string,
	id_ukuran string, id_stock string, jumlah string, satuan string) (tools.Response, error) {

	var res tools.Response

	con := db.CreateCon()

	//input order

	sqlstatement := "INSERT INTO `order` (id_pelanggan, tanggal_pesanan, no_order, pembayaran, down_payment, tanggal_pembayaran, catatan, id_sales) values(?,?,?,?,?,?,?,?)"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	date_pemesanan, _ := time.Parse("02-01-2006", tanggal_pemesanan)
	date_pemesanan_sql := date_pemesanan.Format("2006-01-02")

	date_pembayaran, _ := time.Parse("02-01-2006", tanggal_pembayaran)
	date_pembayaran_sql := date_pembayaran.Format("2006-01-02")

	result, err := stmt.Exec(id_pelanggan, date_pemesanan_sql, no_order,
		pembayaran, down_payment, date_pembayaran_sql,
		catatan, id_sales)

	if err != nil {
		return res, err
	}

	getIdLast, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	id_order := 0

	sqlStatement := "SELECT id_order FROM `order` WHERE `order`.id_sales=? ORDER BY id_order DESC LIMIT 1"

	err = con.QueryRow(sqlStatement, id_sales).Scan(&id_order)

	if err != nil {
		return res, err
	}

	//input detail order
	id_s := tools.String_Separator_To_String(id_stock)
	total_b := tools.String_Separator_To_float64(total_barang)
	harga_j := tools.String_Separator_To_Int64(harga_jual)
	sub_t := tools.String_Separator_To_Int64(sub_total)

	id_u := tools.String_Separator_To_String(id_ukuran)
	jumlah_order := tools.String_Separator_To_String(jumlah)
	satuan_order := tools.String_Separator_To_String(satuan)

	for i := 0; i < len(id_s); i++ {

		sqlstatement = "INSERT INTO detail_order (id_order, id_stock, jumlah, satuan, harga_jual, sub_total) values(?,?,?,?,?,?)"

		stmt, err = con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(id_order, id_s[i], total_b[i], "pcs", harga_j[i], sub_t[i])

		if err != nil {
			return res, err
		}

		var id odr.ShowID_DetailOrderBarang

		sqlStatement = "SELECT id_detail_order FROM detail_order WHERE id_order=? && id_stock=?"

		err := con.QueryRow(sqlStatement, id_order, id_s[i]).Scan(&id.Id_detail_order)

		if err != nil {
			return res, err
		}

		//input_detail_order_barang
		id_u_2_ss := tools.String_Separator_To_Int_Tanda_Seru(id_u[i])
		jumlah_order_2_ss := tools.String_Separator_To_Int_Tanda_Seru(jumlah_order[i])
		satuan_order_2_ss := tools.String_Separator_To_String_Tanda_Seru(satuan_order[i])

		for j := 0; j < len(id_u_2_ss); j++ {

			sqlstatement = "INSERT INTO detail_order_barang (id_detail_order, id_ukuran, jumlah, satuan) values(?,?,?,?)"

			stmt, err = con.Prepare(sqlstatement)

			if err != nil {
				return res, err
			}

			_, err = stmt.Exec(id.Id_detail_order, id_u_2_ss[j], jumlah_order_2_ss[j], satuan_order_2_ss[j])

			if err != nil {
				return res, err
			}
		}

	}

	res.Status = http.StatusOK
	res.Message = "Suksess"
	res.Data = map[string]int64{
		"getIdLast": getIdLast,
	}

	return res, nil
}

// show order header
func ShowOrderHeader(id_sales int) (tools.Response, error) {
	var orderheader odr.ShowOrderHeader
	var arr_orderheader []odr.ShowOrderHeader
	var res tools.Response

	con := db.CreateCon()

	sqlStatement := "SELECT `order`.id_order, no_order, SUM(detail_order.sub_total), name, nama_toko, SUM(detail_order.jumlah), kota FROM `order` JOIN user ON `order`.id_sales = user.id_user JOIN pelanggan ON pelanggan.id_pelanggan = `order`.id_pelanggan JOIN detail_order ON detail_order.id_order = `order`.id_order WHERE id_sales=? GROUP BY id_order"

	rows, err := con.Query(sqlStatement, id_sales)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&orderheader.Id_order, &orderheader.No_order,
			&orderheader.Sub_total, &orderheader.Nama_sales,
			&orderheader.Nama_toko, &orderheader.Total_barang, &orderheader.Kota)

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

// Show Detail Transaksi
func ShowDetailTransaksi(id_order int) (tools.Response, error) {
	var DetailTransaksi odr.ShowDetailTransaksi
	var arr_DetailTransaksi []odr.ShowDetailTransaksi
	var res tools.Response

	con := db.CreateCon()

	sqlStatement := "SELECT `order`.id_order, nama_toko, pelanggan.alamat, kota, pelanggan.no_telp, SUM(detail_order.jumlah), `order`.no_order, tanggal_pesanan, tanggal_pengiriman, pembayaran, down_payment, tanggal_pembayaran, name, SUM(detail_order.sub_total) FROM `order` JOIN `user` ON `order`.id_sales = `user`.id_user JOIN pelanggan ON pelanggan.id_pelanggan = `order`.id_pelanggan JOIN detail_order ON detail_order.id_order=`order`.`id_order` WHERE `order`.id_order=?"

	err := con.QueryRow(sqlStatement, id_order).Scan(&DetailTransaksi.Id_order, &DetailTransaksi.Nama_toko, &DetailTransaksi.Alamat,
		&DetailTransaksi.Kota, &DetailTransaksi.No_telp, &DetailTransaksi.Total_barang, &DetailTransaksi.No_order,
		&DetailTransaksi.Tanggal_pesanan, &DetailTransaksi.Tanggal_pengiriman, &DetailTransaksi.Pembayaran,
		&DetailTransaksi.Down_payment, &DetailTransaksi.Tanggal_pembayaran, &DetailTransaksi.Nama_sales,
		&DetailTransaksi.Sub_total)

	if err != nil {
		return res, err
	}

	if DetailTransaksi.Tanggal_pengiriman == "0000-00-00" {
		DetailTransaksi.Tanggal_pengiriman = "-"
	}

	var id_retur = 0

	sqlStatement = "SELECT retur.id_retur FROM retur WHERE no_order=?"

	err = con.QueryRow(sqlStatement, DetailTransaksi.No_order).Scan(&id_retur)

	sqlStatement = "SELECT SUM(sub_total) FROM detail_retur WHERE id_retur=?"

	err = con.QueryRow(sqlStatement, id_retur).Scan(&DetailTransaksi.Retur)

	var id odr.Detail_Order
	var arr_id []odr.Detail_Order

	sqlStatement = "SELECT id_detail_order, nama_barang, jumlah, satuan, harga_jual, sub_total FROM detail_order JOIN stock s ON s.id = detail_order.id_stock WHERE detail_order.id_order=?"

	rows, err := con.Query(sqlStatement, id_order)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&id.Id_detail_order, &id.Nama_barang, &id.Jumlah,
			&id.Satuan, &id.Harga_jual, &id.Sub_total)

		if err != nil {
			return res, err
		}

		arr_id = append(arr_id, id)

	}

	for i := 0; i < len(arr_id); i++ {

		var D_O_B odr.Detail_Order_Barang
		var arr_D_O_B []odr.Detail_Order_Barang

		sqlStatement = "SELECT jenis_ukuran, jumlah, satuan FROM detail_order_barang JOIN ukuran u ON u.id_ukuran = detail_order_barang.id_ukuran WHERE id_detail_order=?"

		rows, err = con.Query(sqlStatement, arr_id[i].Id_detail_order)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {

			err = rows.Scan(&D_O_B.Jenis_ukuran, &D_O_B.Jumlah, &D_O_B.Satuan)

			if err != nil {
				return res, err
			}

			arr_D_O_B = append(arr_D_O_B, D_O_B)

		}

		arr_id[i].Detail_order_barang = arr_D_O_B

	}

	DetailTransaksi.Detail_order = arr_id

	arr_DetailTransaksi = append(arr_DetailTransaksi, DetailTransaksi)

	if arr_DetailTransaksi == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_DetailTransaksi
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_DetailTransaksi
	}

	return res, nil
}

func ShowOrderHeader_Admin() (tools.Response, error) {
	var orderheader odr.ShowOrderHeader
	var arr_orderheader []odr.ShowOrderHeader
	var res tools.Response

	con := db.CreateCon()

	sqlStatement := "SELECT `order`.id_order, no_order, SUM(detail_order.sub_total), name, nama_toko, SUM(detail_order.jumlah), kota FROM `order` JOIN user ON `order`.id_sales = user.id_user JOIN pelanggan ON pelanggan.id_pelanggan = `order`.id_pelanggan JOIN detail_order ON detail_order.id_order = `order`.id_order GROUP BY id_order"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&orderheader.Id_order, &orderheader.No_order,
			&orderheader.Sub_total, &orderheader.Nama_sales,
			&orderheader.Nama_toko, &orderheader.Total_barang, &orderheader.Kota)

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
