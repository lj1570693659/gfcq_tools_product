SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
echo now the CGO_ENABLED:
go env CGO_ENABLED

echo now the GOOS:
go env GOOS

echo now the GOARCH:
go env GOARCH

go build -o main main.go

@REM docker build -t wxmsg_all -f .\Dockerfile .
@REM docker run -itd -p 8089:8089  --name wxmsg_all d54c7370f99d99a149d9e7c77e897859bd26e304e7b35fa115b7d4b3cc0953c3