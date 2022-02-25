# backend_go
backend go safebsc
# Framework
> Project with Golang, Gin

### Boilerplate structure

```
.
├── Makefile
├── Procfile
├── README.md
├── config
│   ├── config.go
│   ├── development.yaml
│   ├── production.yaml
│   └── test.yaml
├── controllers
│   └── user.go
├── db
│   └── db.go
├── forms
│   └── user.go
├── header.jpg
├── main.go
├── middlewares
│   └── auth.go
├── models
│   └── user.go
└── server
    ├── router.go
    └── server.go
```

## Installation

```sh
make deps
```

## Usage example

`curl http://localhost:8888/health`

## Development setup

Running Posgresql and Mongodb with docker compose local file.


## Run migration file for Postgresql
cat /migrates/postgresql-migratation.sql | docker exec -i {docker_postgresql_name} psql -U postgres