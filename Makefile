test:: lintt
	go test ./...
run:: test
	go run cmd/app/main.go

lintt::
	golangci-lint run

update-deps::
	go get -u ./...
	go mod tidy
