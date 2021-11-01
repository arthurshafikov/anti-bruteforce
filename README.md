[![Go Report Card](https://goreportcard.com/badge/github.com/thewolf27/anti-bruteforce)](https://goreportcard.com/report/github.com/thewolf27/anti-bruteforce)

# Анти-Брутфорс

Сервис предназначен для борьбы с подбором паролей при авторизации в какой-либо системе.

Архитектура сервиса соответствует [стандартам](https://github.com/golang-standards/project-layout) в Golang среде

Кроме того, соблюдаются принципы [чистой архитектуры](https://habr.com/ru/post/269589/), бизнес правила отделены от базы данных и REST API, упрощая замену модулей в проекте

Проект был выполнен в рамках обучения языка Golang на курсе [Golang Developer. Professional от OTUS](https://otus.ru/lessons/golang-professional/)

Полное ТЗ можно посмотреть [здесь](https://github.com/thewolf27/anti-bruteforce/blob/main/docs/tz.md)


---
## Docker-compose

Для поднятия приложения достаточно ввести команду
```
make up
```

Для остановки и сбрасывания приложения
```
make down
```

---
## Database

Войти в консоль базы данных 
```
make enterdb
```

---
## gRPC

Сгенерировать gRPC файлы
```
make generate
```

---
## Тесты

Запустить юнит тесты
```
make test
```

Запустить интеграционные тесты
```
make integration-tests
```

Сбросить контейнеры интеграционных тестов
```
make reset-integration-tests
```
