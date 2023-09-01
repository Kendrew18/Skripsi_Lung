package list

type ShowDetailTransaksi struct {
	IdTransaksi       int    `json:"id_transaksi"`
	IdDetailTransaksi int    `json:"id_detail_transaksi"`
	IdProduct         int    `json:"id_product"`
	NamaProduct       string `json:"nama_product"`
	JenisProduct      string `json:"jenis_product"`
	Jumlah            int    `json:"jumlah"`
	Total             int    `json:"total"`
}
