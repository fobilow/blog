PHONY: compile postman

compile:
	ham build -w ./src -out ./docs
postman:
	go build -o postman main.go
