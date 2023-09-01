package list

type ShowPelanggan struct {
	Id_pelanggan         int    `json:"id_pelanggan"`
	Nama_toko            string `json:"nama_toko"`
	No_telp              string `json:"no_telp"`
	Alamat               string `json:"alamat"`
	Kota                 string `json:"kota"`
	Provinsi             string `json:"provinsi"`
	Nama_penanggungjawab string `json:"nama_penanggungjawab"`
}
