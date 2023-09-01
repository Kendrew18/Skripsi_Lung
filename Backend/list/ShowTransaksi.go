package list

type ShowTransaksi struct {
	IdTransaksi      int    `json:"id_transaksi"`
	Tanggal          string `json:"tanggal"`
	NamaSales        string `json:"nama_sales"`
	NamaPelanggan    string `json:"nama_pelanggan"`
	AlamatPelanggan  string `json:"alamat_pelanggan"`
	NoHpPelanggan    string `json:"no_hp_pelanggan"`
	TotalKeseluruhan int    `json:"total_keseluruhan"`
	Pembayaran       string `json:"pembayaran"`
}
