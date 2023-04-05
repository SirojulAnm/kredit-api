package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	FindTransaksiByUserID(ID int) (User, error)
	CreatePhoto(ID int, namefotoKTP string, namefotoSelfie string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(ID int) (User, error) {
	var user User
	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindTransaksiByUserID(ID int) (User, error) {
	var user User
	err := r.db.Preload("Transaksi").Find(&user, ID).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) CreatePhoto(ID int, namefotoKTP string, namefotoSelfie string) (User, error) {
	var user User
	err := r.db.Model(&user).Where("ID = ?", ID).Updates(User{FotoKtp: namefotoKTP, FotoSelfie: namefotoSelfie}).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
