#!/usr/bin/env bash

version=$@

echo "Create assets for $version"

for os in darwin linux windows; do
  for arch in amd64 386; do
    echo "Build $os $arch"
    GOOS=$o GOARCH=$a go build -v github.com/bzlparty/tool-versions-generator/cmd/tvg

    if [[ "$os" == "windows" ]]; then
      zip tvg-$version-$os_$arch.zip tvg.exe LICENSE README.md
      rm tvg.exe
    else
      tar cvf tvg-$version-$os_$arch.tar.gz tvg LICENSE README.md
      rm tvg
    fi
  done
done
