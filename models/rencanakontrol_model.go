package models

type RequestRencanaKontrol struct {
	NoSEP             string `json:"noSEP"`
	KodeDokter        string `json:"kodeDokter"`
	PoliKontrol       string `json:"poliKontrol"`
	TglRencanaKontrol string `json:"tglRencanaKontrol"`
	User              string `json:"user"`
}

type RequestUpdateRencanaKontrol struct {
	NoSurat           string `json:"noSuratKontrol"`
	NoSEP             string `json:"noSEP"`
	KodeDokter        string `json:"kodeDokter"`
	PoliKontrol       string `json:"poliKontrol"`
	TglRencanaKontrol string `json:"tglRencanaKontrol"`
	User              string `json:"user"`
}

type RequestDeleteSuratKontrol struct {
	SuratKontrol RequestDeleteRencanaKontrolModel `json:"t_suratkontrol"`
}

type RequestDeleteRencanaKontrolModel struct {
	NoSurat string `json:"noSuratKontrol"`
	User    string `json:"user"`
}

type RequestSPRI struct {
	NoKartu           string `json:"noKartu"`
	KodeDokter        string `json:"kodeDokter"`
	PoliKontrol       string `json:"poliKontrol"`
	TglRencanaKontrol string `json:"tglRencanaKontrol"`
	User              string `json:"user"`
}

type ResponseRencanaKontrolModel struct {
	NoSuratKontrol    string      `json:"noSuratKontrol"`
	TglRencanaKontrol string      `json:"tglRencanaKontrol"`
	NamaDokter        string      `json:"namaDokter"`
	NoKartu           string      `json:"noKartu"`
	Nama              string      `json:"nama"`
	Kelamin           string      `json:"kelamin"`
	TglLahir          string      `json:"tglLahir"`
	NamaDiagnosa      interface{} `json:"namaDiagnosa"`
}

type ResponseSPRIModel struct {
	NoSPRI            string      `json:"noSPRI"`
	TglRencanaKontrol string      `json:"tglRencanaKontrol"`
	NamaDokter        string      `json:"namaDokter"`
	NoKartu           string      `json:"noKartu"`
	Nama              string      `json:"nama"`
	Kelamin           string      `json:"kelamin"`
	TglLahir          string      `json:"tglLahir"`
	NamaDiagnosa      interface{} `json:"namaDiagnosa"`
}

type ResponseListSuratKontrolModel struct {
	List []ResponseSuratKontrol `json:"list"`
}

type ResponseListDataSuratKontrolModel struct {
	List []ResponseSuratKontrol `json:"list"`
}

type ResponseSEPRencanaKontrol struct {
	NoSep        string `json:"noSep"`
	TglSep       string `json:"tglSep"`
	JnsPelayanan string `json:"jnsPelayanan"`
	Poli         string `json:"poli"`
	Diagnosa     string `json:"diagnosa"`
	Peserta      struct {
		NoKartu  string `json:"noKartu"`
		Nama     string `json:"nama"`
		TglLahir string `json:"tglLahir"`
		Kelamin  string `json:"kelamin"`
		HakKelas string `json:"hakKelas"`
	} `json:"peserta"`
	ProvUmum struct {
		KdProvider string `json:"kdProvider"`
		NmProvider string `json:"nmProvider"`
	} `json:"provUmum"`
	ProvPerujuk struct {
		KdProviderPerujuk string `json:"kdProviderPerujuk"`
		NmProviderPerujuk string `json:"nmProviderPerujuk"`
		AsalRujukan       string `json:"asalRujukan"`
		NoRujukan         string `json:"noRujukan"`
		TglRujukan        string `json:"tglRujukan"`
	} `json:"provPerujuk"`
}

type ListPoliSpesialistik struct {
	List []PoliSpesialistikModel `json:"list"`
}

type ListDokterKontrol struct {
	List []DokterKontrolModel `json:"list"`
}

//RencanaKontrol/ListSpesialistik/JnsKontrol/2/nomor/0059R0050422V000002/TglRencanaKontrol/2022-04-16
type PoliSpesialistikModel struct {
	KodePoli                    string `json:"kodePoli"`
	NamaPoli                    string `json:"namaPoli"`
	Kapasitas                   string `json:"kapasitas"`
	JmlRencanaKontroldanRujukan string `json:"jmlRencanaKontroldanRujukan"`
	Persentase                  string `json:"persentase"`
}

type DokterKontrolModel struct {
	KodeDokter    string `json:"kodeDokter"`
	NamaDokter    string `json:"namaDokter"`
	JadwalPraktek string `json:"jadwalPraktek"`
	Kapasitas     string `json:"kapasitas"`
}

type RencanaKontrolModel struct {
	NoSuratKontrol    string      `json:"noSuratKontrol"`
	TglRencanaKontrol string      `json:"tglRencanaKontrol"`
	TglTerbit         string      `json:"tglTerbit"`
	JnsKontrol        string      `json:"jnsKontrol"`
	PoliTujuan        string      `json:"poliTujuan"`
	NamaPoliTujuan    string      `json:"namaPoliTujuan"`
	KodeDokter        string      `json:"kodeDokter"`
	NamaDokter        string      `json:"namaDokter"`
	FlagKontrol       string      `json:"flagKontrol"`
	KodeDokterPembuat interface{} `json:"kodeDokterPembuat"`
	NamaDokterPembuat interface{} `json:"namaDokterPembuat"`
	NamaJnsKontrol    string      `json:"namaJnsKontrol"`
	Sep               struct {
		NoSep        string `json:"noSep"`
		TglSep       string `json:"tglSep"`
		JnsPelayanan string `json:"jnsPelayanan"`
		Poli         string `json:"poli"`
		Diagnosa     string `json:"diagnosa"`
		Peserta      struct {
			NoKartu  string `json:"noKartu"`
			Nama     string `json:"nama"`
			TglLahir string `json:"tglLahir"`
			Kelamin  string `json:"kelamin"`
			HakKelas string `json:"hakKelas"`
		} `json:"peserta"`
		ProvUmum struct {
			KdProvider string `json:"kdProvider"`
			NmProvider string `json:"nmProvider"`
		} `json:"provUmum"`
		ProvPerujuk struct {
			KdProviderPerujuk string `json:"kdProviderPerujuk"`
			NmProviderPerujuk string `json:"nmProviderPerujuk"`
			AsalRujukan       string `json:"asalRujukan"`
			NoRujukan         string `json:"noRujukan"`
			TglRujukan        string `json:"tglRujukan"`
		} `json:"provPerujuk"`
	} `json:"sep"`
}

type ResponseSuratKontrol struct {
	NoSuratKontrol    string `json:"noSuratKontrol"`
	JnsPelayanan      string `json:"jnsPelayanan"`
	JnsKontrol        string `json:"jnsKontrol"`
	NamaJnsKontrol    string `json:"namaJnsKontrol"`
	TglRencanaKontrol string `json:"tglRencanaKontrol"`
	TglTerbitKontrol  string `json:"tglTerbitKontrol"`
	NoSepAsalKontrol  string `json:"noSepAsalKontrol"`
	PoliAsal          string `json:"poliAsal"`
	NamaPoliAsal      string `json:"namaPoliAsal"`
	PoliTujuan        string `json:"poliTujuan"`
	NamaPoliTujuan    string `json:"namaPoliTujuan"`
	TglSEP            string `json:"tglSEP"`
	KodeDokter        string `json:"kodeDokter"`
	NamaDokter        string `json:"namaDokter"`
	NoKartu           string `json:"noKartu"`
	Nama              string `json:"nama"`
	TerbitSEP         string `json:"terbitSEP"`
}
