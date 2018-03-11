#! /bin/sh
set -e

export GOARCH="amd64" GOOS="linux" CGO_ENABLED=0

if [ $# -eq 0 ]; then
    echo "version required"
    exit
fi

echo "Getting dependencies..."
go get -u -v .
echo "Building application binary..."
go build -o build/belle-stream -v .
ehco "Building docker image..."
docker build . -t sakibsami/belle-stream:$1
echo "Tasks has been completed"
