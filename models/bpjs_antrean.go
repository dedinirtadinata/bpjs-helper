package models

import "time"

type RequestStatusAntrianPoli struct {
	Kodepoli       string `json:"kodepoli"`
	Kodedokter     int    `json:"kodedokter"`
	Tanggalperiksa string `json:"tanggalperiksa"`
	Jampraktek     string `json:"jampraktek"`
}

type ResponseStatusAntrianPoli struct {
	Namapoli        string `json:"namapoli"`
	Namadokter      string `json:"namadokter"`
	Totalantrean    int    `json:"totalantrean"`
	Sisaantrean     int    `json:"sisaantrean"`
	Antreanpanggil  string `json:"antreanpanggil"`
	Sisakuotajkn    int    `json:"sisakuotajkn"`
	Kuotajkn        int    `json:"kuotajkn"`
	Sisakuotanonjkn int    `json:"sisakuotanonjkn"`
	Kuotanonjkn     int    `json:"kuotanonjkn"`
	Keterangan      string `json:"keterangan"`
}

type RequestGetAntrean struct {
	Nomorkartu     string `json:"nomorkartu"`
	Nik            string `json:"nik"`
	Nohp           string `json:"nohp"`
	Kodepoli       string `json:"kodepoli"`
	Norm           string `json:"norm"`
	Tanggalperiksa string `json:"tanggalperiksa"`
	Kodedokter     int    `json:"kodedokter"`
	Jampraktek     string `json:"jampraktek"`
	Jeniskunjungan int    `json:"jeniskunjungan"`
	Nomorreferensi string `json:"nomorreferensi"`
}

type ResponseGetAntrean struct {
	Nomorantrean     string `json:"nomorantrean"`
	Angkaantrean     int    `json:"angkaantrean"`
	Kodebooking      string `json:"kodebooking"`
	Norm             string `json:"norm"`
	Namapoli         string `json:"namapoli"`
	Namadokter       string `json:"namadokter"`
	Estimasidilayani int64  `json:"estimasidilayani"`
	Sisakuotajkn     int    `json:"sisakuotajkn"`
	Kuotajkn         int    `json:"kuotajkn"`
	Sisakuotanonjkn  int    `json:"sisakuotanonjkn"`
	Kuotanonjkn      int    `json:"kuotanonjkn"`
	Keterangan       string `json:"keterangan"`
}
type RequestSisaAntrean struct {
	Kodebooking string `json:"kodebooking"`
}
type ResponseSisaAntrean struct {
	Nomorantrean   string `json:"nomorantrean"`
	Namapoli       string `json:"namapoli"`
	Namadokter     string `json:"namadokter"`
	Sisaantrean    int    `json:"sisaantrean"`
	Antreanpanggil string `json:"antreanpanggil"`
	Waktutunggu    int    `json:"waktutunggu"`
	Keterangan     string `json:"keterangan"`
}

type RequestBatalAntrean struct {
	Kodebooking string `json:"kodebooking"`
	Keterangan  string `json:"keterangan"`
}

type RequestCheckIn struct {
	Kodebooking string `json:"kodebooking"`
	Waktu       int64  `json:"waktu"`
}

