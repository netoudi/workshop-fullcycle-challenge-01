# Project
A simple webserver in [Golang](https://golang.org) with [Docker](https://www.docker.com).

## Development

```bash
# start project
docker-compose up -d

# stop the container
docker-compose stop

# stop and destroy the container
docker-compose down

# re-build
docker-compose up -d --force-recreate --build

# run app
docker-compose exec app bash
go run main.go
```

## Docker push
```bash
# generate image
docker build -t tineto/workshop-fullcycle-challenge-01:latest .

# send image to docker hub
docker push tineto/workshop-fullcycle-challenge-01:latest
````

## Docker pull
````bash
# run app on port 8000 (http://localhost:8000)
docker run --rm --name challenge-01 -p 8000:8000 tineto/workshop-fullcycle-challenge-01
````

# Links
- https://golang.org
- https://www.docker.com
- https://github.com/gofiber/fiber
- https://github.com/gofiber/template/tree/master/html
