package retur

type ShowID_DetailRetur struct {
	Id_detail_retur int `json:"Id_detail_retur "`
}

type ShowReturHeader struct {
	Id_retur            int                   `json:"id_retur"`
	Tanggal_retur       string                `json:"tanggal_retur"`
	Nama_toko           string                `json:"nama_toko"`
	Kota                string                `json:"kota"`
	Total_barang        string                `json:"total_barang"`
	Detail_retur_header []Detail_Retur_Header `json:"detail_retur_header"`
}

type Detail_Retur_Header struct {
	Id_detail_retur int    `json:"id_detail_retur"`
	Nama_barang     string `json:"nama_barang"`
	Jumlah          int    `json:"jumlah"`
	Satuan          string `json:"satuan"`
}

type ShowDetailRetur struct {
	Id_retur           int           `json:"id_retur"`
	Nama_toko          string        `json:"nama_toko"`
	Alamat             string        `json:"alamat"`
	Kota               string        `json:"kota"`
	No_telp            string        `json:"no_telp"`
	Total_barang       string        `json:"total_barang"`
	Tanggal_retur      string        `json:"tanggal_retur"`
	No_order           string        `json:"no_order"`
	Pembayaran         string        `json:"pembayaran"`
	Down_payment       string        `json:"down_payment"`
	Tanggal_pembayaran string        `json:"tanggal_pembayaran"`
	Nama_sales         string        `json:"nama_sales"`
	Sub_total_retur    int64         `json:"sub_total_retur"`
	Detail_order       []DetailRetur `json:"detail_Order"`
}

type DetailRetur struct {
	Id_detail_retur     int                 `json:"id_detail_retur"`
	Nama_barang         string              `json:"nama_barang"`
	Jumlah              int                 `json:"jumlah"`
	Satuan              string              `json:"satuan"`
	Harga_jual          int64               `json:"harga_jual"`
	Sub_total           int64               `json:"sub_total"`
	Detail_retur_barang []DetailReturBarang `json:"detail_retur_barang"`
}

type DetailReturBarang struct {
	Jenis_ukuran string `json:"jenis_ukuran"`
	Jumlah       int    `json:"jumlah"`
	Satuan       string `json:"satuan"`
}
