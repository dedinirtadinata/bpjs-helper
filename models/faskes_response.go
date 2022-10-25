package models

type FaskesResponse struct {
	Metadata struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"metaData"`
	Response struct {
		Faskes []struct {
			Kode string `json:"kode"`
			Nama string `json:"nama"`
		} `json:"faskes"`
	} `json:"response"`
}