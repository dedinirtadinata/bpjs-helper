package models

type PoliResponse struct {
	Metadata struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"metaData"`
	Response struct {
		Poli []struct {
			Kode string `json:"kode"`
			Nama string `json:"nama"`
		} `json:"poli"`
	} `json:"response"`
}

type ReferensiPoliModel struct {
	Nmpoli         string `gorm:"column:nama" json:"nmpoli"`
	Nmsubspesialis string `gorm:"column:namasubspesialis" json:"nmsubspesialis"`
	Kdsubspesialis string `gorm:"column:kdsubspesialis" json:"kdsubspesialis"`
	Kdpoli         string `gorm:"column:kode" json:"kdpoli"`
}

func (ReferensiPoliModel) TableName() string {
	return "bpjs_poli"
}

type ReferensiDokterModel struct {
	Namadokter string `json:"namadokter"`
	Kodedokter int    `json:"kodedokter"`
}

type ResponseListReferensiDokterModel struct {
	List []ReferensiDokterModel `json:"list"`
}

type BPJSJadwalDokterModel struct {
	Kodesubspesialis string `json:"kodesubspesialis"`
	Hari             int    `json:"hari"`
	Kapasitaspasien  int    `json:"kapasitaspasien"`
	Libur            int    `json:"libur"`
	Namahari         string `json:"namahari"`
	Jadwal           string `json:"jadwal"`
	Namasubspesialis string `json:"namasubspesialis"`
	Namadokter       string `json:"namadokter"`
	Kodepoli         string `json:"kodepoli"`
	Namapoli         string `json:"namapoli"`
	Kodedokter       int    `json:"kodedokter"`
}

type ResponseListBPJSJadwalDokterModel struct {
	List []BPJSJadwalDokterModel `json:"list"`
}
