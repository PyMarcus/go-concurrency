run:
	go run .

test:
	go test -v ./...

race:
	go run -race .