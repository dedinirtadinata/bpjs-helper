package simrs

type UserApi struct {
	ID       int64  `gorm:"primary_key;column:au_id"`
	Username string `gorm:"column:au_username"`
	Password string `gorm:"column:au_password"`
} // default table name is `users`

// Set User's table name to be `profiles`
func (UserApi) TableName() string {
	return "user_api"
}

type JadwalOperasi struct {
	ID             int64  `gorm:"primary_key;column:id_jadwal_operasi"`
	KodeBooking    string `gorm:"column:kode_booking"`
	TanggalOperasi string `gorm:"column:tanggal"`
	NoBPJS         string `gorm:"column:no_bpjs"`
	KodePoli       string `gorm:"column:kode_polibpjs"`
	NamaPoli       string `gorm:"column:nama_polibpjs"`
	JenisTindakan  string `gorm:"column:jenis_tindakan"`
	Status         string `gorm:"column:status"`
	UpdateTime     string `gorm:"column:updateddate"`
}

func (JadwalOperasi) TableName() string {
	return "jadwal_operasi"
}

type PoliBPJS struct {
	IDPoli   int64  `gorm:"column:kd_rs"`
	KodeBPJS string `gorm:"primary_key;column:kd_bpjs"`
	Nama     string `gorm:"column:nm_bpjs"`
}

func (PoliBPJS) TableName() string {
	return "maping_poli_bpjs"
}

type PoliModel struct {
	ID              int64  `gorm:"column:id"`
	NamaPoli        string `gorm:"column:nama"`
	JenisDepartemen string `gorm:"column:jenis"`
	KodeBPJS        string `gorm:"column:kode_bpjs"`
	NamaBPJS        string `gorm:"column:nama_bpjs"`
}

type PasienModel struct {
	ID           int    `gorm:"column:id";json:"id"`
	NIK          string `gorm:"column:nik";json:"nik"`
	Asuransi     string `gorm:"column:asuransi";json:"asuransi"`
	Nokartu      string `gorm:"column:nokartu";json:"nokartu"`
	Kelas        string `gorm:"column:kelas";json:"kelas"`
	NoRM         string `gorm:"column:norm";json:"norm"`
	Nama         string `gorm:"column:nama";json:"nama"`
	JenisKelamin string `gorm:"column:jenis_kelamin";json:"jenis_kelamin"`
	TanggalLahir string `gorm:"column:tanggal_lahir";json:"tanggal_lahir"`
	Alamat       string `gorm:"column:alamat";json:"alamat"`
	NoHP         string `gorm:"column:nohp";json:"nohp"`
}

type ReservasiOnlineModel struct {
	ID               int    `gorm:"primary_key;column:ba_id";json:"id"`
	IDSchedulestring int    `gorm:"column:ba_schedule";json:"schedule_id"`
	Nomorkartu       string `gorm:"column:ba_nomorkartu";json:"nomorkartu"`
	Nik              string `gorm:"column:ba_nik";json:"nik"`
	Nohp             string `gorm:"column:ba_nohp";json:"nohp"`
	Kodepoli         string `gorm:"column:ba_kodepoli";json:"kodepoli"`
	Norm             string `gorm:"column:ba_norm";json:"norm"`
	Tanggalperiksa   string `gorm:"column:ba_tglperiksa";json:"tanggalperiksa"`
	Kodedokter       int    `gorm:"column:ba_kodedokter";json:"kodedokter"`
	Jampraktek       string `gorm:"column:ba_jampraktek";json:"jampraktek"`
	Jeniskunjungan   int    `gorm:"column:ba_jeniskunjungan";json:"jeniskunjungan"`
	Nomorreferensi   string `gorm:"column:ba_noreferensi";json:"nomorreferensi"`
	KodeBooking      string `gorm:"column:ba_kodebooking";json:"nomorreferensi"`
	NoAntrean        string `gorm:"column:ba_nomorantrean";json:"noantrean"`
	Deleted          int    `gorm:"column:ba_deleted";json:"ba_deleted"`
	ISCheckin        int    `gorm:"column:ba_checkin";json:"checkin"`
	WaktuCheckin     int64  `gorm:"column:ba_waktucheckin";json:"waktu_checkin"`
	Keterangan       string `gorm:"column:ba_keterangan";json:"keterangan"`
}

