package dto

type GetAdminPakarResponse struct {
	Email            string `json:"email"`
	NamaLengkap      string `json:"nama_lengkap"`
	Alamat           string `json:"alamat"`
	NoTelepon        string `json:"no_telepon"`
	FotoProfil       string `json:"foto_profil"`
	JenisAkun        string `json:"JenisAkun"`
	RequestJenisAkun string `json:"RequestJenisAkun"`
}

type GetSiswaResponse struct {
	Email            string `json:"email"`
	NIS              string `json:"nis"`
	NamaLengkap      string `json:"nama_lengkap"`
	TempatLahir      string `json:"tempat_lahir"`
	TanggalLahir     string `json:"tanggal_lahir"`
	Alamat           string `json:"alamat"`
	NoTelepon        string `json:"no_telepon"`
	Kelas            string `json:"kelas"`
	Agama            string `json:"agama"`
	FotoProfil       string `json:"foto_profil"`
	JenisAkun        string `json:"JenisAkun"`
	RequestJenisAkun string `json:"RequestJenisAkun"`
}
