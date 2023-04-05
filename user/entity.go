package user

import (
	"kredit-api/transaksi"
	"time"
)

type User struct {
	ID           int
	Email        string
	PasswordHash string
	Nik          int
	FullName     string
	LegalName    string
	TempatLahir  string
	TanggalLahir time.Time
	Gaji         int
	FotoKtp      string
	FotoSelfie   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Transaksi    []transaksi.Transaksi
}
