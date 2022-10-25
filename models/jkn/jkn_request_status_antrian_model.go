package jkn

type RequestStatusAntrianModel struct {
	Kodepoli       string `json:"kodepoli"`
	Kodedokter     int    `json:"kodedokter"`
	Tanggalperiksa string `json:"tanggalperiksa"`
	Jampraktek     string `json:"jampraktek"`
}
