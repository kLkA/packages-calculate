FROM golang:1.24.1

COPY . /go/src/app

WORKDIR /go/src/app/cmd/app

RUN go build -o app main.go

EXPOSE 8080

CMD ["./app"]
