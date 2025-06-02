package app

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"web-scrapper/internal/data"
	"web-scrapper/internal/scrapper"
)

func RunServer() error {
	if err := loadEnv(); err != nil {
		return fmt.Errorf("ошибка загрузки конфигурации %w", err)
	}

	dataSrc, err := data.GetProviderData(os.Getenv("DATA_SRC"))
	if err != nil {
		return fmt.Errorf("ошибка инициализации источника данных: %w", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	tasks := prepareTasks(dataSrc.GetUrlList())

	workerCnt := getWorkerCount()
	log.Printf("Запуск воркеров (количество: %d)", workerCnt)
	scrapper.Run(ctx, tasks, workerCnt)
	return nil
}

func prepareTasks(urls []string) []scrapper.Task {
	tasks := make([]scrapper.Task, len(urls))
	for id, url := range urls {
		tasks[id] = scrapper.Task{
			Id:     int64(id),
			Url:    url,
			Method: http.MethodGet,
		}
	}
	return tasks
}

func loadEnv() error {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Не удалось загрузить .env файл")
		return err
	}
	return nil
}

func getWorkerCount() int {
	workerCnt, err := strconv.Atoi(os.Getenv("WORKER_CNT"))
	if err != nil || workerCnt <= 0 {
		log.Println("WORKER_CNT не задан или некорректен, используем значение по умолчанию (5)")
		return 5
	}
	return workerCnt
}
