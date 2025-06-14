# denezhki

denezhki - это простой сервис для управления банковскими счетами

## В проекте реализовано
- Создание пользователя
- Создание счёта (account)
- Получение/обновление баланса
- Кеширование баланса через Redis
- Выполнение переводов между счетами
- Тестирование бизнес-логики
- Аутентификация с JWT-токеном

---
## Технологии

- Go 1.24.2
- Gin web-framework
- PostgreSQL
- GORM
- Docker
- Redis
- SwaggerUI
- JWT
- Unit-тестирование (testify + mockery)
---

## Запуск
### Linux/MacOS
1. Клонировать репозиторий:
   ```bash
   git clone https://github.com/diemensa/denezhki
   cd denezhki
2. Запустить тесты и приложение:
   ```bash
   make
   
### Windows
1. Клонировать репозиторий:
   ```bash
   git clone https://github.com/diemensa/denezhki
   cd denezhki
2. Запустить тесты:
   ```bash
   go test -v ./...
3. Собрать и запустить проект через docker-compose:
   ```bash
   docker-compose up --build

### Документация API
После запуска можно открыть документацию:
- SwaggerUI: http://localhost:8080/docs/index.html