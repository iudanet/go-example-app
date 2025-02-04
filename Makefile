test::
	go test ./...
run:: test
	go run cmd/app/main.go

update-deps::
	go get -u ./...
	go mod tidy
