package models

type BPJSKlaimModel struct {
	Inacbg struct {
		Kode string `json:"kode"`
		Nama string `json:"nama"`
	} `json:"Inacbg"`
	Biaya struct {
		ByPengajuan   string `json:"byPengajuan"`
		BySetujui     string `json:"bySetujui"`
		ByTarifGruper string `json:"byTarifGruper"`
		ByTarifRS     string `json:"byTarifRS"`
		ByTopup       string `json:"byTopup"`
	} `json:"biaya"`
	KelasRawat string `json:"kelasRawat"`
	NoFPK      string `json:"noFPK"`
	NoSEP      string `json:"noSEP"`
	Peserta    struct {
		Nama    string `json:"nama"`
		NoKartu string `json:"noKartu"`
		NoMR    string `json:"noMR"`
	} `json:"peserta"`
	Poli      string `json:"poli"`
	Status    string `json:"status"`
	TglPulang string `json:"tglPulang"`
	TglSep    string `json:"tglSep"`
}
