.PHONY: run
run:
	clear; go build -o bin/ && ./bin/olist-api

.PHONY: build
build:
	go build -o bin/

.PHONY: serve
serve:
	clear; ./node_modules/.bin/nodemon --exec go run *.go