func (ReservasiOnlineModel) TableName() string {
	return "bpjs_booking"
}

type ScheduleDokter struct {
	SsId          int    `gorm:"primary_key;column:ss_id";json:"ss_id"`
	SsDokter      int    `gorm:"column:ss_dokter";json:"ss_dokter"`
	SsName        string `gorm:"column:ss_name";json:"ss_name"`
	SsSpesialis   string `gorm:"column:ss_spesialis";json:"ss_spesialis"`
	SsAntrean     int    `gorm:"column:ss_antrean";json:"ss_antrean"`
	SsSlot        int    `gorm:"column:ss_slot";json:"ss_slot"`
	SsDate        string `gorm:"column:ss_date";json:"ss_date"`
	SsStarttime   string `gorm:"column:ss_starttime";json:"ss_starttime"`
	SsEndtime     string `gorm:"column:ss_endtime";json:"ss_endtime"`
	SsDeleted     string `gorm:"column:ss_deleted";json:"ss_deleted"`
	SsQuotanonjkn int    `gorm:"column:ss_quotanonjkn";json:"ss_quotanonjkn"`
	SsQuotajkn    int    `gorm:"column:ss_quotajkn";json:"ss_quotajkn"`
	Sisanonjkn    int    `gorm:"column:ss_sisanonjkn";json:"ss_sisanonjkn"`
	Sisajkn       int    `gorm:"column:ss_sisajkn";json:"ss_sisajkn"`
}

func (ScheduleDokter) TableName() string {
	return "schedule_shift"
}

type MappingAntreanModel struct {
	Title        string `json:"title"`
	Poli         int    `json:"poli"`
	Antrean      int    `json:"antrean"`
	Slot         int    `json:"slot"`
	WaktuPraktek string `json:"waktupraktek"`
}

type DokterBPJS struct {
	ID       int    `gorm:"primary_key;column:kd_dokter";json:"kd_dokter"`
	KodeBPJS int    `gorm:"column:kd_dokter_bpjs";json:"kd_dokter_bpjs"`
	Nama     string `gorm:"column:nm_dokter_bpjs";json:"nm_dokter_bpjs"`
}

func (DokterBPJS) TableName() string {
	return "maping_dokter_dpjp_bpjs"
}

type JadwalPraktekDokterModel struct {
	ID           int    `gorm:"primary_key;column:ss_id";json:"id"`
	Dokter       int    `gorm:"column:ss_dokter";json:"dokter"`
	NamaDokter   string `gorm:"column:ss_name";json:"nama_dokter"`
	Spesialis    string `gorm:"column:ss_spesialis";json:"spesialis"`
	Antrean      int    `gorm:"column:ss_antrean";json:"antrean"`
	Slot         int    `gorm:"column:ss_slot";json:"slot"`
	Tanggal      string `gorm:"column:ss_date";json:"tanggal"`
	WaktuMulai   string `gorm:"column:ss_starttime";json:"waktu_mulai"`
	WaktuSelesai string `gorm:"column:ss_endtime";json:"waktu_selesai"`
	QuotaNonJKN  int    `gorm:"column:ss_quotanonjkn";json:"kuota_nonjkn"`
	QuotaJKN     int    `gorm:"column:ss_quotajkn";json:"kuota_jkn"`
	SisaNonJKN   int    `gorm:"column:ss_sisanonjkn";json:"sisa_nonjkn"`
	SisaJKN      int    `gorm:"column:ss_sisajkn";json:"sisa_jkn"`

	Source string `gorm:"column:ss_source";json:"source"`
}

func (JadwalPraktekDokterModel) TableName() string {
	return "schedule_shift"
}
