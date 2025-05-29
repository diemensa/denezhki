# denezhki

denezhki - это простой сервис для управления банковскими счетами

## В проекте реализовано
- Создание пользователя
- Создание счёта (account)
- Получение/обновление баланса
- Кеширование баланса через Redis
- Выполнение переводов между счетами
- Юнит-тестирование

## TODO

- JWT-аутентификация

---
## Технологии

- Go 1.24.2
- Gin web-framework
- PostgreSQL
- GORM
- Docker
- Redis
- SwaggerUI
- Тестирование (testify + mockery)
---

## Запуск


### Инструкция
1. Клонировать репозиторий:
   ```bash
   git clone https://github.com/diemensa/denezhki
   cd denezhki
2. Запустить тесты (пока что тестами покрыты только usecase):
   ```bash
   go test ./...
3. Собрать и запустить проект через docker-compose:
   ```bash
   docker-compose up --build

### Документация API
После запуска можно открыть документацию:
- SwaggerUI: http://localhost:8080/docs