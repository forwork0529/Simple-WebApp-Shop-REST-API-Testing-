package memDB

import (
	"reflect"
	"testing"
	"time"
)



func Test_CRUD_memDB(t *testing.T) {

	db := New()
	amount := 10

	car1 := Product{ID : 101, Name: "model1", Price: 100000}
	car2 := Product{ID : 102, Name: "model2", Price: 200000}
	testOrder := Order {amount, true, 100, "Manhattan", []Product{car1, car2}}

	t.Run("NewOrder(Create)", func(t * testing.T){
		for i := 1; i <= amount ; i ++{
			db.NewOrder(testOrder)
		}
		if db.store[1].ID != 1 || db.store[amount].ID != amount {
			t.Errorf("NewOrder(Create) function auto-increment errogo r ")
		}
		if len(db.store) < amount{
			t.Errorf("NewOrder(Create) function dont created all input orders")
		}
		if !(reflect.DeepEqual(testOrder, db.store[amount])){
			t.Errorf("got: %v, не равно want: %v", db.store[amount], testOrder)
		}
	})

	t.Run("Orders(Read)", func(t *testing.T){

		orders := db.Orders()
		if len(db.store) < 2{
			t.Errorf("Orders(Read) dont have anough data for testing")
		}
		if len (orders) < len(db.store){
			t.Errorf( "Orders(Read) function dont returned  all stored orders")
		}
		var found bool
		for _, order := range orders{
			if order.ID == 1{
				found = true
			}
		}
		if ! found{
			t.Errorf( "Orders(Read) wrong data read")
		}
	})

	t.Run("UpdateOrder(Update)", func(t *testing.T){
		lenS := len(db.store)
		if lenS < 2{
			t.Errorf("UpdateOrder(Update) dont have anough data for testing")
		}
		db.UpdateOrder(Order {1, true, 100, "Manhattan", []Product{car1}})
		if len(db.store) != lenS {
			t.Errorf("UpdateOrder(Update) update operation create/delete order in storage")
		}
		if len(db.store[1].Products) != 1{
			t.Errorf("UpdateOrder(Update) changes by function dont saved in memory")
		}
	})

	t.Run("DeleteOrder(Delete)", func(t *testing.T){
		lenS := len(db.store)
		if lenS < 2{
			t.Errorf("DeleteOrder(Delete) dont have anough data for testing")
		}

		db.DeleteOrder(1)
		if _, ok := db.store[1]; ok{
			t.Errorf("DeleteOrder(Delete) data dont deleted")
		}
		if lenS  <= len(db.store){
			t.Errorf("DeleteOrder(Delete) data dont deleted")
		}

	})

}

func Test_CRUD_memDB2(t *testing.T) {
	// Создаем БД.
	db := New()
	o := Order{
		IsOpen:       true,
		DeliveryTime: time.Now().Unix(),
	}

	// Тест создания записи в БД.
	o.ID = db.NewOrder(o)
	// Проверка.
	ord := db.Orders()
	if !reflect.DeepEqual(ord[0], o) {
		t.Errorf("не найден созданный заказ")
	}

	// Тест обновления записи в БД.
	o.IsOpen = false
	o.DeliveryAddress = "Адрес доставки"
	db.UpdateOrder(o)
	// Проверка.
	ord = db.Orders()
	if !reflect.DeepEqual(ord[0], o) {
		t.Errorf("не найден обновленный заказ")
	}

	// Тест удаления записи из БД.
	db.DeleteOrder(o.ID)
	// Проверка.
	ord = db.Orders()
	if len(ord) != 0 {
		t.Errorf("заказ не был удален")
	}
}

