# NetLook

## Dev setup

```sh
docker-compose -f docker-compose.local.yaml up --build --remove-orphans'
```

### Master

```sh
cd server/
air server --port 10000
```

### Agent

```sh
cd agent/
air server --port 20000
```
