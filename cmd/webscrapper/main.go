package main

import (
	"fmt"
	"log"
	"os"
	"web-scrapper/internal/app"
)

func main() {
	fmt.Println("Запуск сервера")
	if err := app.RunServer(); err != nil {
		log.Printf("Сервер завершился с ошибкой: %v", err)
		os.Exit(1)
	}

	fmt.Println("Сервер завершился корректно")
}
