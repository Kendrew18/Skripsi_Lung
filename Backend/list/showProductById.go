package list

type ShowUkuranProduct struct {
	NamaProduct int `json:"id_detail_ukuran"`
	IdProduct   int `json:"id_product"`
	IdUkuran    int `json:"id_ukuran"`
	Stock       int `json:"stock"`
}
