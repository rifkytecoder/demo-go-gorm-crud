package crud

import (
	"fmt"
	"lab-gorm-crud/config"
	"lab-gorm-crud/entity"
	"testing"
)

func TestShowProducts(t *testing.T) {
	config.ConnectDB()

	produk := ShowAllProducts(config.DB)

	for _, val := range produk {
		fmt.Println(val)
	}
}

func TestShowOneRow(t *testing.T) {
	config.ConnectDB()
	kode_produk := "001"
	produk := GetOneRow(config.DB, kode_produk)

	if produk.KodeProduk == "" {
		t.Errorf("Data dengan kode " + kode_produk + " kosong")
	}

	fmt.Println(produk)

}

func TestUpdateProduk(t *testing.T) {
	config.ConnectDB()

	kode_produk := "001"
	data_produk_update := entity.Products{
		NamaProduk: "Indomie Goreng",
		Stok:       500,
		Harga:      3500,
	}

	UpdateProduct(config.DB, kode_produk, data_produk_update)

	produk := GetOneRow(config.DB, kode_produk)
	fmt.Println("Data produk baru")
	fmt.Println(produk)
	var (
		nama_produk = "Indomie"
		stok        = 100
		harga       = 2500
	)

	if produk.NamaProduk == nama_produk && produk.Stok == stok && produk.Harga == harga {
		t.Errorf("Gagal melakukan update data produk" + kode_produk)
	}
}

func TestDeleteProduct(t *testing.T) {
	config.ConnectDB()

	DeleteProduct(config.DB, "002")
}
