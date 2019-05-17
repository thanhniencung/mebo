GO_BUILD=CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v

build:
	cd cmd/;$(GO_BUILD);
	docker-compose up