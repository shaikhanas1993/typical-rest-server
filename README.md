<!-- Autogenerated by Typical-Go. DO NOT EDIT. -->

# Typical REST Server

Example of typical and scalable RESTful API Server for Go

### Usage

### Configuration

| Name | Type | Default | Required |
|---|---|---|:---:|
|SERVER_DEBUG|bool|false||
|PG_DBNAME|string|typical-rest|Yes|
|PG_USER|string|postgres|Yes|
|PG_PASSWORD|string|pgpass|Yes|
|PG_HOST|string|localhost||
|PG_PORT|int|5432||
|REDIS_HOST|string|localhost|Yes|
|REDIS_PORT|string|6379|Yes|
|REDIS_PASSWORD|string|redispass||
|REDIS_DB|int|0||
|REDIS_POOL_SIZE|int|20|Yes|
|REDIS_DIAL_TIMEOUT|Duration|5s|Yes|
|REDIS_READ_WRITE_TIMEOUT|Duration|3s|Yes|
|REDIS_IDLE_TIMEOUT|Duration|5m|Yes|
|REDIS_IDLE_CHECK_FREQUENCY|Duration|1m|Yes|
|REDIS_MAX_CONN_AGE|Duration|30m|Yes|
|APP_ADDRESS|string|:8089|Yes|

----

## Development Guide

### Prerequisite

Install [Go](https://golang.org/doc/install) (It is recommend to install via [Homebrew](https://brew.sh/) `brew install go`)

### Build & Run

Use `./typicalw run` to build and run the project.

### Test

Use `./typicalw test` to test the project.

### Release the destribution

Use `./typicalw release` to make the release. [Learn More](https://typical-go.github.io/learn-more/build-tool/release-distribution.html)

### Other Command

- `./typicalw docker`: Docker utility
	- `./typicalw docker compose`: Generate docker-compose.yaml
	- `./typicalw docker up`: Spin up docker containers
	- `./typicalw docker down`: Take down all docker containers
- `./typicalw postgres`: Postgres Database Tool
	- `./typicalw postgres create`: Create New Database
	- `./typicalw postgres drop`: Drop Database
	- `./typicalw postgres migrate`: Migrate Database
	- `./typicalw postgres rollback`: Rollback Database
	- `./typicalw postgres seed`: Database Seeding
	- `./typicalw postgres console`: PostgreSQL Interactive
- `./typicalw redis`: Redis Tool
	- `./typicalw redis console`: Redis Interactive
