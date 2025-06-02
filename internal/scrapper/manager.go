package scrapper

import (
	"context"
	"fmt"
	"log"
	"sync"
)

func Run(ctx context.Context, taskList []Task, WorkerCnt int) {
	tasks := make(chan Task)
	results := make(chan Result)
	var wg sync.WaitGroup

	log.Printf("Запускаем %d воркеров", WorkerCnt)

	for i := 0; i < WorkerCnt; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			Worker(ctx, id, tasks, results)
		}(i)
	}

	go func() {
		defer close(tasks)
		fmt.Println("Отправка задач в канал")

		for _, task := range taskList {
			fmt.Println("sending task url", task.Url)
			select {
			case tasks <- task:
				log.Printf("Задача отправлена %s", task.Url)
			case <-ctx.Done():
				log.Println("Отмена отправки задач по сигналу")
				return
			}
		}

		log.Println("Отправка задач завершена")
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	log.Println("Получаем результаты от воркеров")
	var SuccessCount, FailedCount int
	for result := range results {
		log.Printf("Результат: %v", result)
		if result.Err != "" {
			FailedCount++
		} else {
			SuccessCount++
		}
	}

	log.Printf("Завершено. Успешных %d, Ошибок %d", SuccessCount, FailedCount)
}
