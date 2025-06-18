# Simple monitoring tool

## Dev setup

```sh
docker-compose -f docker-compose.local.yaml up --build --remove-orphans'
```

### Server

```sh
cd server/
air server --port 11000
```

### Agent

```sh
cd agent/
air server --port 10000
```
