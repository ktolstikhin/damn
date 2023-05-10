FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download -x
RUN go build -o=bin/api cmd/api/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin .

EXPOSE 8000

CMD ["./api", "-addr", ":8000"]
