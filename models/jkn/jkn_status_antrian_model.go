package jkn

type StatusAntrianModel struct {
	Namapoli        string `json:"namapoli"`
	Namadokter      string `json:"namadokter"`
	Totalantrean    int    `json:"totalantrean"`
	Sisaantrean     int    `json:"sisaantrean"`
	Antreanpanggil  string `json:"antreanpanggil"`
	Sisakuotajkn    int    `json:"sisakuotajkn"`
	Kuotajkn        int    `json:"kuotajkn"`
	Sisakuotanonjkn int    `json:"sisakuotanonjkn"`
	Kuotanonjkn     int    `json:"kuotanonjkn"`
	Keterangan      string `json:"keterangan"`
}
