package retur

import (
	"kendrew/db"
	"kendrew/list"
	"kendrew/tools"
	"net/http"
	"time"
)

func SaveRetur(no_order string, tanggal_retur string, total_barang int,
	harga_jual int64, sub_total int64, id_stock string, id_ukuran string,
	jumlah string, satuan string) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlstatement := "INSERT INTO retur (no_order, tanggal_retur, total_barang, harga_jual, sub_total) values(?,?,?,?,?)"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	date_retur, _ := time.Parse("02-01-2006", tanggal_retur)
	date_retur_sql := date_retur.Format("2006-01-02")

	result, err := stmt.Exec(no_order, date_retur_sql, total_barang, harga_jual, sub_total)

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

	id_retur := int(getIdLast)

	for i := 0; i < len(id_u); i++ {

		sqlstatement = "INSERT INTO detail_retur (id_ukuran, jumlah, id_retur, satuan, id_stock) values(?,?,?,?,?)"

		stmt, err = con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(id_u[i], jumlah_order[i], id_retur, satuan_order[i], id_s[i])

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

func ShowAllTransaksi() (tools.Response, error) {

	var res tools.Response
	var obj list.ShowTransaksi
	var arrobj []list.ShowTransaksi

	con := db.CreateCon()

	sqlStatement := "SELECT transaksi.id_transaksi, pelanggan.nama, pelanggan.alamat, pelanggan.no_hp, user.name, transaksi.tanggal, transaksi.pembayaran, transaksi.total from transaksi join pelanggan on transaksi.id_pelanggan = pelanggan.id_pelanggan join user on transaksi.id_user = user.id_user;"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdTransaksi, &obj.NamaPelanggan, &obj.AlamatPelanggan, &obj.NoHpPelanggan, &obj.NamaSales, &obj.Tanggal, &obj.Pembayaran, &obj.TotalKeseluruhan)
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

func ShowAllTransaksiBySalesId(idSales int) (tools.Response, error) {

	var res tools.Response
	var obj list.ShowTransaksi
	var arrobj []list.ShowTransaksi

	con := db.CreateCon()

	sqlStatement := "SELECT transaksi.id_transaksi, pelanggan.nama, pelanggan.alamat, pelanggan.no_hp, user.name, transaksi.tanggal, transaksi.pembayaran, transaksi.total from transaksi join pelanggan on transaksi.id_pelanggan = pelanggan.id_pelanggan join user on transaksi.id_user = user.id_user WHERE user.id_user = ?;"

	rows, err := con.Query(sqlStatement, idSales)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdTransaksi, &obj.NamaPelanggan, &obj.AlamatPelanggan, &obj.NoHpPelanggan, &obj.NamaSales, &obj.Tanggal, &obj.Pembayaran, &obj.TotalKeseluruhan)
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

func ShowDetailTransaksi(IdTransaksi int) (tools.Response, error) {

	var res tools.Response
	var obj list.ShowDetailTransaksi
	var arrobj []list.ShowDetailTransaksi

	con := db.CreateCon()

	sqlStatement := "SELECT detail_transaksi.id_detail_transaksi, detail_transaksi.id_transaksi, detail_transaksi.id_product, product.nama_product, product.jenis_product, detail_transaksi.jumlah, detail_transaksi.total from detail_transaksi join transaksi on detail_transaksi.id_transaksi = transaksi.id_transaksi join product on detail_transaksi.id_product = product.id_product where detail_transaksi.id_transaksi = ?;"

	rows, err := con.Query(sqlStatement, IdTransaksi)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.IdDetailTransaksi, &obj.IdTransaksi, &obj.IdProduct, &obj.NamaProduct, &obj.JenisProduct, &obj.Jumlah, &obj.Total)
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

func SimpanPesananDetail(idTransaksi int, idProduct int, jumlah int, total int) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	sqlstatement := "INSERT INTO detail_transaksi (idTransaksi, idProduct, jumlah, total) values(?,?,?,?)"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(idTransaksi, idProduct, jumlah, total)

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
