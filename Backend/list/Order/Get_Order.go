package order

type ShowOrderHeader struct {
	Id_order     int    `json:"id_order"`
	No_order     string `json:"no_order"`
	Sub_total    int64  `json:"sub_total"`
	Nama_sales   string `json:"nama_sales"`
	Nama_toko    string `json:"nama_toko"`
	Kota         string `json:"kota"`
	Status_order string `json:"status_order"`
	Total_barang string `json:"total_barang"`
}
type ShowDetailTransaksi struct {
	Id_order           int            `json:"id_order"`
	Nama_toko          string         `json:"nama_toko"`
	Alamat             string         `json:"alamat"`
	Kota               string         `json:"kota"`
	No_telp            string         `json:"no_telp"`
	Total_barang       string         `json:"total_barang"`
	No_order           string         `json:"no_order"`
	Tanggal_pesanan    string         `json:"tanggal_pesanan"`
	Tanggal_pengiriman string         `json:"tanggal_pengiriman"`
	Pembayaran         string         `json:"pembayaran"`
	Down_payment       string         `json:"down_payment"`
	Tanggal_pembayaran string         `json:"tanggal_pembayaran"`
	Nama_sales         string         `json:"nama_sales"`
	Sub_total          int64          `json:"sub_total"`
	Retur              int64          `json:"retur"`
	Status_order       string         `json:"status_order"`
	Detail_order       []Detail_Order `json:"detail_Order"`
}

type Detail_Order struct {
	Id_detail_order     int                   `json:"Id_detail_order"`
	Nama_barang         string                `json:"Nama_barang"`
	Jumlah              int                   `json:"jumlah"`
	Satuan              string                `json:"satuan"`
	Harga_jual          int64                 `json:"harga_jual"`
	Sub_total           int64                 `json:"sub_total"`
	Detail_order_barang []Detail_Order_Barang `json:"detail_order_barang"`
}

type Detail_Order_Barang struct {
	Jenis_ukuran string `json:"jenis_ukuran"`
	Jumlah       int    `json:"jumlah"`
	Satuan       string `json:"satuan"`
}

type ShowID_DetailOrderBarang struct {
	Id_detail_order int `json:"Id_detail_order "`
}
