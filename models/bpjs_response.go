package models

type BPJSResponse struct {
	Metadata struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	} `json:"metadata"`

	Response string `json:"response"`
}

type BPJSResponse2 struct {
	Metadata struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"metadata"`

	Response string `json:"response"`
}

type BPJSResponse3 struct {
	Metadata struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"metadata"`

	Response interface{} `json:"response"`
}
