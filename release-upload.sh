#!/usr/bin/env bash

TAG="$1"

if [ "${TAG}" == "" ]; then
    echo "Need tag"
    exit 1
fi

for FILE in `ls build/cp-vanitygen-*`; do
    echo "uploading ${FILE} ..."
    FILENAME=$(echo ${FILE} | sed 's/build\///g')
    echo "$GOPATH/bin/github-release upload -u rubensayshi -r cp-vanitygen --tag ${TAG} --name ${FILENAME} --file ${FILE}"
    $GOPATH/bin/github-release upload -u rubensayshi -r cp-vanitygen --tag ${TAG} --name ${FILENAME} --file ${FILE}
done
