# gin-gonic gin starter redis

## init

```
dep init
dep ensure --add github.com/gin-gonic/gin

dep ensure --add github.com/go-redis/redis
```

## launch redis first

```docker run --rm --name redis -d -p 6379:6379 redis:5-alpine```

## run

```
go run main.go
```

## Docker

```docker build -t gin-starter-redis .```

```
docker run --rm \
    -p 3000:3000 \
    --name gin-starter-redis \
    -e REDIS_HOST=192.168.1.21:6379 \
    gin-starter-redis
```

## test

````
curl localhost:3000/api/debug/ping
````

``
curl localhost:3000/api/todo/todo3
curl localhost:3000/api/todo/todo3 -X DELETE
```

```
curl -X POST \
    -H 'Content-Type: application/json' \
    -d '{"task":"something new"}' \
    localhost:3000/api/todo/todo3
```

```
curl -X POST \
    -H 'Content-Type: application/json' \
    -d '{"task":"something different"}' \
    localhost:3000/api/todo/todo3
```