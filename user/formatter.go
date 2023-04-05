package user

import "kredit-api/transaksi"

type LoginFormatter struct {
	Nik       int    `json:"nik"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	LegalName string `json:"LegalName"`
}

type UserFormatter struct {
	Email        string                         `json:"email"`
	Nik          int                            `json:"nik"`
	FullName     string                         `json:"full_name"`
	LegalName    string                         `json:"legal_name"`
	TempatLahir  string                         `json:"tempat_lahir"`
	TanggalLahir string                         `json:"tanggal_lahir"`
	Gaji         int                            `json:"gaji"`
	FotoKtp      string                         `json:"foto_ktp"`
	FotoSelfie   string                         `json:"foto_selfie"`
	Transaksi    []transaksi.TransaksiFormatter `json:"history_transaksi"`
}

func FormatLogin(user User) LoginFormatter {
	formatter := LoginFormatter{}
	formatter.Nik = user.Nik
	formatter.Email = user.Email
	formatter.FullName = user.FullName
	formatter.LegalName = user.LegalName

	return formatter
}

func FormatUser(user User, baseURL string) UserFormatter {
	formatter := UserFormatter{}
	formatter.Nik = user.Nik
	formatter.Email = user.Email
	formatter.FullName = user.FullName
	formatter.LegalName = user.LegalName
	formatter.TempatLahir = user.TempatLahir
	formatter.Gaji = user.Gaji
	formatter.FotoKtp = baseURL + "/" + user.FotoKtp
	formatter.FotoSelfie = baseURL + "/" + user.FotoSelfie

	return formatter
}

func FormatHistoryTransaksi(user User, baseURL string) UserFormatter {
	formatter := UserFormatter{}
	formatter.Nik = user.Nik
	formatter.Email = user.Email
	formatter.FullName = user.FullName
	formatter.LegalName = user.LegalName
	formatter.TempatLahir = user.TempatLahir
	formatter.Gaji = user.Gaji
	formatter.FotoKtp = baseURL + "/" + user.FotoKtp
	formatter.FotoSelfie = baseURL + "/" + user.FotoSelfie

	var transaksiFormatter []transaksi.TransaksiFormatter
	if len(user.Transaksi) > 0 {
		for _, item := range user.Transaksi {
			formatterTransaksi := transaksi.FormatTransaksi(item)
			transaksiFormatter = append(transaksiFormatter, formatterTransaksi)
		}
	}

	formatter.Transaksi = transaksiFormatter

	return formatter
}
