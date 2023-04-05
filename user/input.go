package user

type LoginAdminRequest struct {
	Email         string `json:"email" binding:"required"`
	Password      string `json:"password" binding:"required"`
	FirebaseToken string `json:"firebase_token"`
}

type InputRegister struct {
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Nik          int    `json:"nik" binding:"required"`
	FullName     string `json:"full_name" binding:"required"`
	LegalName    string `json:"legal_name" binding:"required"`
	TempatLahir  string `json:"tempat_lahir" binding:"required"`
	TanggalLahir string `json:"tanggal_lahir" binding:"required"`
	Gaji         int    `json:"gaji" binding:"required"`
}
