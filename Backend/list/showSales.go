package list

type SalesNames struct {
	IdSales     int     `json:"id_sales"`
	NamaSales   string  `json:"nama_sales"`
	Alamat      string  `json:"alamat"`
	Nomor_hp    string  `json:"nomor_hp"`
	Bank        string  `json:"bank"`
	No_rekening string  `json:"no_rekening"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
}
