package models

type BPJSPelayananModel struct {
	Diagnosa     string `json:"diagnosa"`
	JnsPelayanan string `json:"jnsPelayanan"`
	KelasRawat   string `json:"kelasRawat"`
	NamaPeserta  string `json:"namaPeserta"`
	NoKartu      string `json:"noKartu"`
	NoSep        string `json:"noSep"`
	NoRujukan    string `json:"noRujukan"`
	Poli         string `json:"poli"`
	PpkPelayanan string `json:"ppkPelayanan"`
	TglPlgSep    string `json:"tglPlgSep"`
	TglSep       string `json:"tglSep"`
}
