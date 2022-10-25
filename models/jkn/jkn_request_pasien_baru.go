package jkn

type RequestRegisterPasienBaru struct {
	Nomorkartu   string `json:"nomorkartu"`
	Nik          string `json:"nik"`
	Nomorkk      string `json:"nomorkk"`
	Nama         string `json:"nama"`
	Jeniskelamin string `json:"jeniskelamin"`
	Tanggallahir string `json:"tanggallahir"`
	Nohp         string `json:"nohp"`
	Alamat       string `json:"alamat"`
	Kodeprop     string `json:"kodeprop"`
	Namaprop     string `json:"namaprop"`
	Kodedati2    string `json:"kodedati2"`
	Namadati2    string `json:"namadati2"`
	Kodekec      string `json:"kodekec"`
	Namakec      string `json:"namakec"`
	Kodekel      string `json:"kodekel"`
	Namakel      string `json:"namakel"`
	Rw           string `json:"rw"`
	Rt           string `json:"rt"`
}
