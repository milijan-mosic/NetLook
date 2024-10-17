# Simple monitoring tool

## Dev setup

### Server

```sh
cd server/
air server --port 11000
```

### Client

```sh
cd client/
go mod download
go install github.com/air-verse/air@latest
go mod tidy
air server --port 10000
```
