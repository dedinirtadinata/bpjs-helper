package models

type ListRujukan struct {
	List []RujukanModel `json:"rujukan"`
}

type ResponseListSaranaRujukan struct {
	List []SaranaRujukanModel `json:"list"`
}

type ResponseListSpesialistikModel struct {
	List []DataRujukanSpesialistikModel `json:"list"`
}

type RequestUpdateRujukanModel struct {
	Rujukan UpdateRujukanModel `json:"t_rujukan"`
}

type RequestDeleteRujukanModel struct {
	Rujukan DeleteRujukanModel `json:"t_rujukan"`
}

type ResponseRujukanKeluarRSModel struct {
	Rujukan RujukanRSModel `json:"rujukan"`
}

type ResponseRujukan struct {
	AsalFaskes  string       `json:"asalFaskes"`
	DataRujukan RujukanModel `json:"rujukan"`
}

type RequestRujukanSpesialistik struct {
	Faskes  string `json:"faskes" form:"faskes"`
	Tanggal string `json:"tgl" form:"tgl"`
}

type RequestRujukanSarana struct {
	Faskes string `json:"faskes" form:"faskes"`
}

type RequestRujukan struct {
	Rujukan PostRujukanModel `json:"t_rujukan" form:"t_rujukan"`
}

type RequestRujukanRS struct {
	NoRujukan string `json:"norujukan" form:"norujukan"`
}

type RujukanRSModel struct {
	NoRujukan           string `json:"noRujukan"`
	NoSep               string `json:"noSep"`
	NoKartu             string `json:"noKartu"`
	Nama                string `json:"nama"`
	KelasRawat          string `json:"kelasRawat"`
	Kelamin             string `json:"kelamin"`
	TglLahir            string `json:"tglLahir"`
	TglSep              string `json:"tglSep"`
	TglRujukan          string `json:"tglRujukan"`
	TglRencanaKunjungan string `json:"tglRencanaKunjungan"`
	PpkDirujuk          string `json:"ppkDirujuk"`
	NamaPpkDirujuk      string `json:"namaPpkDirujuk"`
	JnsPelayanan        string `json:"jnsPelayanan"`
	Catatan             string `json:"catatan"`
	DiagRujukan         string `json:"diagRujukan"`
	NamaDiagRujukan     string `json:"namaDiagRujukan"`
	TipeRujukan         string `json:"tipeRujukan"`
	NamaTipeRujukan     string `json:"namaTipeRujukan"`
	PoliRujukan         string `json:"poliRujukan"`
	NamaPoliRujukan     string `json:"namaPoliRujukan"`
}

type SaranaRujukanModel struct {
	Kode string `json:"kodeSarana"`
	Nama string `json:"namaSarana"`
}

type PostRujukanModel struct {
	NoSep               string `json:"noSep"`
	TglRujukan          string `json:"tglRujukan"`
	TglRencanaKunjungan string `json:"tglRencanaKunjungan"`
	PpkDirujuk          string `json:"ppkDirujuk"`
	JnsPelayanan        string `json:"jnsPelayanan"`
	Catatan             string `json:"catatan"`
	DiagRujukan         string `json:"diagRujukan"`
	TipeRujukan         string `json:"tipeRujukan"`
	PoliRujukan         string `json:"poliRujukan"`
	User                string `json:"user"`
}

type DeleteRujukanModel struct {
	NoRujukan string `json:"noRujukan"`
	User      string `json:"user"`
}

type UpdateRujukanModel struct {
	NoRujukan           string `json:"noRujukan"`
	TglRujukan          string `json:"tglRujukan"`
	TglRencanaKunjungan string `json:"tglRencanaKunjungan"`
	PpkDirujuk          string `json:"ppkDirujuk"`
	JnsPelayanan        string `json:"jnsPelayanan"`
	Catatan             string `json:"catatan"`
	DiagRujukan         string `json:"diagRujukan"`
	TipeRujukan         string `json:"tipeRujukan"`
	PoliRujukan         string `json:"poliRujukan"`
	User                string `json:"user"`
}

type PostRujukanKhususModel struct {
	NoRujukan string `json:"noRujukan"`
	Diagnosa  []struct {
		Kode string `json:"kode"`
	} `json:"diagnosa"`
	Procedure []struct {
		Kode string `json:"kode"`
	} `json:"procedure"`
	User string `json:"user"`
}

type DataRujukanSpesialistikModel struct {
	Kode          string `json:"kodeSpesialis"`
	Nama          string `json:"namaSpesialis"`
	Kapasitas     string `json:"kapasitas"`
	JumlahRujukan string `json:"jumlahRujukan"`
	Persentase    string `json:"persentase"`
}

type DataRujukanModel struct {
	Norujukan          string `json:"norujukan"`
	Nokapst            string `json:"nokapst"`
	Nmpst              string `json:"nmpst"`
	Diagppk            string `json:"diagppk"`
	TglrujukanAwal     string `json:"tglrujukan_awal"`
	TglrujukanBerakhir string `json:"tglrujukan_berakhir"`
}

type RujukanModel struct {
	Diagnosa struct {
		Kode string `json:"kode"`
		Nama string `json:"nama"`
	} `json:"diagnosa"`
	Keluhan     string `json:"keluhan"`
	NoKunjungan string `json:"noKunjungan"`
	Pelayanan   struct {
		Kode string `json:"kode"`
		Nama string `json:"nama"`
	} `json:"pelayanan"`
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
			NoMR      string      `json:"noMR"`
			NoTelepon interface{} `json:"noTelepon"`
		} `json:"mr"`
		Nama     string      `json:"nama"`
		Nik      interface{} `json:"nik"`
		NoKartu  string      `json:"noKartu"`
		Pisa     string      `json:"pisa"`
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
	PoliRujukan struct {
		Kode string `json:"kode"`
		Nama string `json:"nama"`
	} `json:"poliRujukan"`
	ProvPerujuk struct {
		Kode string `json:"kode"`
		Nama string `json:"nama"`
	} `json:"provPerujuk"`
	TglKunjungan string `json:"tglKunjungan"`
}
