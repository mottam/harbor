#!/bin/bash
set -e

COMPILE_OS=$1
COMPILE_ARCH=$2
VERSION=$(grep "const VERSION" main.go | awk '{print substr($NF,2,length($NF)-2)}')
if [ -z "$VERSION" ]; then
	echo "could not find project version"
	exit 1
fi

echo "building harbor v$VERSION for os[$COMPILE_OS] arch[$COMPILE_ARCH]"
docker run --rm -i -v "$(pwd)":/gopath/src/github.com/elo7/harbor -e "GOPATH=/gopath" -w /gopath/src/github.com/elo7/harbor golang:latest sh -c "pwd && go version && go test ./... && CGO_ENABLED=0 GOOS=$COMPILE_OS GOARCH=$COMPILE_ARCH go build -v -a -installsuffix cgo --ldflags=\"-s\" -o harbor-v$VERSION-$COMPILE_OS-$COMPILE_ARCH ."
