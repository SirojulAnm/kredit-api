package transaksi

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(transaksi Transaksi) (Transaksi, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(transaksi Transaksi) (Transaksi, error) {
	err := r.db.Create(&transaksi).Error
	if err != nil {
		return transaksi, err
	}

	return transaksi, nil
}
