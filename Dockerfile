FROM golang:1.12-alpine as builder
COPY . /go/src/github.com/arroyo-piedra/first-app
WORKDIR /go/src/github.com/arroyo-piedra/first-app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o first-app main.go

RUN ls /go/src/github.com/arroyo-piedra/first-app/first-app
RUN chmod +x /go/src/github.com/arroyo-piedra/first-app/first-app
ENTRYPOINT ["/go/src/github.com/arroyo-piedra/first-app/first-app"]
