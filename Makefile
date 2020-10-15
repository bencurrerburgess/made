
.PHONY: build
build:
	go build -o ./bin/made ./main.go

.PHONY: run
run:
	go run ./main.go

.PHONY: test
test:
	go test -v ./logic
