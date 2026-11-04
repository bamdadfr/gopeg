#!/bin/bash

mkdir -p packages

# macos universal (arm64 + x64), must run on macos: cgo links system frameworks
if [ "$(uname)" = "Darwin" ]; then
  GOARCH=arm64 CGO_ENABLED=1 go build -o packages/gopeg_macos_arm64
  GOARCH=amd64 CGO_ENABLED=1 go build -o packages/gopeg_macos_x64
  lipo -create -output packages/gopeg_macos packages/gopeg_macos_arm64 packages/gopeg_macos_x64
  rm packages/gopeg_macos_arm64 packages/gopeg_macos_x64
  exit 0
fi

# linux
go build -o packages/gopeg

# requires `gcc-mingw-w64-x86-64` on debian
# windows x64
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o packages/gopeg_x64.exe
