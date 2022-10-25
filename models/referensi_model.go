package models

type ReferensiModel struct {
	Kode string `json:"kode"`
	Nama string `json:"nama"`
}

type ListPropinsi struct {
	List []ReferensiModel `json:"list"`
}

type ListKabupaten struct {
	List []ReferensiModel `json:"list"`
}

type ListKecamatan struct {
	List []ReferensiModel `json:"list"`
}

type ListPascaPulang struct {
	List []ReferensiModel `json:"list"`
}

type ListCaraKeluar struct {
	List []ReferensiModel `json:"list"`
}

type ListRuangRawat struct {
	List []ReferensiModel `json:"list"`
}

type ListSpesialistik struct {
	List []ReferensiModel `json:"list"`
}

type ListKelasRawat struct {
	List []ReferensiModel `json:"list"`
}

type ListDiagnosaPRB struct {
	List []ReferensiModel `json:"list"`
}

type ListDiagnosa struct {
	List []ReferensiModel `json:"diagnosa"`
}

type ListPoli struct {
	List []ReferensiModel `json:"poli"`
}

type ListFaskes struct {
	List []ReferensiModel `json:"faskes"`
}

type ListProcedure struct {
	List []ReferensiModel `json:"procedure"`
}

type ListDokter struct {
	List []ReferensiModel `json:"list"`
}

type ListDokterDPJP struct {
	List []ReferensiModel `json:"list"`
}
