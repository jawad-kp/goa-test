.PHONY: start-server
start-server:
	go build -o app ./cmd/calc
	./app

.PHONY: start-server-win
start-server-win:
	go build -o app.exe goa-test/cmd/calc
	app.exe