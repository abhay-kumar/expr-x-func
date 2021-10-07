# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

RUN cd example

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./
RUN go build -o /example

EXPOSE 8080

CMD [ "/example" ]
