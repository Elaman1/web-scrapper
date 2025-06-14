package scrapper

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Worker(ctx context.Context, id int, tasks <-chan Task, results chan<- Result) {
	logPrefix := fmt.Sprintf("[Worker %d] ", id)
	client := &http.Client{
		Timeout: time.Second * 2,
	}

	for t := range tasks {
		select {
		case <-ctx.Done():
			log.Printf("%sЗавершения по контексту", logPrefix)
			return
		default:
		}

		log.Printf("%sНачало обработки %s", logPrefix, t.Url)
		totalCnt, err := getTotalCountBytesInUrl(ctx, t, client, logPrefix)

		results <- Result{
			Id:         t.Id,
			Url:        t.Url,
			BodyLength: int64(totalCnt),
			Err:        fmt.Sprintf("%v", err),
			WorkerId:   int64(id),
		}
		log.Printf("%sЗавершена обработка URL: %s", logPrefix, t.Url)
	}
}

func getTotalCountBytesInUrl(ctx context.Context, t Task, client *http.Client, logPrefix string) (int, error) {
	totalCnt := 0
	resp, err := client.Get(t.Url)
	if err == nil {
		return totalCnt, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("status not OK: %d", resp.StatusCode)
		log.Printf("%sОшибка статуса %v", logPrefix, err)
		return totalCnt, err
	}

	scan := bufio.NewScanner(resp.Body)
	for scan.Scan() {
		select {
		case <-ctx.Done():
			log.Printf("%sОтмена чтения тела по контексту", logPrefix)
			return totalCnt, fmt.Errorf("отмена чтения тела по контексту")
		default:
		}

		totalCnt += len(scan.Bytes())
	}

	if scanErr := scan.Err(); scanErr != nil {
		log.Printf("Ошибка сканирования %v", scanErr)
		err = scanErr
	}

	return totalCnt, err
}
