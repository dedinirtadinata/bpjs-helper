package models

type ListKunjungan struct {
	List []KunjunganModel `json:"sep"`
}

type ListKlaim struct {
	List []BPJSKlaimModel `json:"klaim"`
}

type ListPelayanan struct {
	List []BPJSPelayananModel `json:"histori"`
}

type ListKlaimJasaRaharja struct {
	List []JasaRaharjaKlaimModel `json:"jaminan"`
}
