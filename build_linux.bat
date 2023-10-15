SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
SET CC=gcc
SET CXX=g++
SET CGO_LDFLAGS="-g -O2"
go env CGO_ENABLED

echo now the GOOS:
go env GOOS

echo now the GOARCH:
go env GOARCH

go build -o main main.go