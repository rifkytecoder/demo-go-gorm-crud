package crud

import (
	"fmt"
	"lab-gorm-crud/config"
	"lab-gorm-crud/entity"
	"strconv"
	"testing"
)

func TestShowOrders(t *testing.T) {
	config.ConnectDB()

	orders := ShowAllOrders(config.DB)

	for _, val := range orders {
		fmt.Println(val)
		//fmt.Println(val.Products)
	}
}

func TestUpdateOrder(t *testing.T) {
	config.ConnectDB()

	id_order := 1
	data_order := entity.Orders{
		PaymentMethod: "Credit Card",
		Total:         50000,
	}
	UpdateOrder(config.DB, id_order, data_order)

	order := GetOneRow(config.DB, id_order)

	fmt.Println("Data order baru")
	fmt.Println(order)

	var (
		payment_method = "Cash"
		total          = 30000
	)

	if order.PaymentMethod == payment_method && order.Total == total {
		t.Errorf("Gagal melakukan update data order" + strconv.Itoa(id_order))
	}

}

func TestDeleteOrderDetails(t *testing.T) {
	config.ConnectDB()
	DeleteOrderDetail(config.DB, 2)
}
