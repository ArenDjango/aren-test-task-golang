#FROM golang:1.22
#
#COPY . /go/src/app
#
#WORKDIR /go/src/app/cmd/api
#
## Проверка содержимого директории cmd/api
#RUN ls -l /go/src/app/cmd/api
#
#RUN go build -o grpc-rate-api main.go
#
#EXPOSE 8080
#
#CMD ["./grpc-rate-api"]


FROM golang:1.22 as Builder
COPY . go/src/test-task-golang-aren
WORKDIR go/src/test-task-golang-aren/cmd

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags='-w -s' -o /go/bin/service

FROM alpine:latest
COPY --from=builder go/bin/service go/bin/service
ENTRYPOINT ["go/bin/service"]