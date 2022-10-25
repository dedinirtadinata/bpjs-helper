package models

type RequestSEP struct {
	Sep PostSepModel `json:"t_sep"`
}

type RequestUpdateSEP struct {
	Sep UpdatedSEPModel `json:"t_sep"`
}

type RequestPengajuan struct {
	Pengajuan PengajuanModel `json:"t_sep"`
}

type RequestUpdatePulang struct {
	UpdatePulang UpdatePulangModel `json:"t_sep"`
}

type RequestDeleteSep struct {
	Sep DeleteSepModel `json:"t_sep"`
}

type RequestDeleteSepInternal struct {
	SepInternal DeleteSepInternalModel `json:"t_sep"`
}

type UpdatePulangModel struct {
	NoSep            string `json:"noSep"`
	StatusPulang     string `json:"statusPulang"`
	NoSuratMeninggal string `json:"noSuratMeninggal"`
	TglMeninggal     string `json:"tglMeninggal"`
	TglPulang        string `json:"tglPulang"`
	NoLaporan        string `json:"noLPManual"`
	User             string `json:"user"`
}

type DeleteSepModel struct {
	NoSep string `json:"noSep"`
	User  string `json:"user"`
}

type DeleteSepInternalModel struct {
	NoSep      string `json:"noSep"`
	NoSurat    string `json:"noSurat"`
	TglRujukan string `json:"tglRujukanInternal"`
	PoliTujuan string `json:"kdPoliTuj"`
	User       string `json:"user"`
}

type PengajuanModel struct {
	NoKartu      string `json:"noKartu"`
	TglSep       string `json:"tglSep"`
	JnsPelayanan string `json:"jnsPelayanan"`
	JnsPengajuan string `json:"jnsPengajuan"`
	Keterangan   string `json:"keterangan"`
	User         string `json:"user"`
}

type UpdatedSEPModel struct {
	NoSep    string `json:"noSep"`
	KlsRawat struct {
		KlsRawatHak     string `json:"klsRawatHak"`
		KlsRawatNaik    string `json:"klsRawatNaik"`
		Pembiayaan      string `json:"pembiayaan"`
		PenanggungJawab string `json:"penanggungJawab"`
	} `json:"klsRawat"`
	NoMR     string `json:"noMR"`
	Catatan  string `json:"catatan"`
	DiagAwal string `json:"diagAwal"`
	Poli     struct {
		Tujuan    string `json:"tujuan"`
		Eksekutif string `json:"eksekutif"`
	} `json:"poli"`
	Cob struct {
		Cob string `json:"cob"`
	} `json:"cob"`
	Katarak struct {
		Katarak string `json:"katarak"`
	} `json:"katarak"`
	Jaminan struct {
		LakaLantas string `json:"lakaLantas"`
		Penjamin   struct {
			TglKejadian string `json:"tglKejadian"`
			Keterangan  string `json:"keterangan"`
			Suplesi     struct {
				Suplesi      string `json:"suplesi"`
				NoSepSuplesi string `json:"noSepSuplesi"`
				LokasiLaka   struct {
					KdPropinsi  string `json:"kdPropinsi"`
					KdKabupaten string `json:"kdKabupaten"`
					KdKecamatan string `json:"kdKecamatan"`
				} `json:"lokasiLaka"`
			} `json:"suplesi"`
		} `json:"penjamin"`
	} `json:"jaminan"`
	DpjpLayan string `json:"dpjpLayan"`
	NoTelp    string `json:"noTelp"`
	User      string `json:"user"`
}

type PostSepModel struct {
	NoKartu      string `json:"noKartu"`
	TglSep       string `json:"tglSep"`
	PpkPelayanan string `json:"ppkPelayanan"`
	JnsPelayanan string `json:"jnsPelayanan"`
	KlsRawat     struct {
		KlsRawatHak     string `json:"klsRawatHak"`
		KlsRawatNaik    string `json:"klsRawatNaik"`
		Pembiayaan      string `json:"pembiayaan"`
		PenanggungJawab string `json:"penanggungJawab"`
	} `json:"klsRawat"`
	NoMR    string `json:"noMR"`
	Rujukan struct {
		AsalRujukan string `json:"asalRujukan"`
		TglRujukan  string `json:"tglRujukan"`
		NoRujukan   string `json:"noRujukan"`
		PpkRujukan  string `json:"ppkRujukan"`
	} `json:"rujukan"`
	Catatan  string `json:"catatan"`
	DiagAwal string `json:"diagAwal"`
	Poli     struct {
		Tujuan    string `json:"tujuan"`
		Eksekutif string `json:"eksekutif"`
	} `json:"poli"`
	Cob struct {
		Cob string `json:"cob"`
	} `json:"cob"`
	Katarak struct {
		Katarak string `json:"katarak"`
	} `json:"katarak"`
	Jaminan struct {
		LakaLantas string `json:"lakaLantas"`
		NoLP       string `json:"noLP"`
		Penjamin   struct {
			TglKejadian string `json:"tglKejadian"`
			Keterangan  string `json:"keterangan"`
			Suplesi     struct {
				Suplesi      string `json:"suplesi"`
				NoSepSuplesi string `json:"noSepSuplesi"`
				LokasiLaka   struct {
					KdPropinsi  string `json:"kdPropinsi"`
					KdKabupaten string `json:"kdKabupaten"`
					KdKecamatan string `json:"kdKecamatan"`
				} `json:"lokasiLaka"`
			} `json:"suplesi"`
		} `json:"penjamin"`
	} `json:"jaminan"`
	TujuanKunj    string `json:"tujuanKunj"`
	FlagProcedure string `json:"flagProcedure"`
	KdPenunjang   string `json:"kdPenunjang"`
	AssesmentPel  string `json:"assesmentPel"`
	Skdp          struct {
		NoSurat  string `json:"noSurat"`
		KodeDPJP string `json:"kodeDPJP"`
	} `json:"skdp"`
	DpjpLayan string `json:"dpjpLayan"`
	NoTelp    string `json:"noTelp"`
	User      string `json:"user"`
}

