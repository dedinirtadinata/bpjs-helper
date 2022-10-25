package jkn

type RequestAmbilAntrianModel struct {
	Nomorkartu     string `json:"nomorkartu"`
	Nik            string `json:"nik"`
	Nohp           string `json:"nohp"`
	Kodepoli       string `json:"kodepoli"`
	Norm           string `json:"norm"`
	Tanggalperiksa string `json:"tanggalperiksa"`
	Kodedokter     int    `json:"kodedokter"`
	Jampraktek     string `json:"jampraktek"`
	Jeniskunjungan int    `json:"jeniskunjungan"`
	Nomorreferensi string `json:"nomorreferensi"`
}
