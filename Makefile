test::
	go test ./...
run::
	go run main.go

update-deps::
	go get -u ./...
	go mod tidy
