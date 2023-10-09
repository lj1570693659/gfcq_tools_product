SET CGO_ENABLED=1
SET GOOS=linux
SET GOARCH=amd64
SET CC=x86_64-w64-mingw32-gcc
SET CXX=x86_64-w64-mingw32-g++
SET CGO_LDFLAGS="-static"
echo now the CGO_ENABLED:
go env CGO_ENABLED

echo now the GOOS:
go env GOOS

echo now the GOARCH:
go env GOARCH

go build -o main main.go