cover: 
	go test -cover ./...

test: 
	go test ./...

build:
	go build -o gotion-cli ./cmd/...
