FROM golang:1.21.1-alpine as builder

RUN mkdir app

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main ./

FROM alpine:latest 

RUN mkdir /app

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]
