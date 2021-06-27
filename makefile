up:
	sudo docker-compose up

in:
	sudo docker exec -ti cql /bin/bash

build:
	sudo docker exec -ti cql go build -o ./bin

run:
	sudo docker exec -ti cql ./bin/cql

test:
	sudo docker exec -ti cql go test ./lexer