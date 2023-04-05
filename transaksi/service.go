package transaksi

import "github.com/google/uuid"

type Service interface {
	AddTransaksi(input TransaksiInput) (Transaksi, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) AddTransaksi(input TransaksiInput) (Transaksi, error) {
	transaksi := Transaksi{}
	transaksi.UserID = input.UserID
	transaksi.NomorKontrak = uuid.New().String()
	transaksi.Otr = input.Otr
	transaksi.AdminFee = input.AdminFee
	transaksi.JumlahCicilan = input.JumlahCicilan
	transaksi.JumlahBunga = input.JumlahBunga
	transaksi.NamaAsset = input.NamaAsset

	newTransaksi, err := s.repository.Save(transaksi)
	if err != nil {
		return newTransaksi, err
	}

	return newTransaksi, nil
}
