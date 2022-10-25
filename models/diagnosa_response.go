package models

type DiagnosaResponse struct {
	Metadata struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"metaData"`

	Response struct {
		Diagnosa []*ReferensiModel `json:"diagnosa"`
	} `json:"response"`
}