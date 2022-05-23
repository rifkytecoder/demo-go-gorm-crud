package crud

import (
	"lab-gorm-crud/entity"

	"gorm.io/gorm"
)

// Create data
func InsertProduct(db *gorm.DB, product *entity.Products) {
	result := db.Create(product)

	if result.Error != nil {
		panic("Gagal melakukan proses insert data produk")
	}
}

// Show all data
func ShowAllProducts(db *gorm.DB) []entity.Products {
	var produk []entity.Products

	//result := db.Find(&produk)
	result := db.Where("kode_produk = ?", "001").Find(&produk) //filter

	if result.Error != nil {
		panic("Gagal menampilkan data produk")
	}
	return produk
}

func GetOneRow(db *gorm.DB, kode_produk string) entity.Products {
	var product entity.Products

	db.Where("kode_produk = ?", kode_produk).First(&product)

	return product
}

func UpdateProduct(db *gorm.DB, kode_produk string, data_produk entity.Products) {
	result := db.Where("kode_produk = ?", kode_produk).Updates(&data_produk)

	if result.Error != nil {
		panic("Gagal melakukan update data produk")
	}
}

func DeleteProduct(db *gorm.DB, kode_produk string) {
	result := db.Delete(&entity.Products{
		KodeProduk: kode_produk,
	})

	if result.Error != nil {
		panic("Gagal melakukan delete produk")
	}
}
