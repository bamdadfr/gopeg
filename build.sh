#!/bin/bash

mkdir -p packages

# linux
go build -o packages/gopeg

# requires `gcc-mingw-w64-x86-64` on debian
# windows x64
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o packages/gopeg_x64.exe
