#!/usr/bin/env bash

version=$@

function build_and_pack()
{
  GOOS=$1 GOARCH=$2 go build -v github.com/bzlparty/tool-versions-generator/cmd/tvg
  if [[ "$1" == "windows" ]]; then
    zip "tvg-${version}-${1}_${2}.zip" tvg.exe LICENSE README.md
    rm tvg.exe
  else
    tar cvf "tvg-${version}-${1}_${2}.tar.gz" tvg LICENSE README.md
    rm tvg
  fi
}

echo "Create assets for $version"

build_and_pack linux amd64
build_and_pack linux 386
build_and_pack darwin amd64
build_and_pack windows amd64
build_and_pack windows 386
