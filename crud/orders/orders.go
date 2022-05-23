package crud

import (
	"lab-gorm-crud/entity"

	"gorm.io/gorm"
)

func InsertOrder(db *gorm.DB, order *entity.Orders) {
	result := db.Create(order)

	if result.Error != nil {
		panic("Gagal insert data order")
	}
}

func InsertOrderDetails(db *gorm.DB, details *[]entity.OrderDetails) {
	result := db.Create(details)

	if result.Error != nil {
		panic("Gagal insert data orders")
	}
}

func ShowAllOrders(db *gorm.DB) []entity.OrderDetails {
	var orders []entity.OrderDetails

	//result := db.Joins("Products").Joins("Orders").Find(&orders)
	result := db.Joins("join products as pd ON pd.kode_produk = order_details.kode_produk AND pd.stok > 90").Preload("Products").
		Joins("Orders").Find(&orders)

	if result.Error != nil {
		panic("Gagal menampilkan data order detail")
	}

	return orders
}

func GetOneRow(db *gorm.DB, id_order int) entity.Orders {
	var order entity.Orders

	db.Where("id_order = ?", id_order).First(&order)

	return order
}

func UpdateOrder(db *gorm.DB, id_order int, data_order entity.Orders) {
	//result := db.Where("id_order = ?", id_order).Updates(data_order)
	result := db.Where(&entity.Orders{
		Id_Order: id_order,
		// PaymentMethod: "Tunai",
		// Total:         40000,
	}).Updates(&data_order)

	if result.Error != nil {
		panic("Gagal melakukan update data order")
	}
}

func DeleteOrderDetail(db *gorm.DB, id_order_detail int) {
	result := db.Where("id_order_detail = ?", id_order_detail).Delete(&entity.OrderDetails{})

	if result.Error != nil {
		panic("Gagal menghapus order detail")
	}
}
