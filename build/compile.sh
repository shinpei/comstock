#!/bin/bash

set -e

DIR=$(cd $(dirname ${0})/.. && pwd)
cd ${DIR}

XC_ARCH=${XC_ARCH:-amd64}
XC_OS=${XC_OS:-darwin linux}

rm -rf pkg/
gox \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -output "pkg/{{.OS}}_{{.Arch}}/{{.Dir}}-cli" \
    -ldflags '-w -s'


