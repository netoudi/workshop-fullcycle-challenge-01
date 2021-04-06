FROM golang:1.16-alpine

WORKDIR /app

RUN apk update
RUN apk add --no-cache \
    bash

CMD ["tail", "-f", "/dev/null"]
