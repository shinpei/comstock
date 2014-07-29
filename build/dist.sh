#!/bin/bash

set -e

DIR=$(cd $(dirname ${0})/.. && pwd)
cd ${DIR}

VERSION=$(grep "const Version " version.go | sed -E 's/.*"(.+)"$/\1/')

# Compile
./build/compile.sh

# Copy dist scripts
./build/copyshell.sh

# Zip all pacakges
mkdir -p ./pkg/dist

for PLATFORM in $(find ./pkg -mindepth 1 -maxdepth 1 -type d); do
    PLATFORM_NAME=$(basename ${PLATFORM})
    ARCHIVE_NAME=comstock_${VERSION}_${PLATFORM_NAME}

    if [ $PLATFORM_NAME = "dist" ]; then
        continue
    fi

    pushd ${PLATFORM}
    zip ${DIR}/pkg/dist/${ARCHIVE_NAME}.zip ./*
    popd
done

pushd ./pkg/dist
shasum * > ./${VERSION}_SHASUMS
popd
