

Порядок действий:

1 постановка задачи и формализация требований к приложению,
2 документирование схемы API, описание параметров каждого метода,
3 написание обработчиков для каждого метода в отдельном импортируемом пакете,
4 регистрация обработчиков в маршрутизаторе HTTP-запросов,
5 разработка исполняемого пакета, запускающего HTTP-сервер.

Мы должны разработать сервер для приложения, позволяющего создавать,
просматривать и отслеживать выполнение заказов.
Сервер должен содержать все необходимые методы
для просмотра и редактирования задач.

Метод HTTP	Метод API	    Обработчик	            Описание
GET	        /orders	        ordersHandler()	        Получение списка всех заказов
POST	    /orders	        newOrderHandler()	    Создание нового заказа
PATCH	    /orders/{id}	updateOrderHanler()	    Обновление информации заказа по его номеру — ID
DELETE	    /orders/{id}	deleteOrderHandler()	Удаление заказов по ID