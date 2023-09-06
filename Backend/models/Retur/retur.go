package retur

import (
	"kendrew/db"
	retur "kendrew/list/Retur"
	"kendrew/tools"
	"net/http"
	"time"
)

// Save_Retur
func SaveRetur(no_order string, tanggal_retur string, total_barang string,
	harga_jual string, sub_total string, id_stock string, id_ukuran string,
	jumlah string, satuan string, id_sales int) (tools.Response, error) {
	var res tools.Response

	con := db.CreateCon()

	id_pelanggan := 0

	sqlstatement := "SELECT id_pelanggan FROM `order` WHERE no_order=?"

	err := con.QueryRow(sqlstatement, no_order).Scan(&id_pelanggan)

	if err != nil {
		return res, err
	}

	sqlstatement = "INSERT INTO retur (no_order, tanggal_retur,id_sales,id_pelanggan) values(?,?,?,?)"

	stmt, err := con.Prepare(sqlstatement)

	if err != nil {
		return res, err
	}

	date_retur, _ := time.Parse("02-01-2006", tanggal_retur)
	date_retur_sql := date_retur.Format("2006-01-02")

	result, err := stmt.Exec(no_order, date_retur_sql, id_sales, id_pelanggan)

	if err != nil {
		return res, err
	}

	getIdLast, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	id_retur := 0

	sqlStatement := "SELECT id_retur FROM retur WHERE retur.id_sales=? ORDER BY id_retur DESC LIMIT 1"

	err = con.QueryRow(sqlStatement, id_sales).Scan(&id_retur)

	if err != nil {
		return res, err
	}

	//input detail retur
	id_s := tools.String_Separator_To_String(id_stock)
	total_b := tools.String_Separator_To_float64(total_barang)
	harga_j := tools.String_Separator_To_Int64(harga_jual)
	sub_t := tools.String_Separator_To_Int64(sub_total)

	id_u := tools.String_Separator_To_String(id_ukuran)
	jumlah_retur := tools.String_Separator_To_String(jumlah)
	satuan_retur := tools.String_Separator_To_String(satuan)

	for i := 0; i < len(id_s); i++ {

		sqlstatement = "INSERT INTO detail_retur (id_retur, id_stock, jumlah, satuan, harga_jual, sub_total) values(?,?,?,?,?,?)"

		stmt, err = con.Prepare(sqlstatement)

		if err != nil {
			return res, err
		}

		_, err = stmt.Exec(id_retur, id_s[i], total_b[i], "pcs", harga_j[i], sub_t[i])

		if err != nil {
			return res, err
		}

		var id retur.ShowID_DetailRetur

		sqlStatement = "SELECT id_detail_retur FROM detail_retur WHERE id_retur=? && id_stock=?"

		err := con.QueryRow(sqlStatement, id_retur, id_s[i]).Scan(&id.Id_detail_retur)

		if err != nil {
			return res, err
		}

		// input_detail_order_barang
		id_u_2_ss := tools.String_Separator_To_Int_Tanda_Seru(id_u[i])
		jumlah_order_2_ss := tools.String_Separator_To_Int_Tanda_Seru(jumlah_retur[i])
		satuan_order_2_ss := tools.String_Separator_To_String_Tanda_Seru(satuan_retur[i])

		for j := 0; j < len(id_u_2_ss); j++ {

			sqlstatement = "INSERT INTO detail_retur_barang (id_detail_retur, id_ukuran, jumlah, satuan) values(?,?,?,?)"

			stmt, err = con.Prepare(sqlstatement)

			if err != nil {
				return res, err
			}

			_, err = stmt.Exec(id.Id_detail_retur, id_u_2_ss[j], jumlah_order_2_ss[j], satuan_order_2_ss[j])

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

func ShowReturHeader(id_sales int) (tools.Response, error) {
	var res tools.Response
	var obj retur.ShowReturHeader
	var arrobj []retur.ShowReturHeader

	con := db.CreateCon()

	sqlStatement := "SELECT retur.id_retur, tanggal_retur, SUM(detail_retur.jumlah), nama_toko, kota FROM retur JOIN user ON retur.id_sales = user.id_user JOIN pelanggan ON pelanggan.id_pelanggan = retur.id_pelanggan JOIN detail_retur ON detail_retur.id_retur = retur.id_retur WHERE id_sales=? GROUP BY id_retur"

	rows, err := con.Query(sqlStatement, id_sales)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		var DRB retur.Detail_Retur_Header
		var arr_DRB []retur.Detail_Retur_Header

		err = rows.Scan(&obj.Id_retur, &obj.Tanggal_retur,
			&obj.Total_barang, &obj.Nama_toko, &obj.Kota)

		if err != nil {
			return res, err
		}

		sqlStatement = "SELECT id_detail_retur, nama_barang, jumlah, satuan FROM detail_retur JOIN stock s ON detail_retur.id_stock = s.id WHERE id_retur=?"

		rows2, err := con.Query(sqlStatement, obj.Id_retur)
		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows2.Next() {
			err = rows2.Scan(&DRB.Id_detail_retur,
				&DRB.Nama_barang, &DRB.Jumlah, &DRB.Satuan)

			if err != nil {
				return res, err
			}

			arr_DRB = append(arr_DRB, DRB)
		}

		obj.Detail_retur_header = arr_DRB

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Sukses"
	res.Data = arrobj

	return res, nil
}

func ShowDetailRetur(id_retur int) (tools.Response, error) {
	var RTR retur.ShowDetailRetur
	var arr_RTR []retur.ShowDetailRetur
	var res tools.Response

	con := db.CreateCon()

	sqlStatement := "SELECT retur.id_retur, nama_toko, pelanggan.alamat, kota, no_telp, SUM(dr.jumlah), no_order, tanggal_retur, `user`.`name`, SUM(sub_total) FROM retur JOIN `user` ON retur.id_sales = `user`.id_user JOIN pelanggan ON pelanggan.id_pelanggan = retur.id_pelanggan JOIN detail_retur dr ON dr.id_retur=retur.id_retur WHERE retur.id_retur=?"

	err := con.QueryRow(sqlStatement, id_retur).Scan(&RTR.Id_retur, &RTR.Nama_toko,
		&RTR.Alamat, &RTR.Kota, &RTR.No_telp, &RTR.Total_barang,
		&RTR.No_order, &RTR.Tanggal_pembayaran, &RTR.Nama_sales,
		&RTR.Sub_total_retur)

	if err != nil {
		return res, err
	}

	sqlStatement = "SELECT down_payment, tanggal_pembayaran, pembayaran FROM `order` WHERE no_order=?"

	err = con.QueryRow(sqlStatement, RTR.No_order).Scan(&RTR.Tanggal_retur,
		&RTR.Pembayaran, &RTR.Down_payment)

	if err != nil {
		return res, err
	}

	var id retur.DetailRetur
	var arr_id []retur.DetailRetur

	sqlStatement = "SELECT id_detail_retur, nama_barang, jumlah, satuan, harga_jual, sub_total FROM detail_retur JOIN stock s ON s.id = detail_retur.id_stock WHERE detail_retur.id_retur=?"

	rows, err := con.Query(sqlStatement, id_retur)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {

		err = rows.Scan(&id.Id_detail_retur, &id.Nama_barang, &id.Jumlah,
			&id.Satuan, &id.Harga_jual, &id.Sub_total)

		if err != nil {
			return res, err
		}

		arr_id = append(arr_id, id)

	}

	for i := 0; i < len(arr_id); i++ {

		var D_R_B retur.DetailReturBarang
		var arr_D_R_B []retur.DetailReturBarang

		sqlStatement = "SELECT jenis_ukuran, jumlah, satuan FROM detail_retur_barang JOIN ukuran u ON u.id_ukuran = detail_retur_barang.id_ukuran WHERE id_detail_retur=?"

		rows, err = con.Query(sqlStatement, arr_id[i].Id_detail_retur)

		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {

			err = rows.Scan(&D_R_B.Jenis_ukuran, &D_R_B.Jumlah, &D_R_B.Satuan)

			if err != nil {
				return res, err
			}

			arr_D_R_B = append(arr_D_R_B, D_R_B)

		}

		arr_id[i].Detail_retur_barang = arr_D_R_B

	}

	RTR.Detail_order = arr_id

	arr_RTR = append(arr_RTR, RTR)

	if arr_RTR == nil {
		res.Status = http.StatusNotFound
		res.Message = "Not Found"
		res.Data = arr_RTR
	} else {
		res.Status = http.StatusOK
		res.Message = "Sukses"
		res.Data = arr_RTR
	}

	return res, nil
}
