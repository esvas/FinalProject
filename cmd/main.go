package main

import (
	"log"
	"net/http"
	"os"

	"github.com/esvas/FinalProject/api/routes"
)

func main() {
	log.Println("Запуск сервера")
	r := routes.CreateRouter()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Printf("Не удалось запустить сервер:\n%v", err)
		os.Exit(1)
	}
}