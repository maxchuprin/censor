//api сервера сервиса цензурирования
package api

import (
	"censor/pkg/censor"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// API приложения.
type API struct {
	r *mux.Router // маршрутизатор запросов
}

// Конструктор API.
func New() *API {
	api := API{}
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
	api.r.HandleFunc("/check", api.check).Methods(http.MethodPost)
	api.r.Use(api.HeadersMiddleware)
	api.r.Use(api.RequestIDMiddleware)
	api.r.Use(api.LoggingMiddleware)
}

// проверяет текст, переданный в теле запроса на стоп-слова
func (api *API) check(w http.ResponseWriter, r *http.Request) {

	reqID := r.Context().Value(contextKey("requestID")).(int)

	var c = struct {
		Text string
	}{}

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error()+":requestID: "+strconv.Itoa(reqID), http.StatusInternalServerError)
		return
	}

	if censor.Censored(c.Text) {
		http.Error(w, "Censored!", http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	//запись ID запроса в тело ответа
	a := struct {
		RequestID int
	}{
		RequestID: reqID,
	}
	//Отправка данных клиенту в формате JSON.
	json.NewEncoder(w).Encode(a)

}