type RequestInfoPasienBaru struct {
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

type ResponsePasienBaru struct {
	Norm string `json:"norm"`
}

type RequestJadwalOperasiRs struct {
	Tanggalawal  string `json:"tanggalawal"`
	Tanggalakhir string `json:"tanggalakhir"`
}

type JadwalOperasi struct {
	Kodebooking    string `json:"kodebooking"`
	Tanggaloperasi string `json:"tanggaloperasi"`
	Jenistindakan  string `json:"jenistindakan"`
	Kodepoli       string `json:"kodepoli"`
	Namapoli       string `json:"namapoli"`
	Terlaksana     int    `json:"terlaksana"`
	Nopeserta      string `json:"nopeserta"`
	Lastupdate     int64  `json:"lastupdate"`
}

type ResponseJadwalOperasiRs struct {
	List []*JadwalOperasi `json:"list"`
}

type RequestJadwalOperasiPasien struct {
	Nopeserta string `json:"nopeserta"`
}

type JadwalOperasiPasien struct {
	Kodebooking    string `json:"kodebooking"`
	Tanggaloperasi string `json:"tanggaloperasi"`
	Jenistindakan  string `json:"jenistindakan"`
	Kodepoli       string `json:"kodepoli"`
	Namapoli       string `json:"namapoli"`
	Terlaksana     int    `json:"terlaksana"`
}

type ResponseJadwalOperasiPasien struct {
	List []*JadwalOperasiPasien `json:"list"`
}

type ResponseToken struct {
	Token string `json:"token"`
}

type StatusAntreanService struct {
	Status       string `json:"status"`
	Message      string `json:"message"`
	AntrianNomor string `json:"antrian_nomor"`
	ServiceName  string `json:"service_name"`
	LeftQueue    int    `json:"left_queue"`
	TotalQueue   int    `json:"total"`
}

type BPJSAntreanModel struct {
	Kodebooking      string `gorm:"column:kodebooking";json:"kodebooking"`
	Jenispasien      string `gorm:"column:jenispasien";json:"jenispasien"`
	Nomorkartu       string `gorm:"column:nomorkartu";json:"nomorkartu"`
	Nik              string `json:"nik"`
	Nohp             string `json:"nohp"`
	Kodepoli         string `gorm:"column:kodepoli";json:"kodepoli"`
	Namapoli         string `gorm:"column:namapoli";json:"namapoli"`
	Pasienbaru       int    `json:"pasienbaru"`
	Norm             string `gorm:"column:norm";json:"norm"`
	Tanggalperiksa   string `json:"tanggalperiksa"`
	Kodedokter       int    `gorm:"column:kodedokter";json:"kodedokter"`
	Namadokter       string `gorm:"column:namadokter";json:"namadokter"`
	Jampraktek       string `json:"jampraktek"`
	Jeniskunjungan   int    `json:"jeniskunjungan"`
	Nomorreferensi   string `gorm:"column:nomorreferensi";json:"nomorreferensi"`
	Nomorantrean     string `gorm:"column:nomorantrean";json:"nomorantrean"`
	Angkaantrean     int    `json:"angkaantrean"`
	Estimasidilayani int64  `json:"estimasidilayani"`
	Sisakuotajkn     int    `json:"sisakuotajkn"`
	Kuotajkn         int    `json:"kuotajkn"`
	Sisakuotanonjkn  int    `json:"sisakuotanonjkn"`
	Kuotanonjkn      int    `json:"kuotanonjkn"`
	Keterangan       string `json:"keterangan"`
}

type BPJSAntreanModel2 struct {
	Kodebooking      string `json:"kodebooking"`
	Jenispasien      string `json:"jenispasien"`
	Nomorkartu       string `json:"nomorkartu"`
	Nik              string `json:"nik"`
	Nohp             string `json:"nohp"`
	Kodepoli         string `json:"kodepoli"`
	Namapoli         string `json:"namapoli"`
	Pasienbaru       int    `json:"pasienbaru"`
	Norm             string `json:"norm"`
	Tanggalperiksa   string `json:"tanggalperiksa"`
	Kodedokter       int    `json:"kodedokter"`
	Namadokter       string `json:"namadokter"`
	Jampraktek       string `json:"jampraktek"`
	Jeniskunjungan   int    `json:"jeniskunjungan"`
	Nomorreferensi   string `json:"nomorreferensi"`
	Nomorantrean     string `json:"nomorantrean"`
	Angkaantrean     int    `json:"angkaantrean"`
	Estimasidilayani int64  `json:"estimasidilayani"`
	Sisakuotajkn     int    `json:"sisakuotajkn"`
	Kuotajkn         int    `json:"kuotajkn"`
	Sisakuotanonjkn  int    `json:"sisakuotanonjkn"`
	Kuotanonjkn      int    `json:"kuotanonjkn"`
	Keterangan       string `json:"keterangan"`
}

// "kodebooking": "{kodebooking yang didapat dari servis tambah antrean}",
//   "taskid": {
//                1 (mulai waktu tunggu admisi),
//                2 (akhir waktu tunggu admisi/mulai waktu layan admisi),
//                3 (akhir waktu layan admisi/mulai waktu tunggu poli),
//                4 (akhir waktu tunggu poli/mulai waktu layan poli),
//                5 (akhir waktu layan poli/mulai waktu tunggu farmasi),
//                6 (akhir waktu tunggu farmasi/mulai waktu layan farmasi membuat obat),
//                7 (akhir waktu obat selesai dibuat),
//                99 (tidak hadir/batal)
//            },
//   "waktu": {waktu dalam timestamp milisecond}
//  /*

type MappingRegistrasiAntrean struct {
	ID           int    `gorm:"primary_key;column:ma_id";json:"id"`
	RegistrasiId int64  `gorm:"column:ma_registrationid";json:"registrasi_id"`
	KodeBooking  string `gorm:"column:ma_kodebooking";json:"kode_booking"`
	Step         int    `gorm:"column:ma_step";json:"step"`
	Waktu        int64  `gorm:"column:ma_waktu";json:"waktu"`
}

func (m MappingRegistrasiAntrean) LastAction() time.Time {
	return time.Unix(0, m.Waktu*int64(time.Millisecond))
}

func (MappingRegistrasiAntrean) TableName() string {
	return "mapping_registrasi_antrean"
}
