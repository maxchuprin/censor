// сервер сервиса цензурирования
package main

import (
	"censor/pkg/api"
	"log"
	"net/http"
)

func main() {

	api := api.New()

	// Запуск сетевой службы и HTTP-сервера
	// на всех локальных IP-адресах на порту 8787.
	err := http.ListenAndServe(":8080", api.Router())
	if err != nil {
		log.Fatal(err)
	}
}
