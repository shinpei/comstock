#!/bin/bash

set -e
DIR=$(cd $(dirname ${0})/.. && pwd)
echo "dir: " ${DIR}
cd ${DIR}
mkdir -p ${DIR}/pkg/scripts
for file in ${DIR}/dist/*; do
    echo ${file}
    cp ${file} ${DIR}/pkg/scripts/
done

