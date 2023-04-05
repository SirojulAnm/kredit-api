package transaksi

import "time"

type Transaksi struct {
	ID            int
	UserID        int
	NomorKontrak  string
	Otr           int
	AdminFee      int
	JumlahCicilan int
	JumlahBunga   int
	NamaAsset     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Tabler interface {
	TableName() string
}

func (Transaksi) TableName() string {
	return "transaksi"
}
