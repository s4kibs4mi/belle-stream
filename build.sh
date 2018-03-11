#! /bin/sh
set -e

export GOARCH="amd64" GOOS="linux" CGO_ENABLED=0

if [ $# -eq 0 ]; then
    echo "version required"
    exit
fi

go get -u .
go build -o build/belle-stream -v .

docker build . -t sakibsami/belle-stream:$1
