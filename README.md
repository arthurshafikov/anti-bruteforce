# Anti-bruteforce

[![Go Report Card](https://goreportcard.com/badge/github.com/thewolf27/anti-bruteforce)](https://goreportcard.com/report/github.com/thewolf27/anti-bruteforce)
![Tests](https://github.com/arthurshafikov/anti-bruteforce/actions/workflows/tests.yml/badge.svg)
![License](https://img.shields.io/github/license/arthurshafikov/anti-bruteforce)

This microservice is intended to fight against ***password brute-force attack***.

The microservice **architecture** corresponds to [Golang standarts](https://github.com/golang-standards/project-layout)

Besides, [the Clean Architecture](https://clevercoder.net/2018/09/08/clean-architecture-summary-review/) principles are **applied**, business-logic are separated from database and REST API, so we can easily replace any module on the service.

For counting authorization tries for every login/password/IP pair there is a [Leaky Bucket](https://en.wikipedia.org/wiki/Leaky_bucket) algorithm in the app.

The project was done as a final step to complete [Golang Developer. Professional from OTUS](https://otus.ru/lessons/golang-professional/) course

## Commands

### Docker-compose

Build all containers and run the app
```
make up
```

Stop all containers with the app.
```
make down
```

---

### Database

Enter PostgreSQL database console
```
make enterdb
```

---

### gRPC

Regenerate gRPC files
```
make generate
```

---

### Tests

Run unit-tests
```
make test
```

Run integration-tests
```
make integration-tests
```

Delete integration-tests containers (if there's some cache and it keeps failing)
```
make reset-integration-tests
```

## Command Line Interface

If you want to use CLI you need to build the app binary file 
```
make build
```

### Blacklist

Add subnet to the blacklist
```
./bin/app blacklist add *SUBNET*
```

Remove subnet from the blacklist
```
./bin/app blacklist rm *SUBNET*
```

### Whitelist

Add subnet to the whitelist
```
./bin/app whitelist add *SUBNET*
```

Remove subnet from the whitelist
```
./bin/app whitelist rm *SUBNET*
```

---

### Leaky Bucket

Reset Leaky-Bucket counter
```
./bin/app bucket-reset
```

---

### Server

Run the app's HTTP and gRPC servers
```
./bin/app serve --configFolder *CONFIG_FOLDER*
```

