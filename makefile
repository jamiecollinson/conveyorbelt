build:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/conveyorbelt_darwin_amd64 cmd/main.go && GOOS=windows GOARCH=amd64 go build -o bin/conveyorbelt_windows_amd64.exe cmd/main.go && GOOS=linux GOARCH=amd64 go build -o bin/conveyorbelt_linux_amd64 cmd/main.go

run:
	go run cmd/main.go

test:
	go test

cover:
	go test -coverprofile=coverage.out && go tool cover -func=coverage.out
