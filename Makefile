up:
	docker-compose up

down:
	docker-compose down

tidy:
	cd src/core && go mod tidy
	cd src/adapter && go mod tidy
	cd src/migration && go mod tidy
	cd src/public && go mod tidy

fmt:
	cd src/core && go fmt ./...
	cd src/adapter && go fmt ./...
	cd src/migration && go fmt ./...
	cd src/public && go fmt ./...

test:
	cd src/core && go mod tidy && go test ./...
	cd src/adapter && go mod tidy && go test ./...
	cd src/migration && go mod tidy && go test ./...
	cd src/public && go mod tidy && go test ./...

swagger-public:
	cd src/public && swag init --parseDependency --parseDepth=3

love:
	cd src/core && gosec ./...
	cd src/adapter && gosec ./...
	cd src/migration && gosec ./...
	cd src/public && gosec ./...

vet:
	cd src/core && go vet ./...  && staticcheck ./...
	cd src/adapter && go vet ./... && staticcheck ./...
	cd src/migration && go vet ./... && staticcheck ./...
	cd src/public && go vet ./... && staticcheck ./...
