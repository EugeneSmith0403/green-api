# GREEN-API Integration

Web-приложение для демонстрации работы с [GREEN-API](https://green-api.com) — прокси-сервер на Go + Vue 3 SPA.

## Стек

| Слой | Технологии |
|---|---|
| Backend | Go 1.22, Gin, swaggo/swag |
| Frontend | Vue 3, Vite, TypeScript |
| Тесты | Go testing + httptest, Vitest + @vue/test-utils |
| Deploy | Railway (nixpacks) |

## Архитектура

```
Browser → Gin (/api/green/*) → GREEN-API
Browser → Gin (/*) → embedded Vue SPA
```

Frontend никогда не обращается напрямую к GREEN-API. Все запросы проксируются через backend.

## Локальный запуск

### Требования

- Go 1.22+
- Node.js 20+

### Переменные окружения

Скопировать `.env.example` и при необходимости изменить:

```bash
cp .env.example .env
```

| Переменная | Дефолт | Описание |
|---|---|---|
| `HOST` | `0.0.0.0` | Bind address |
| `PORT` | `8080` | Порт сервера |
| `GREEN_API_URL` | `https://api.green-api.com` | Upstream GREEN-API |

### Режим разработки

Backend и frontend dev-сервер запускаются параллельно:

```bash
make dev
```

Или в отдельных терминалах:

```bash
make backend-run    # Go сервер на :8080
make frontend-dev   # Vite dev-сервер на :5173 с proxy → :8080
```

### Production-сборка и запуск

```bash
make build          # frontend build + swagger + go build → bin/server
./bin/server
```

## Команды

```bash
make dev              # запуск backend + frontend dev параллельно
make backend-run      # только backend
make frontend-dev     # только frontend dev-сервер
make frontend-build   # собрать frontend → internal/static/dist/
make build            # полная production сборка → bin/server
make test             # все тесты (backend + frontend)
make test-backend     # go test с -race
make test-frontend    # vitest run
make lint             # golangci-lint + eslint
make fmt              # gofmt + goimports + prettier
make swagger          # регенерировать docs/ из аннотаций
```

## Тесты

```bash
make test
```

Backend покрывает: success сценарии всех 4 методов, validation errors, upstream errors.  
Frontend покрывает: API вызовы (mock fetch), валидацию форм, форматирование ответа.

## Swagger UI

После запуска сервера: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

Регенерация документации:

```bash
make swagger
```

## Deploy на Railway

1. Создать новый проект в [Railway](https://railway.app)
2. Подключить GitHub репозиторий
3. Выставить env vars в настройках сервиса:
   - `GIN_MODE=release`
4. Railway автоматически использует `nixpacks.toml` для сборки

**Pipeline сборки (nixpacks.toml):**

```
npm ci (frontend deps)
→ npm run build (Vue SPA → internal/static/dist/)
→ swag init (генерация Swagger docs)
→ go build -o bin/server (единый статический бинарник)
```

**Runtime:** один процесс `bin/server`, слушает `$PORT` (Railway выставляет автоматически).

## Методы GREEN-API

| Метод | Backend endpoint | Описание |
|---|---|---|
| `getSettings` | `POST /api/green/settings` | Настройки инстанса |
| `getStateInstance` | `POST /api/green/state` | Состояние инстанса |
| `sendMessage` | `POST /api/green/send-message` | Отправка текстового сообщения |
| `sendFileByUrl` | `POST /api/green/send-file` | Отправка файла по URL |
