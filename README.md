# denezhki

denezhki - это простой сервис для управления банковскими счетами

# В проекте реализовано
- Создание пользователя
- Создание счёта (account)
- Получение/обновление баланса
- Кеширование баланса через Redis
- Выполнение переводов между счетами

---
# TODO

- Тесты
- JWT-аутентификация

## Технологии

- Go 1.24.2
- Gin web-framework
- PostgreSQL
- GORM
- Docker
- Docker Compose
- Redis
---

## Запуск


### Инструкция
1. Клонировать репозиторий:
   ```bash
   git clone https://github.com/diemensa/fastapi_project
   cd fastapi_project
2. Собрать и запустить проект через docker-compose:
   ```bash
   docker-compose up --build