package models

type JasaRaharjaKlaimModel struct {
	Sep struct {
		NoSEP        string `json:"noSEP"`
		TglSEP       string `json:"tglSEP"`
		TglPlgSEP    string `json:"tglPlgSEP"`
		NoMr         string `json:"noMr"`
		JnsPelayanan string `json:"jnsPelayanan"`
		Poli         string `json:"poli"`
		Diagnosa     string `json:"diagnosa"`
		Peserta      struct {
			NoKartu string `json:"noKartu"`
			Nama    string `json:"nama"`
			NoMR    string `json:"noMR"`
		} `json:"peserta"`
	} `json:"sep"`
	JasaRaharja struct {
		TglKejadian        string `json:"tglKejadian"`
		NoRegister         string `json:"noRegister"`
		KetStatusDijamin   string `json:"ketStatusDijamin"`
		KetStatusDikirim   string `json:"ketStatusDikirim"`
		BiayaDijamin       string `json:"biayaDijamin"`
		Plafon             string `json:"plafon"`
		JmlDibayar         string `json:"jmlDibayar"`
		ResultsJasaRaharja string `json:"resultsJasaRaharja"`
	} `json:"jasaRaharja"`
}
