package transaksi

type TransaksiFormatter struct {
	NomorKontrak     string `json:"nomor_kontrak"`
	Otr              int    `json:"otr"`
	AdminFee         int    `json:"admin_fee"`
	JumlahCicilan    int    `json:"jumlah_cicilan"`
	JumlahBunga      int    `json:"jumlah_bunga"`
	NamaAsset        string `json:"nama_asset"`
	TanggalTransaksi string `json:"tanggal_transaksi"`
}

func FormatTransaksi(transaksi Transaksi) TransaksiFormatter {
	formatter := TransaksiFormatter{}
	formatter.NomorKontrak = transaksi.NomorKontrak
	formatter.Otr = transaksi.Otr
	formatter.AdminFee = transaksi.AdminFee
	formatter.JumlahCicilan = transaksi.JumlahCicilan
	formatter.JumlahBunga = transaksi.JumlahBunga
	formatter.NamaAsset = transaksi.NamaAsset
	formatter.TanggalTransaksi = transaksi.CreatedAt.Format("2006-01-02 15:04:05")

	return formatter
}
