# Cloud Native Go

## Setup
```shell
git clone git@github.com:nathanielhall/cloud-native-go.git
cd cloud-native-go
docker compose build && docker compose up
```

## Technical Decisions
- [Google: Resource-oriented design](https://cloud.google.com/apis/design/resources)
- [Chi](https://github.com/go-chi/chi) as the Router
- Postgres as the database
- [GORM](https://gorm.io/) as the ORM
- [Goose](https://github.com/pressly/goose) for database migrations
- [Zerolog](https://github.com/rs/zerolog) as the Logger
