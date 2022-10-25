package jkn

type ListJadwalOperasiModel struct {
	List []JadwalOperasiModel `json:"list"`
}

type JadwalOperasiModel struct {
	Kodebooking    string `json:"kodebooking"`
	Tanggaloperasi string `json:"tanggaloperasi"`
	Jenistindakan  string `json:"jenistindakan"`
	Kodepoli       string `json:"kodepoli"`
	Namapoli       string `json:"namapoli"`
	Terlaksana     int    `json:"terlaksana"`
	Nopeserta      string `json:"nopeserta"`
	Lastupdate     int64  `json:"lastupdate"`
}
