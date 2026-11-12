# sample-health-check

A minimal Go HTTP service with a health check endpoint.

## Run

```sh
go run .
```

Port defaults to `8080`; override with the `PORT` environment variable.

## Endpoints

| Path | Description |
| --- | --- |
| `GET /` | Plain-text "Server is running" message |
| `GET /health` | JSON health check |
| `GET /swagger/` | Swagger UI (API docs) |
| `GET /swagger/openapi.json` | OpenAPI 3 spec |

The spec lives in `openapi.json` and is embedded into the binary; update it when you add endpoints.
Swagger UI assets load from the jsDelivr CDN, so the browser needs internet access.

### `GET /health`

```json
{
  "status": "UP",
  "timestamp": "2026-11-12T10:00:00Z",
  "uptime": "42s",
  "version": "1.0.0"
}
```

## Test

```sh
go test ./...
```
