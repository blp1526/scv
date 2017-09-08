#!/bin/bash

unames=$(uname -s)
if [ ${unames} = "Linux" ]; then
  GOOS="linux"
  cmd="tar zcvf"
  ext=".tar.gz"
elif [ ${unames} = "Darwin" ]; then
  GOOS="darwin"
  cmd="zip -r"
  ext=".zip"
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

wd="tmp"
cd ${wd}
rm -rf "scv_*"
path="scv_${GOOS}_${GOARCH}"
mkdir -p ${path}
mv scv ${path}

${cmd} "${path}${ext}" ${path} > /dev/null
rm -rf "${path}"

echo "${wd}/${path}${ext} created"
