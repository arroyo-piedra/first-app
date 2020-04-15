# My First Application
Demo application for meetup https://www.meetup.com/Golang-Costa-Rica/events/269507531/

# Getting started

## Requirements
- [Docker](https://docs.docker.com/engine/installation/)
- [Docker compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/doc/install)

Note: Don't forget to define $GOROOT AND $GOPATH env variables

## How to use this repo?
1. Clone this repo
``
go get -u github.com/arroyo-piedra/first-app
``

2. Build the image
``
docker-compose build
``

3. Execute the example
``
docker run -it --rm arroyo-piedra/first-app
``

4. Enjoy!!
