#!/bin/bash
set -e

function checkGoEnv() {
	if [ -z "$GOPATH" ]; then
		echo "GOPATH is missing"
		exit 1
	fi

	if [[ "$(pwd)" != $GOPATH* ]]; then  
	  echo "project should be in [$GOPATH]"  
	  exit 1
	fi

}

function execLocal() {
	echo "[local] building harbor v$VERSION for os[$COMPILE_OS] arch[$COMPILE_ARCH]"
	checkGoEnv
	eval $1
}

function execDocker() {
	echo "[docker] building harbor v$VERSION for os[$COMPILE_OS] arch[$COMPILE_ARCH]"
	docker run --rm -i -v "$(pwd)":/gopath/src/github.com/elo7/harbor -e "GOPATH=/gopath" -w /gopath/src/github.com/elo7/harbor golang:latest sh -c "$1"
}

function compress() {
	mkdir gzout
	mv harbor-v$1-$2-$3 gzout/harbor
	tar -z -C gzout -cvf harbor_$1_$2_$3.tar.gz harbor
	rm -rf gzout
}


COMPILE_OS=$1
COMPILE_ARCH=$2

VERSION=$(grep "const VERSION" main.go | awk '{print substr($NF,2,length($NF)-2)}')

if [ -z "$VERSION" ]; then
	echo "could not find project version"
	exit 1
fi

BUILD_CMD="go version && go test ./... && CGO_ENABLED=0 GOOS=$COMPILE_OS GOARCH=$COMPILE_ARCH go build -v -a -installsuffix cgo --ldflags=\"-s\" -o harbor-v$VERSION-$COMPILE_OS-$COMPILE_ARCH ."
GOEXISTS=$(hash go)
if [ -z "$GOEXISTS" ]; then
	execLocal "$BUILD_CMD"
else
    execDocker "$BUILD_CMD"
fi

if [[ "$OUTPUT" == "gz" ]]; then
	compress $VERSION $COMPILE_OS $COMPILE_ARCH
fi
