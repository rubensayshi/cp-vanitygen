#!/usr/bin/env bash

TAG="$1"

if [ "${TAG}" == "" ]; then
    echo "Need tag"
    exit 1
fi

git tag $TAG || exit 1
git push --tags || exit 1

$GOPATH/bin/github-release release -u rubensayshi -r cp-vanitygen --tag ${TAG} --name "Counterparty VanityGen ${TAG}"

./release-upload.sh ${TAG}
