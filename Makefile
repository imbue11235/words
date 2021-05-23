build:
	go build ./...

escape:
	go build -gcflags '-m -l'

benchmark:
	go test ./... -bench=. -benchmem

test:
	go mod tidy -v
	go test ./...