type ResponseSEP struct {
	sep ResponseSEPModel `json:"sep"`
}

type ResponseSEPModel struct {
	NoSep         string `json:"noSep"`
	TglSep        string `json:"tglSep"`
	JnsPelayanan  string `json:"jnsPelayanan"`
	KelasRawat    string `json:"kelasRawat"`
	Diagnosa      string `json:"diagnosa"`
	NoRujukan     string `json:"noRujukan"`
	Poli          string `json:"poli"`
	PoliEksekutif string `json:"poliEksekutif"`
	Catatan       string `json:"catatan"`
	Penjamin      string `json:"penjamin"`
	Peserta       struct {
		Asuransi   string `json:"asuransi"`
		HakKelas   string `json:"hakKelas"`
		JnsPeserta string `json:"jnsPeserta"`
		Kelamin    string `json:"kelamin"`
		Nama       string `json:"nama"`
		NoKartu    string `json:"noKartu"`
		NoMr       string `json:"noMr"`
		TglLahir   string `json:"tglLahir"`
	} `json:"peserta"`
	Informasi struct {
		Dinsos      interface{} `json:"Dinsos"`
		ProlanisPRB interface{} `json:"prolanisPRB"`
		NoSKTM      interface{} `json:"noSKTM"`
	} `json:"informasi:"`
}

type ResponseSEPInternalModel struct {
	Count string                 `json:"count"`
	List  []DataSEPInternalModel `json:"list"`
}

type DataSEPModel struct {
	NoSep              string `json:"noSep"`
	TglSep             string `json:"tglSep"`
	JnsPelayanan       string `json:"jnsPelayanan"`
	KelasRawat         string `json:"kelasRawat"`
	Diagnosa           string `json:"diagnosa"`
	NoRujukan          string `json:"noRujukan"`
	Poli               string `json:"poli"`
	PoliEksekutif      string `json:"poliEksekutif"`
	Catatan            string `json:"catatan"`
	Penjamin           string `json:"penjamin"`
	KdStatusKecelakaan string `json:"kdStatusKecelakaan"`
	NmstatusKecelakaan string `json:"nmstatusKecelakaan"`
	LokasiKejadian     struct {
		KdKab       string `json:"kdKab"`
		KdKec       string `json:"kdKec"`
		KdProp      string `json:"kdProp"`
		KetKejadian string `json:"ketKejadian"`
		Lokasi      string `json:"lokasi"`
		TglKejadian string `json:"tglKejadian"`
	} `json:"lokasiKejadian"`
	Dpjp struct {
		KdDPJP string `json:"kdDPJP"`
		NmDPJP string `json:"nmDPJP"`
	} `json:"dpjp"`
	Peserta struct {
		Asuransi   interface{} `json:"asuransi"`
		HakKelas   string      `json:"hakKelas"`
		JnsPeserta string      `json:"jnsPeserta"`
		Kelamin    string      `json:"kelamin"`
		Nama       string      `json:"nama"`
		NoKartu    string      `json:"noKartu"`
		NoMr       string      `json:"noMr"`
		TglLahir   string      `json:"tglLahir"`
	} `json:"peserta"`
	KlsRawat struct {
		KlsRawatHak     string `json:"klsRawatHak"`
		KlsRawatNaik    string `json:"klsRawatNaik"`
		Pembiayaan      string `json:"pembiayaan"`
		PenanggungJawab string `json:"penanggungJawab"`
	} `json:"klsRawat"`
	Kontrol struct {
		KdDokter string `json:"kdDokter"`
		NmDokter string `json:"nmDokter"`
		NoSurat  string `json:"noSurat"`
	} `json:"kontrol"`
	Cob     string `json:"cob"`
	Katarak string `json:"katarak"`
}

type DataSEPInternalModel struct {
	Tujuanrujuk      string      `json:"tujuanrujuk"`
	Nmtujuanrujuk    string      `json:"nmtujuanrujuk"`
	Nmpoliasal       string      `json:"nmpoliasal"`
	Tglrujukinternal string      `json:"tglrujukinternal"`
	Nosep            string      `json:"nosep"`
	Nosepref         interface{} `json:"nosepref"`
	Ppkpelsep        string      `json:"ppkpelsep"`
	Nokapst          string      `json:"nokapst"`
	Tglsep           string      `json:"tglsep"`
	Nosurat          string      `json:"nosurat"`
	Flaginternal     string      `json:"flaginternal"`
	Kdpoliasal       string      `json:"kdpoliasal"`
	Kdpolituj        string      `json:"kdpolituj"`
	Kdpenunjang      string      `json:"kdpenunjang"`
	Nmpenunjang      string      `json:"nmpenunjang"`
	Diagppk          string      `json:"diagppk"`
	Kddokter         string      `json:"kddokter"`
	Nmdokter         string      `json:"nmdokter"`
	Flagprosedur     string      `json:"flagprosedur"`
	Opsikonsul       string      `json:"opsikonsul"`
	Flagsep          string      `json:"flagsep"`
	Fuser            string      `json:"fuser"`
	Fdate            string      `json:"fdate"`
	Nmdiag           string      `json:"nmdiag"`
}
