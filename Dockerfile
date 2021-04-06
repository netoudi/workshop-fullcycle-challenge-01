FROM golang:1.16-alpine as builder

WORKDIR /app

#### COPY ####
COPY . .

#### BUILD ###
RUN go get -d -v
RUN GOOS=linux go build -ldflags="-w -s" -o /app/main

#### OPTIMIZE ####
FROM alpine:3.13.4

COPY --from=builder /app/main /app/main
COPY --from=builder /app/views /app/views
COPY --from=builder /app/public /app/public

ENTRYPOINT ["/app/main"]
