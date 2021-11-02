[![Go Report Card](https://goreportcard.com/badge/github.com/thewolf27/anti-bruteforce)](https://goreportcard.com/report/github.com/thewolf27/anti-bruteforce)
![Tests](https://github.com/thewolf27/anti-bruteforce/actions/workflows/tests.yml/badge.svg)
![License](https://img.shields.io/github/license/thewolf27/anti-bruteforce)


# Анти-Брутфорс

Сервис предназначен для борьбы с подбором паролей при авторизации в какой-либо системе.

Архитектура сервиса соответствует [стандартам](https://github.com/golang-standards/project-layout) в Golang среде

Кроме того, соблюдаются принципы [чистой архитектуры](https://habr.com/ru/post/269589/), бизнес правила отделены от базы данных и REST API, упрощая замену модулей в проекте

Для учёта количества попыток авторизации каждого login/password/IP в приложении используется алгоритм [Leaky Bucket](https://en.wikipedia.org/wiki/Leaky_bucket)

Полное ТЗ можно посмотреть [здесь](https://github.com/thewolf27/anti-bruteforce/blob/main/docs/tz.md)

Проект был выполнен в рамках обучения языка Golang на курсе [Golang Developer. Professional от OTUS](https://otus.ru/lessons/golang-professional/)


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

# Command Line Interface

Для работы с CLI необходимо сбилдить бинарник приложения командой 
```
make build
```

Доступные команды для CLI:

## Blacklist
Добавить подсеть в черный список
```
./bin/app blacklist add *SUBNET*
```
Удалить подсеть из черного списка
```
./bin/app blacklist rm *SUBNET*
```

## Whitelist
Добавить подсеть в белый список
```
./bin/app whitelist add *SUBNET*
```
Удалить подсеть из белого списка
```
./bin/app whitelist rm *SUBNET*
```

## Bucket
Сбросить счётчик bucket для каждого логина/пароля/IP
```
./bin/app bucket-reset
```

## Server
Поднять HTTP и GRPC сервера приложения
```
./bin/app serve --configFolder *CONFIG_FOLDER*
```

