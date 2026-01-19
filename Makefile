build: client server

server:
	go build -o ./bin/server.exe ./cmd/server/

client:
	go build -o ./bin/client.exe ./cmd/client/
