# Web Scraper с Worker pool

Этот проект демонстрирует многопоточный web scraper, использующий паттерны **Worker Pool** и **Fan-out/Fan-in**. Здесь показывается работа с горутинами, каналами, контекстом для graceful shutdown, а также обработку ошибок.

---

## Архитектура

Основные компоненты:
- **cmd/webscrapper/main.go** — точка входа
- **internal/app** — запуск приложения
- **internal/scrapper** — логика воркеров и менеджера задач (fan-out/fan-in)
- **internal/data** — источник URL данных в текущем случае в памяти, но можно изменить на бд и т.д.

---

Чтобы запустить
```bash
go run cmd/webscrapper/main.go
```