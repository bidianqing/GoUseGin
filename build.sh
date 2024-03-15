CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./publish cmd/myappname1/main.go
cp -r configs/ ./publish
cp -r wwwroot/ ./publish