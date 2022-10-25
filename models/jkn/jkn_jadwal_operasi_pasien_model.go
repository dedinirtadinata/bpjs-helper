package jkn

type ListJadwalOperasiPasienModel struct {
	List []JadwalOperasiPasienModel `json:"list"`
}

type JadwalOperasiPasienModel struct {
	Kodebooking    string `json:"kodebooking"`
	Tanggaloperasi string `json:"tanggaloperasi"`
	Jenistindakan  string `json:"jenistindakan"`
	Kodepoli       string `json:"kodepoli"`
	Namapoli       string `json:"namapoli"`
	Terlaksana     int    `json:"terlaksana"`
}
