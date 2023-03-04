package api

import (
	"encoding/json"
	"net/http"
	"server/packages/api/middlewares"
	"server/packages/storage/memDB"
	"strconv"

	"github.com/gorilla/mux"
)

// API приложения.
type API struct {
	r  *mux.Router // маршрутизатор запросов
	db *memDB.DB      // база данных
}

// Конструктор API.
func New(db *memDB.DB) *API {
	api := API{}
	api.db = db
	api.r = mux.NewRouter()
	api.endpoints()
	return &api
}

// Router возвращает маршрутизатор запросов.
func (api *API) Router() *mux.Router {
	return api.r
}

// Регистрация методов API в маршрутизаторе запросов.
func (api *API) endpoints() {
	api.r.HandleFunc("/orders", api.ordersHandler).Methods(http.MethodGet)
	api.r.HandleFunc("/orders", api.newOrderHandler).Methods(http.MethodPost)
	api.r.HandleFunc("/orders/{id}", api.updateOrderHandler).Methods(http.MethodPatch)
	api.r.HandleFunc("/orders/{id}", api.deleteOrderHandler).Methods(http.MethodDelete)
	api.r.Use(mux.MiddlewareFunc(middlewares.MddlwrJsonTypeHeader))
}

// ordersHandler возвращает все заказы.
func (api *API) ordersHandler(w http.ResponseWriter, r *http.Request) {
	// Получение данных из БД.
	orders := api.db.Orders()
	// Отправка данных клиенту в формате JSON.
	json.NewEncoder(w).Encode(orders)
}

// newOrderHandler создает новый заказ.
func (api *API) newOrderHandler(w http.ResponseWriter, r *http.Request) {
	var o memDB.Order
	err := json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id := api.db.NewOrder(o)
	w.Write([]byte(strconv.Itoa(id)))
}

// updateOrderHandler обновляет данные заказа по ID.
func (api *API) updateOrderHandler(w http.ResponseWriter, r *http.Request) {
	// Считывание параметра {id} из пути запроса.
	// Например, /orders/45.
	s := mux.Vars(r)["id"]

	id, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Декодирование в переменную тела запроса,
	// которое должно содержать JSON-представление
	// обновляемого объекта.
	var o memDB.Order
	err = json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	o.ID = id
	// Обновление данных в БД.
	api.db.UpdateOrder(o)
	// Отправка клиенту статуса успешного выполнения запроса
	w.WriteHeader(http.StatusOK)
}

// deleteOrderHandler удаляет заказ по ID.
func (api *API) deleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	s := mux.Vars(r)["id"]
	id, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	api.db.DeleteOrder(id)
	w.WriteHeader(http.StatusOK)
}