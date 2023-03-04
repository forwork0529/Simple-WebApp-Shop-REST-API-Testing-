package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"server/packages/storage/memDB"
	"strconv"
	"testing"
)

func TestAPI_ordersHandler(t *testing.T) {

	db := memDB.New()
	o := memDB.Order{}
	db.NewOrder(o)

	api := New(db)

	rr := httptest.NewRequest(http.MethodGet, "/orders", nil)
	rp := httptest.NewRecorder()


	api.r.ServeHTTP(rp, rr)

	if !(rp.Code == http.StatusOK){
		t.Fatalf("не верный код, получен: %v, должен быть: %v\n", rp.Code, http.StatusOK)
	}
	b, err := ioutil.ReadAll(rp.Body)
	if err != nil{
		t.Fatalf("не удалось прочитать содержимое ответа")
	}
	var result []memDB.Order
	err = json.Unmarshal(b, &result)
	if err != nil{
		t.Fatalf("cand recognize objects in the answer")
	}
	if reflect.DeepEqual(result[0], o){
		t.Fatalf("полученная запись: %v, не соответствует ожидаемой: %v\n", result[0], o)
	}
}

func TestAPI_newOrderHandler(t *testing.T){
	db := memDB.New()
	api := New(db)
	toReq, _ := json.Marshal(memDB.Order{DeliveryAddress: "Moscow"})
	t.Log(string(toReq))
	rq := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(toReq) )
	rp := httptest.NewRecorder()
	api.r.ServeHTTP(rp, rq)
	if !(rp.Code == http.StatusOK){
		t.Fatalf("go response code: %v, want: %v", rp.Code, http.StatusOK)
	}
	res, err := ioutil.ReadAll(rp.Body)
	if err != nil{
		t.Fatalf("cant read from response")
	}
	id, err :=strconv.Atoi(string(res))
	if !(id == 1){
		t.Fatalf("wrong data in response")
	}
}

func TestAPI_updateOrderHandler (t *testing.T){
	o := memDB.Order{DeliveryAddress: "Lucky street"}
	db := memDB.New()
	db.NewOrder(o)
	api := New(db)
	oJ, _ := json.Marshal(o)
	req := httptest.NewRequest(http.MethodPatch, "/orders/1", bytes.NewBuffer(oJ))
	rp := httptest.NewRecorder()
	api.r.ServeHTTP(rp, req)

	if !(rp.Code == http.StatusOK){
		t.Fatalf("err: %v, code: %v", rp.Body, rp.Code)
	}
}


func TestAPI_deleteOrderHandler(t *testing.T){

	o := memDB.Order{DeliveryAddress: "Pattaya"}
	db := memDB.New()
	db.NewOrder(o)
	api := New(db)
	rq := httptest.NewRequest(http.MethodDelete, "/orders/1",nil)
	rp := httptest.NewRecorder()
	api.r.ServeHTTP(rp, rq)

	if !(rp.Code == http.StatusOK){
		t.Fatalf("err : %v, code: %v\n", rp.Body, rp.Code)
	}
}


/*func TestAPI_ordersHandler(t *testing.T) {
	// Создаем чистый объект API для теста.
	dbase := db.New()
	dbase.NewOrder(db.Order{})
	api := New(dbase)
	// Создаем HTTP-запрос.
	req := httptest.NewRequest(http.MethodGet, "/orders", nil)
	// Создаем объект для записи ответа обработчика.
	rr := httptest.NewRecorder()
	// Вызываем маршрутизатор. Маршрутизатор для пути и метода запроса
	// вызовет обработчик. Обработчик запишет ответ в созданный объект.
	api.r.ServeHTTP(rr, req)
	// Проверяем код ответа.
	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	// Читаем тело ответа.
	b, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("не удалось раскодировать ответ сервера: %v", err)
	}
	// Раскодируем JSON в массив заказов.
	var data []db.Order
	err = json.Unmarshal(b, &data)
	if err != nil {
		t.Fatalf("не удалось раскодировать ответ сервера: %v", err)
	}
	// Проверяем, что в массиве ровно один элемент.
	const wantLen = 1
	if len(data) != wantLen {
		t.Fatalf("получено %d записей, ожидалось %d", len(data), wantLen)
	}
	// Также можно проверить совпадение заказов в результате
	// с добавленными в БД для теста.
}

func TestAPI_newOrderHandler(t *testing.T) {
	dbase := db.New()
	dbase.NewOrder(db.Order{})
	api := New(dbase)

	o := db.Order{}
	b, _ := json.Marshal(o)
	req := httptest.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(b))
	rr := httptest.NewRecorder()
	api.r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
}

func TestAPI_updateOrderHandler(t *testing.T) {
	dbase := db.New()
	dbase.NewOrder(db.Order{})
	api := New(dbase)

	o := db.Order{}
	b, _ := json.Marshal(o)
	req := httptest.NewRequest(http.MethodPatch, "/orders/100", bytes.NewBuffer(b))
	rr := httptest.NewRecorder()
	api.r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
}

func TestAPI_deleteOrderHandler(t *testing.T) {
	dbase := db.New()
	dbase.NewOrder(db.Order{})
	api := New(dbase)

	req := httptest.NewRequest(http.MethodDelete, "/orders/100", nil)
	rr := httptest.NewRecorder()
	api.r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
}*/