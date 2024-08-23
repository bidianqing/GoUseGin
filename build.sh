CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./publish cmd/main.go
cp -r configs/*.jso ./publish
cp -r wwwroot/ ./publish
cp ./cmd/Dockerfile ./publish