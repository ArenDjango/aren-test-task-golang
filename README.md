# Aren Test task for GOLANG Interview

### I created a test task where it was necessary to make an API request to receive a rate, save it to the database, make a grpc server, healt check grpc server, and write a unit test, also connect prometeus, grafana, opentelemetry for tracing, also take configs from env file or from command line flags

This is Golang GRPC server example including the following features:
*   made with Clean Architecture in mind (controller -> service -> repository)
*   has services that work with both PostgreSQL or other database if you need, because its working by interface
*   Go go tests based on mocks auto-generated with go:generate and mockery (<https://github.com/vektra/mockery>)
*   config based on envconfig
* Prometeus, opentelemetry
* Database migrations
* Smart logging architecture
* Make file


### Run

docker-compose up