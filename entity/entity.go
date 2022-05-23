package entity

import "time"

type Products struct {
	KodeProduk string `gorm:"column:kode_produk;primaryKey"`
	NamaProduk string `gorm:"column:nama_produk"`
	Stok       int    `gorm:"column:stok"`
	Harga      int    `gorm:"column:harga"`
}

type Orders struct {
	Id_Order      int       `gorm:"column:id_order;primaryKey;autoIncrement"`
	TanggalOrder  time.Time `gorm:"column:tanggal_order"`
	PaymentMethod string    `gorm:"column:payment_method"`
	Total         int       `gorm:"column:total"`
}

type OrderDetails struct {
	Id_OrderDetails int      `gorm:"column:id_order_detail;primaryKey;autoIncrement"`
	FK_IdOrder      int      `gorm:"column:id_order"`
	Orders          Orders   `gorm:"foreignKey:FK_IdOrder"`
	FK_KodeProduk   string   `gorm:"column:kode_produk"`
	Products        Products `gorm:"foreignKey:FK_KodeProduk"`
	Qty             int      `gorm:"column:qty"`
}
