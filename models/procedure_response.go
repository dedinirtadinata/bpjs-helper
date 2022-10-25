package models

type ProcedureResponse struct {
	Metadata struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"metaData"`
	Response struct {
		Procedure []struct {
			Kode string `json:"kode"`
			Nama string `json:"nama"`
		} `json:"procedure"`
	} `json:"response"`
}