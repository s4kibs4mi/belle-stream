#! /bin/sh
set -e

export GOARCH="amd64" GOOS="linux" CGO_ENABLED=0
export GOPATH=$HOME/go

if [ $# -eq 0 ]; then
    echo "version required"
    exit
fi

echo "Building application binary..."
go build -o ./belle-stream -v .

echo "Building docker image..."
docker build . -t sakibsami/belle-stream:$1
echo "Tasks has been completed"
