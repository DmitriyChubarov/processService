# builder
FROM golang:1.25.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o processService ./cmd/processing/
RUN go build -o migrate ./cmd/migrate/

# runtime
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/processService .
COPY --from=builder /app/migrate .
COPY migrations ./migrations

CMD ["sh", "-c", "./migrate && ./processService"]

