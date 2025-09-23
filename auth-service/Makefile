include .env
export

deps: 
	go mod tidy && go mod verify

run: deps
	go mod download && \
	CGO_ENABLED=0 go run ./cmd/app