# URL Shortener

[![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A simple URL shortener service built with Go, GORM, and MySQL.

## Features

- Shorten long URLs with configurable expiration
- Redirect short codes to original URLs (with expiration check)
- Automatic cleanup of expired URLs via background scheduler
- RESTful JSON API

## Tech Stack

- **Go 1.25** — application logic
- **net/http** — HTTP router (standard library)
- **GORM** — ORM / database layer
- **MySQL** — database (compatible with MariaDB)
- **Docker Compose** — local database setup

## Getting Started

### Prerequisites

- Go 1.25+
- Docker & Docker Compose (for database)

### Setup

1. Clone the repo
2. Copy `.env` and adjust if needed (see configuration below)
3. Start the database: `docker compose up -d`
4. Run the app: `go run main.go`
5. Server starts at `http://localhost:8080`

## Configuration (`.env`)

| Variable | Default | Description |
|---|---|---|
| DB_HOST | 127.0.0.1 | Database host |
| DB_DRIVER | mysql | Database driver |
| DB_PORT | 3306 | Database port |
| DB_USER | root | Database user |
| DB_PASSWORD | password | Database password |
| DB_NAME | url-shortener | Database name |
| DEFAULT_LIFESPAN | 1440 | Default URL lifespan in minutes (24h) |
| BASE_URL | http://localhost:8080 | Base URL for shortened links |
| CLEANUP_INTERVAL | 1 | Expired URL cleanup interval in minutes |
| SHORT_PATH_LENGTH | 6 | Length of generated short codes |

## API Endpoints

### `POST /api/create`

Create a shortened URL.

```json
{ "url": "https://example.com", "lifespan": 60 }
```

Optional `lifespan` in minutes; defaults to `DEFAULT_LIFESPAN`.

### `DELETE /api/delete`

Delete a shortened URL.

```json
{ "short": "abc12345" }
```

### `GET /:short`

Redirect to the original URL. Returns **404** if not found, **410 Gone** if expired.

## License

MIT
