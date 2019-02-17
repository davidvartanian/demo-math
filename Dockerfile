FROM golang:latest

WORKDIR /demo_math
RUN go get -u github.com/swaggo/swag/cmd/swag

COPY . /demo-math
RUN swag init
RUN go build