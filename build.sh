#!/bin/bash
# usage: ./build.sh <linux|windows|macos>

set -e
mkdir -p packages

case "$1" in
linux)
  # requires X11 headers on debian
  go build -o packages/gopeg
  ;;
windows)
  # requires `gcc-mingw-w64-x86-64` on debian
  # windows x64
  GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o packages/gopeg_x64.exe
  ;;
macos)
  # universal (arm64 + x64), must run on macos: cgo links system frameworks
  GOARCH=arm64 CGO_ENABLED=1 go build -o packages/gopeg_macos_arm64
  GOARCH=amd64 CGO_ENABLED=1 go build -o packages/gopeg_macos_x64
  lipo -create -output packages/gopeg_macos packages/gopeg_macos_arm64 packages/gopeg_macos_x64
  rm packages/gopeg_macos_arm64 packages/gopeg_macos_x64
  ;;
*)
  echo "usage: $0 <linux|windows|macos>" >&2
  exit 1
  ;;
esac
