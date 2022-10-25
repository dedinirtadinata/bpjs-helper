package models

type DokterResponse struct {
	Metadata struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"metaData"`
	Response struct {
		List []struct {
			Kode string `json:"kode"`
			Nama string `json:"nama"`
		} `json:"list"`
	} `json:"response"`
}