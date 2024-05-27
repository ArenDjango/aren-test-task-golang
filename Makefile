mocks:
	cd ./store/mocks/; go generate;
	cd ./service/mocks/; go generate;

build:
	go build -o cmd/api/grpc-rate-api cmd/api/main.go

test:
	go test ./...

docker-build:
	docker-compose build