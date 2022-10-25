package models

type ResponsePeserta struct {
	Metadata struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	} `json:"metadata"`

	Peserta DataPeserta `json:"response"`
}

type DataPesertaBPJS struct {
	Peserta struct {
		Cob struct {
			NmAsuransi interface{} `json:"nmAsuransi"`
			NoAsuransi interface{} `json:"noAsuransi"`
			TglTAT     interface{} `json:"tglTAT"`
			TglTMT     interface{} `json:"tglTMT"`
		} `json:"cob"`
		HakKelas struct {
			Keterangan string `json:"keterangan"`
			Kode       string `json:"kode"`
		} `json:"hakKelas"`
		Informasi struct {
			Dinsos      interface{} `json:"dinsos"`
			NoSKTM      interface{} `json:"noSKTM"`
			ProlanisPRB interface{} `json:"prolanisPRB"`
		} `json:"informasi"`
		JenisPeserta struct {
			Keterangan string `json:"keterangan"`
			Kode       string `json:"kode"`
		} `json:"jenisPeserta"`
		Mr struct {
			NoMR      interface{} `json:"noMR"`
			NoTelepon interface{} `json:"noTelepon"`
		} `json:"mr"`
		Nama     string `json:"nama"`
		Nik      string `json:"nik"`
		NoKartu  string `json:"noKartu"`
		Pisa     string `json:"pisa"`
		ProvUmum struct {
			KdProvider string `json:"kdProvider"`
			NmProvider string `json:"nmProvider"`
		} `json:"provUmum"`
		Sex           string `json:"sex"`
		StatusPeserta struct {
			Keterangan string `json:"keterangan"`
			Kode       string `json:"kode"`
		} `json:"statusPeserta"`
		TglCetakKartu string `json:"tglCetakKartu"`
		TglLahir      string `json:"tglLahir"`
		TglTAT        string `json:"tglTAT"`
		TglTMT        string `json:"tglTMT"`
		Umur          struct {
			UmurSaatPelayanan string `json:"umurSaatPelayanan"`
			UmurSekarang      string `json:"umurSekarang"`
		} `json:"umur"`
	} `json:"peserta"`
}

type DataPeserta struct {
	NoKartu string `json:"noKartu"`
	Nik     string `json:"nik"`
	Nama    string `json:"nama"`
	Pisa    string `json:"pisa"`
	Sex     string `json:"sex"`
	Mr      struct {
		NoMR      interface{} `json:"noMR"`
		NoTelepon interface{} `json:"noTelepon"`
	} `json:"mr"`
	TglLahir      string `json:"tglLahir"`
	TglCetakKartu string `json:"tglCetakKartu"`
	TglTAT        string `json:"tglTAT"`
	TglTMT        string `json:"tglTMT"`
	StatusPeserta struct {
		Kode       string `json:"kode"`
		Keterangan string `json:"keterangan"`
	} `json:"statusPeserta"`
	ProvUmum struct {
		KdProvider string `json:"kdProvider"`
		NmProvider string `json:"nmProvider"`
	} `json:"provUmum"`
	JenisPeserta struct {
		Kode       string `json:"kode"`
		Keterangan string `json:"keterangan"`
	} `json:"jenisPeserta"`
	HakKelas struct {
		Kode       string `json:"kode"`
		Keterangan string `json:"keterangan"`
	} `json:"hakKelas"`
	Umur struct {
		UmurSekarang      string `json:"umurSekarang"`
		UmurSaatPelayanan string `json:"umurSaatPelayanan"`
	} `json:"umur"`
	Informasi struct {
		Dinsos      interface{} `json:"dinsos"`
		ProlanisPRB interface{} `json:"prolanisPRB"`
		NoSKTM      interface{} `json:"noSKTM"`
	} `json:"informasi"`
	Cob struct {
		NoAsuransi interface{} `json:"noAsuransi"`
		NmAsuransi interface{} `json:"nmAsuransi"`
		TglTMT     interface{} `json:"tglTMT"`
		TglTAT     interface{} `json:"tglTAT"`
	} `json:"cob"`
}
