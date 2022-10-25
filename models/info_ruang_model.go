package models

type InfoRuangModel struct {
	ID         int    `db:"id_kelas" json:"id"`
	Nama       string `db:"nama_kelas" json:"nama"`
	Jenis      string `db:"jenis_kelas" json:"jenis"`
	TotalKamar string `db:"total_kamar" json:"total_kamar"`
	Tersedia   string `db:"tersedia" json:"tersedia"`
}

func (u *InfoRuangModel) TableName() string {
	return "kelas"
}
