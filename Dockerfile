FROM golang:1.24.1 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY config.docker.toml ./config.toml

WORKDIR /app/cmd/app

RUN go build -o /app/bin/app main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=builder /app/bin/app /app
COPY --from=builder /app/config.toml /config.toml

# Heroku и другие платформы задают PORT через переменную окружения
ENV PORT=8080
EXPOSE 8080

CMD ["/app"]