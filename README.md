# Cloud Native Go
> This repository provides a CRUD API for Todo items.

## Setup
```shell
git clone git@github.com:nathanielhall/cloud-native-go.git
cd cloud-native-go
docker compose build && docker compose up
```

## Todos

| Name        | HTTP Method | Route          |
|-------------|-------------|----------------|
| List  | GET         | /v1/todos      |
| Create | POST        | /v1/todos      |
| Read | GET         | /v1/todos/{id} |
| Update | PUT         | /v1/todos/{id} |
| Delete | DELETE      | /v1/todos/{id} |


## Technical Decisions
- [Chi](https://github.com/go-chi/chi) as the Router
- Postgres as the database
- [GORM](https://gorm.io/) as the ORM
- [Goose](https://github.com/pressly/goose) for database migrations
- [Zerolog](https://github.com/rs/zerolog) as the Logger

## Resources
- [Google: Resource-oriented design](https://cloud.google.com/apis/design/resources)
- [Learning Cloud Native Go](https://learning-cloud-native-go.github.io/)