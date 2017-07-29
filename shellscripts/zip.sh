#!/bin/bash

unames=$(uname -s)
if [ ${unames} = "Linux" ]; then
  GOOS="linux"
elif [ ${unames} = "Darwin" ]; then
  GOOS="darwin"
else
  echo "fatal: Unsupported OS"
  exit 1
fi

unamem=$(uname -m)
if [ ${unamem} = "x86_64" ]; then
  GOARCH="amd64"
else
  echo "fatal: Unsupported ARCH"
  exit 1
fi

cd "tmp"
rm -rf "scv_*"
path="scv_${GOOS}_${GOARCH}"
mkdir -p ${path}
mv scv ${path}
zip -r "${path}.zip" ${path} > /dev/null
rm -rf "${path}"

echo "${path}.zip created"
