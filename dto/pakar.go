package dto

type CreatePakarRequest struct {
	NamaLengkap string `json:"nama_lengkap"`
	Alamat      string `json:"alamat"`
	NoTelepon   string `json:"no_telepon"`
	FotoProfil  string `json:"foto_profil"`
}

type UpdatePhotoPakarResponse struct {
	Email      string `json:"email"`
	FotoProfil string `json:"foto_profil"`
}

type CreatePakarResponse struct {
	Email       string `json:"email"`
	NamaLengkap string `json:"nama_lengkap"`
	Alamat      string `json:"alamat"`
	NoTelepon   string `json:"no_telepon"`
	FotoProfil  string `json:"foto_profil"`
}
