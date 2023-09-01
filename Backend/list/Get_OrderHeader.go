package list

type ShowOrderHeader struct {
	Id_order     int    `json:"id_order"`
	No_order     string `json:"no_order"`
	Sub_total    int64  `json:"sub_total"`
	Nama_sales   string `json:"nama_sales"`
	Nama_toko    string `json:"nama_toko"`
	Kota         string `json:"kota"`
	Total_barang string `json:"total_barang"`
}
