@echo off

set /p name="Enter name: "

mkdir %name%
cd %name%
echo. > main.go
echo package main >> main.go
echo. >> main.go
echo func main() { >> main.go
echo. >> main.go
echo } >> main.go

go mod init %name%

echo Script executed successfully!

