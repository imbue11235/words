build:
	go build ./...

escape:
	go build -gcflags '-m -l'

benchmark:
	go test ./... -bench=. -benchmem

cover:
	go test -coverprofile=cover.out -covermode=atomic -coverpkg=./... ./...
	go tool cover -html=cover.out -o cover.html

test:
	go test ./... -v -cover