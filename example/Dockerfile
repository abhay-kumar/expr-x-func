# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY example/go.mod ./
COPY example/go.sum ./

RUN go mod download

COPY . ./
RUN go build -o /example example/main.go

EXPOSE 8080

RUN chmod +x /example

CMD [ "/example" ]
