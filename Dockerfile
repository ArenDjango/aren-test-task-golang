FROM golang:1.22

COPY . /go/src/app

WORKDIR /go/src/app/cmd/api

RUN go build -o grpc-rate-api main.go

EXPOSE 8080

CMD ["./grpc-rate-api"]
