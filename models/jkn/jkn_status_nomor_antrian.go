package jkn

type InformasiNomorAntrian struct {
	Nomorantrean     string `json:"nomorantrean"`
	Angkaantrean     int    `json:"angkaantrean"`
	Kodebooking      string `json:"kodebooking"`
	Norm             string `json:"norm"`
	Namapoli         string `json:"namapoli"`
	Namadokter       string `json:"namadokter"`
	Estimasidilayani int64  `json:"estimasidilayani"`
	Sisakuotajkn     int    `json:"sisakuotajkn"`
	Kuotajkn         int    `json:"kuotajkn"`
	Sisakuotanonjkn  int    `json:"sisakuotanonjkn"`
	Kuotanonjkn      int    `json:"kuotanonjkn"`
	Keterangan       string `json:"keterangan"`
}
