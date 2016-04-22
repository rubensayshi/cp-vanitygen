#!/usr/bin/env bash

PLATFORMS="darwin/386 darwin/amd64 freebsd/386 freebsd/amd64 freebsd/arm linux/386 linux/amd64 linux/arm windows/386 windows/amd64 openbsd/386 openbsd/amd64"

rm -rf build/
mkdir build/ || exit 1

for PLATFORM in $PLATFORMS; do
    export GOOS=${PLATFORM%/*}
    export GOARCH=${PLATFORM#*/}

    EXT=""
    if [ "$GOOS" == "windows" ]; then
        EXT=".exe"
        export CGO_ENABLED=0
    fi

    echo "go build -o build/cp-vanitygen-${GOOS}-${GOARCH}${EXT} vanity.go"
    go build -o build/cp-vanitygen-${GOOS}-${GOARCH}${EXT} vanity.go
done
