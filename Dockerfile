FROM golang:1.22.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ret3iac .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ret3iac .

ENTRYPOINT ["./ret3iac"]
