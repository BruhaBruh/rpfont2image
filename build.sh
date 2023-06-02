#!/usr/bin/env bash

platforms=("windows/amd64" "windows/386" "linux/amd64")

rm -r -f ./build

GOOS=windows GOARCH=amd64 go build -o ./build/rpfont2image-windows-amd64.exe ./cmd/app/main.go
GOOS=windows GOARCH=386 go build -o ./build/rpfont2image-windows-386.exe ./cmd/app/main.go
GOOS=linux GOARCH=amd64 go build -o ./build/rpfont2image-linux-amd64 ./cmd/app/main.go
