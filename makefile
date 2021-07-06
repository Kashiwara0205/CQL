up:
	sudo docker-compose up

in:
	sudo docker exec -ti cql /bin/bash

build:
	sudo docker exec -ti cql go build -o ./bin

run:
	sudo docker exec -ti cql ./bin/cql

start:
	sudo docker exec -ti cql go run main.go

test:
	sudo docker exec -ti cql go test -v ./lexer
	sudo docker exec -ti cql go test -v ./token
	sudo docker exec -ti cql go test -v ./parser
	sudo docker exec -ti cql go test -v ./evaluator