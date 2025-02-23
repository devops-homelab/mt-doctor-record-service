FROM golang:1.23-alpine AS builder

ENV GOOS=linux GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o doctor-record-service .

FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/doctor-record-service .
COPY --from=builder /app/tests /root/tests
COPY --from=builder /app/go.mod /root/go.mod

EXPOSE 8084

CMD ["./doctor-record-service"]