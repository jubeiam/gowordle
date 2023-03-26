FROM golang:1.20-alpine as base
WORKDIR /app
COPY go.mod ./
RUN go mod download

FROM base as builder
COPY main.go ./
RUN go build -o /usr/local/go/src/bin/wordle main.go
CMD [ "go", "run", "wordle"]

