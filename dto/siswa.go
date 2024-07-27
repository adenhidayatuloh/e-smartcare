package dto

type CreateSiswaRequest struct {
	NIS          string `json:"nis"`
	NamaLengkap  string `json:"nama_lengkap"`
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"`
	Alamat       string `json:"alamat"`
	NoTelepon    string `json:"no_telepon"`
	Kelas        string `json:"kelas"`
	Agama        string `json:"agama"`
	FotoProfil   string `json:"foto_profil"`
}

type UpdatePhotoResponse struct {
	Email      string `json:"email"`
	FotoProfil string `json:"foto_profil"`
}

type CreateSiswaResponse struct {
	Email        string `json:"email"`
	NIS          string `json:"nis"`
	NamaLengkap  string `json:"nama_lengkap"`
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"`
	Alamat       string `json:"alamat"`
	NoTelepon    string `json:"no_telepon"`
	Kelas        string `json:"kelas"`
	Agama        string `json:"agama"`
	FotoProfil   string `json:"foto_profil"`
}
