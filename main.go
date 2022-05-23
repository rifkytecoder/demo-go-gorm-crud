package main

import (
	"lab-gorm-crud/config"
	crud_order "lab-gorm-crud/crud/orders"
	crud_product "lab-gorm-crud/crud/products"
	"lab-gorm-crud/entity"
	"time"
)

func main() {

	// Initialize Database
	config.ConnectDB()

	// Insert data Product
	product := entity.Products{
		KodeProduk: "001",
		NamaProduk: "Indomie",
		Stok:       100,
		Harga:      2500,
	}

	crud_product.InsertProduct(config.DB, &product)

	product = entity.Products{
		KodeProduk: "002",
		NamaProduk: "Sabun Mandi",
		Stok:       200,
		Harga:      1000,
	}

	crud_product.InsertProduct(config.DB, &product)

	// Insert data Order
	order := entity.Orders{
		TanggalOrder:  time.Now(),
		PaymentMethod: "Cash",
		Total:         30000,
	}

	crud_order.InsertOrder(config.DB, &order)

	details := []entity.OrderDetails{
		{
			Orders: order,
			Products: entity.Products{
				KodeProduk: "001",
			},
			Qty: 10,
		},
		{
			Orders: order,
			Products: entity.Products{
				KodeProduk: "002",
			},
			Qty: 5,
		},
	}

	crud_order.InsertOrderDetails(config.DB, &details)
}
