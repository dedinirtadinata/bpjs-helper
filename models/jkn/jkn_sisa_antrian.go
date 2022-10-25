package jkn

type SisaAntrianModel struct {
	Nomorantrean   string `json:"nomorantrean"`
	Namapoli       string `json:"namapoli"`
	Namadokter     string `json:"namadokter"`
	Sisaantrean    int    `json:"sisaantrean"`
	Antreanpanggil string `json:"antreanpanggil"`
	Waktutunggu    int    `json:"waktutunggu"`
	Keterangan     string `json:"keterangan"`
}
