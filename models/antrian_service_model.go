package models

type AntrianServiceModel struct {
	ID   int    `db:"layanan_id" json:"id"`
	Name string `db:"layanan_nama" json:"nama"`
}

func (u *AntrianServiceModel) TableName() string {
	return "tbl_layanan"
}
