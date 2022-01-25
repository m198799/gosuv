#!/bin/bash

CURRENT_DIR=$(pwd)
BUILD_DIR=_output
BUILD_PREFIX=gosuv
GOPATH=${CURRENT_DIR}/${BUILD_DIR}
DESTINATION_DIR=${GOPATH}/src/${BUILD_PREFIX}

echo $DESTINATION_DIR

export GOPATH=$GOPATH
mkdir -p ${DESTINATION_DIR}

echo "sync gosuv to ${DESTINATION_DIR}"

rsync -a ./ --exclude=bin --exclude=${BUILD_DIR}  --exclude=hack  ${DESTINATION_DIR}/

cd ${DESTINATION_DIR} 

echo "ready to build gosuv in ${DESTINATION_DIR}"
go generate
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64
go build -tags vfs "$@"

rsync -a gosuv ${CURRENT_DIR}/bin/

echo "new gosuv in ${CURRENT_DIR}/bin"
ls -l ${CURRENT_DIR}/bin/

#go-bindata-assetfs -tags bindata res/...
#go build -tags bindata
