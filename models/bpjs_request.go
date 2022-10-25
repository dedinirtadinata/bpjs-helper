package models

type BPJSRequest struct {
	Request interface{} `json:"request"`
}

type RequestDiagnosa struct {
	Keyword string `json:"keyword" form:"keyword"`
}

type RequestPoli struct {
	Keyword string `json:"keyword" form:"keyword"`
}

type RequestDokter struct {
	Keyword string `json:"keyword" form:"keyword"`
}

type RequestDokterDPJP struct {
	Jenis         string `json:"jenispelayanan" form:"jenispelayanan"`
	Tanggal       string `json:"tgl" form:"tgl"`
	KodeSpesialis string `json:"spesialis" form:"spesialis"`
}

type RequestFaskes struct {
	Keyword string `json:"keyword" form:"keyword"`
	Jenis   string `json:"jenis" form:"jenis"`
}

type RequestKabupaten struct {
	KodePropinsi string `json:"propinsi" form:"propinsi"`
}

type RequestBySEP struct {
	NoSEP string `json:"sep" form:"sep"`
}

type RequestKecamatan struct {
	KodeKabupaten string `json:"kabupaten" form:"kabupaten"`
}

type RequestPeserta struct {
	Tanggal    string `json:"tgl" form:"tgl"`
	JenisKartu string `json:"identitas" form:"identitas"`
	NoKartu    string `json:"noka" form:"nokartu"`
}

type RequestMonitoring struct {
	Tanggal        string `json:"tgl" form:"tgl"`
	JenisPelayanan string `json:"jenispelayanan" form:"jenispelayanan"`
	StatusKlaim    string `json:"statusklaim" form:"statusklaim"`
}

type RequestMonitoringHistoriPelayanan struct {
	TanggalMulai string `json:"tglmulai" form:"tglmulai"`
	TanggalAkhir string `json:"tglakhir" form:"tglakhir"`
	NoKartu      string `json:"nokartu" form:"nokartu"`
}

type RequestMonitoringJasaRaharja struct {
	TanggalMulai   string `json:"tglmulai" form:"tglmulai"`
	TanggalAkhir   string `json:"tglakhir" form:"tglakhir"`
	JenisPelayanan string `json:"jenispelayanan" form:"jenispelayanan"`
}

type RequestDataRujukan struct {
	Faskes  string `json:"faskes" form:"faskes"`
	NoKartu string `json:"norujukan" form:"norujukan"`
	Limit   string `json:"limit" form:"limit"`
}

type RequestRencanaKontrolSpesialistik struct {
	Jenis   string `json:"jenis" form:"jenis"`
	Nomor   string `json:"nomor" form:"nomor"`
	Tanggal string `json:"tgl" form:"tgl"`
}

type RequestRencanaKontrolDokter struct {
	Jenis   string `json:"jenis" form:"jenis"`
	Poli    string `json:"poli" form:"poli"`
	Tanggal string `json:"tgl" form:"tgl"`
}

type RequestRencanaKontrolbyNoSurat struct {
	NoSurat string `json:"nosurat" form:"nosurat"`
}

type RequestRencanaKontrolbyNoKartu struct {
	Bulan   string `json:"bulan" form:"bulan"`
	Tahun   string `json:"tahun" form:"tahun"`
	NoKartu string `json:"nokartu" form:"nokartu"`
	Filter  string `json:"filter" form:"filter"`
}

type RequestListRencanaKontrol struct {
	TanggalAwal  string `json:"tglawal" form:"tglawal"`
	TanggalAkhir string `json:"tglakhir" form:"tglakhir"`
	Filter       string `json:"filter" form:"filter"`
}

type RequestJadwalPraktek struct {
	Poli    string `json:"poli" form:"poli"`
	Tanggal string `json:"tanggal" form:"tanggal"`
}

type RequestUpdateStep struct {
	RegistrasiID int64 `json:"registrasi_id" form:"registrasi_id"`
	Step         int   `json:"step" form:"step"`
}

type RequestSendAntrean struct {
	RegistrasiID   int64 `json:"registrasi_id" form:"registrasi_id"`
	JenisKunjungan int   `json:"jenis_kunjungan" form:"jenis_kunjungan"`
}

type RequestUpdateWaktuLayanan struct {
	Kodebooking string `json:"kodebooking"`
	Taskid      int    `json:"taskid"`
	Waktu       int64  `json:"waktu"`
}

type RequestSyncRuangRawatan struct {
	RuangID int64 `json:"ruang_id" form:"ruang_id"`
}

type RequestDeleteRuangRawatan struct {
	RuangID int64 `json:"ruang_id" form:"ruang_id"`
}
