FROM golang:1.24.4-alpine AS builder

ENV CGO_ENABLED=1
RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app
RUN go install github.com/air-verse/air@latest
COPY go.mod go.sum ./
RUN go mod download && go mod tidy

COPY . .
# RUN CGO_ENABLED=1 GOOS=linux go build -o /app/main

# ---------------------------------------------------------------- #

FROM golang:1.24.4-alpine

ENV CGO_ENABLED=1
RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app
# COPY --from=builder /app/main .
COPY --from=builder /go/bin/air /usr/local/bin/air
COPY --from=builder /app .
COPY air.toml ./

EXPOSE 11000
CMD ["air", "-c", "air.toml"]
