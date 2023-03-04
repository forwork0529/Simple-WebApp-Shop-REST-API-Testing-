package main

import (
	"log"
	"net/http"
	api2 "server/packages/api"
	"server/packages/storage/memDB"
)

func main() {
	// Инициализация БД в памяти.

	ord := memDB.Order {ID: 1, IsOpen: true, DeliveryTime: 155, DeliveryAddress: "Manhattan", Products: []memDB.Product{}}
	db := memDB.New()
	db.NewOrder(ord)
	// Создание объекта API, использующего БД в памяти.
	api := api2.New(db)
	// Запуск сетевой службы и HTTP-сервера
	// на всех локальных IP-адресах на порту 80.
	err:=http.ListenAndServe(":8080", api.Router())
	if err != nil {
		log.Fatal(err)
	}
}