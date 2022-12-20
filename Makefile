.PHONY: start-server
start-server:
	go build -o app ./cmd/calc/*.go
	./app