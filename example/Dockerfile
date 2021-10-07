# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY example/go.mod ./
COPY example/go.sum ./

RUN go mod download

COPY *.go ./
RUN go build -o /example

EXPOSE 8080

CMD [ "/example" ]
