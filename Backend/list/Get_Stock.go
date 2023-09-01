package list

type ShowAllProduct struct {
	IdProduct    int      `json:"id_product"`
	NamaProduct  string   `json:"nama_product"`
	JenisProduct string   `json:"jenis_product"`
	Harga        int64    `json:"harga"`
	Ukuran       []Ukuran `json:"ukuran"`
}
type Ukuran struct {
	UkuranProduct string `json:"ukuran_product"`
	Stock         int    `json:"stock"`
}
