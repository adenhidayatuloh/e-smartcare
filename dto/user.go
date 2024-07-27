package dto

type RegisterRequest struct {
	Email     string `json:"email" valid:"required~Email cannot be empty,email~Invalid email format"`
	Password  string `json:"password" valid:"required~Password cannot be empty,minstringlength(6)~password min 6 length"`
	JenisAkun string `json:"jenis_akun" valid:"required~Jenis akun cannot be empty,jenisAkunValidator~Jenis akun hanya siswa admin pakar"`
}

type RegisterResponse struct {
	Email     string `json:"email"`
	JenisAkun string `json:"jenis_akun"`
}

type GetAllUsersResponse struct {
	Email            string `gorm:"primaryKey"`
	JenisAkun        string
	RequestJenisAkun string
}

type LoginRequest struct {
	Email    string `json:"email" valid:"required~Email cannot be empty,email~Invalid email format"`
	Password string `json:"password" valid:"required~Password cannot be ampty,minstringlength(6)~password min 6 length" `
}

type LoginResponse struct {
	Token string `json:"token" binding:"jwt"`
}

// type UpdateUserRequest struct {
// 	Email string `json:"email" valid:"email,required~Email cannot be ampty"`
// }

type UpdateUserResponse struct {
	Message string `json:"Message"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}
