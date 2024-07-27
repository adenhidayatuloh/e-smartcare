package dto

type CreateAdminRequest struct {
	NamaLengkap string `json:"nama_lengkap"`
	Alamat      string `json:"alamat"`
	NoTelepon   string `json:"no_telepon"`
	FotoProfil  string `json:"foto_profil"`
}

type UpdatePhotoAdminResponse struct {
	Email      string `json:"email"`
	FotoProfil string `json:"foto_profil"`
}

type CreateAdminResponse struct {
	Email       string `json:"email"`
	NamaLengkap string `json:"nama_lengkap"`
	Alamat      string `json:"alamat"`
	NoTelepon   string `json:"no_telepon"`
	FotoProfil  string `json:"foto_profil"`
}
