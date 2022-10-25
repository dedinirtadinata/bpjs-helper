package models

type KunjunganModel struct {
	Diagnosa     string      `json:"diagnosa"`
	JnsPelayanan string      `json:"jnsPelayanan"`
	KelasRawat   string      `json:"kelasRawat"`
	Nama         string      `json:"nama"`
	NoKartu      string      `json:"noKartu"`
	NoSep        string      `json:"noSep"`
	NoRujukan    string      `json:"noRujukan"`
	Poli         interface{} `json:"poli"`
	TglPlgSep    string      `json:"tglPlgSep"`
	TglSep       string      `json:"tglSep"`
}
