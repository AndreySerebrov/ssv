tests:
	go test -v ./... --timeout 30s 

generate:
	go generate ./...

run:
	go run ./cmd/server/main.go