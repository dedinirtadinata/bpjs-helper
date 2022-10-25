package models

type BPJSRuanganModel struct {
	Kodekelas          string `gorm:"column:kodekelas";json:"kodekelas"`
	Koderuang          string `gorm:"column:koderuang";json:"koderuang"`
	Namaruang          string `gorm:"column:namaruang";json:"namaruang"`
	Kapasitas          string `gorm:"column:kapasitas";json:"kapasitas"`
	Tersedia           string `gorm:"column:tersedia";json:"tersedia"`
	Tersediapria       string `json:"tersediapria"`
	Tersediawanita     string `json:"tersediawanita"`
	Tersediapriawanita string `json:"tersediapriawanita"`
}
type BPJSRuanganModel2 struct {
	Kodekelas          string `json:"kodekelas"`
	Koderuang          string `json:"koderuang"`
	Namaruang          string `json:"namaruang"`
	Kapasitas          string `json:"kapasitas"`
	Tersedia           string `json:"tersedia"`
	Tersediapria       string `json:"tersediapria"`
	Tersediawanita     string `json:"tersediawanita"`
	Tersediapriawanita string `json:"tersediapriawanita"`
}

type KetersediaanRuangModel struct {
	Tersedia           int    `json:"tersedia"`
	Kodekelas          string `json:"kodekelas"`
	Namakelas          string `json:"namakelas"`
	Lastupdateall      string `json:"lastupdateall"`
	Tersediapria       int    `json:"tersediapria"`
	Tersediawanita     int    `json:"tersediawanita"`
	Koderuang          string `json:"koderuang"`
	Kodeppk            string `json:"kodeppk"`
	Tersediapriawanita int    `json:"tersediapriawanita"`
	Namaruang          string `json:"namaruang"`
	Rownumber          int    `json:"rownumber"`
	Kapasitas          int    `json:"kapasitas"`
	Lastupdate         string `json:"lastupdate"`
}

type RuangDeleteModel struct {
	Kodekelas string `json:"kodekelas"`
	Koderuang string `json:"koderuang"`
}

type KelasBPJSModel struct {
	Kodekelas string `json:"kodekelas"`
	Namakelas string `json:"namakelas"`
}

type ResponseListBPJSKelasModel struct {
	List []KelasBPJSModel `json:"list"`
}

type ResponseListKetersediaanRuangModel struct {
	List []KetersediaanRuangModel `json:"list"`
}

type ResponseInformasiKetersediaanKamarRS struct {
	Metadata struct {
		Code       int    `json:"code"`
		Message    string `json:"message"`
		Totalitems int    `json:"totalitems"`
	} `json:"metadata"`
	Response struct {
		List []struct {
			Tersedia           int    `json:"tersedia"`
			Kodekelas          string `json:"kodekelas"`
			Namakelas          string `json:"namakelas"`
			Lastupdateall      string `json:"lastupdateall"`
			Tersediapria       int    `json:"tersediapria"`
			Tersediawanita     int    `json:"tersediawanita"`
			Koderuang          string `json:"koderuang"`
			Kodeppk            string `json:"kodeppk"`
			Tersediapriawanita int    `json:"tersediapriawanita"`
			Namaruang          string `json:"namaruang"`
			Rownumber          int    `json:"rownumber"`
			Kapasitas          int    `json:"kapasitas"`
			Lastupdate         string `json:"lastupdate"`
		} `json:"list"`
	} `json:"response"`
}
