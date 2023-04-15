# Cloud Native Go
> This repository was created as a learning resource. 

## Setup
```shell
git clone git@github.com:nathanielhall/cloud-native-go.git
cd cloud-native-go
docker compose build && docker compose up
```

## Technical Decisions
- [Chi](https://github.com/go-chi/chi) as the Router
- Postgres as the database
- [GORM](https://gorm.io/) as the ORM
- [Goose](https://github.com/pressly/goose) for database migrations
- [Zerolog](https://github.com/rs/zerolog) as the Logger

## Resources
- [Google: Resource-oriented design](https://cloud.google.com/apis/design/resources)
- [Learning Cloud Native Go](https://learning-cloud-native-go.github.io/)