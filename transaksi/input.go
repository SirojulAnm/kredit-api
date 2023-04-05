package transaksi

type TransaksiInput struct {
	UserID        int    `json:"user_id"`
	Otr           int    `json:"otr" binding:"required"`
	AdminFee      int    `json:"admin_fee" binding:"required"`
	JumlahCicilan int    `json:"jumlah_cicilan" binding:"required"`
	JumlahBunga   int    `json:"jumlah_bunga" binding:"required"`
	NamaAsset     string `json:"nama_asset" binding:"required"`
}
