FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .

RUN apk add --no-cache gcc musl-dev sqlite-dev

RUN go mod download && go mod tidy

FROM golang:1.23-alpine

WORKDIR /app
COPY --from=builder /app .

RUN go install github.com/air-verse/air@latest

EXPOSE 10000
CMD air server --port 10000 -c .air.toml
