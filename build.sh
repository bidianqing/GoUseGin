CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./publish/ cmd/main.go
mkdir ./publish/configs
cp -r configs/*.json ./publish/configs
cp -r wwwroot/ ./publish
cp ./cmd/Dockerfile ./publish
