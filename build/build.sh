#!/usr/bin/env bash

set -e

# scan for vulnerabilities first
if [ "$(uname)" == "Darwin" ]; then
    ./build/nancy-darwin.amd64-v0.0.22 go.sum
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
    ./build/nancy-linux.amd64-v0.0.22 go.sum
fi

export GO111MODULE=on

# if success than go ahead an build the binaries
platforms=("windows/amd64" "windows/386" "linux/amd64" "darwin/amd64")

for platform in "${platforms[@]}"
do
    echo "building $platform..."
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name='getprs-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi  

    env GOOS=$GOOS GOARCH=$GOARCH go build -o distributions/$output_name $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution…'
        exit 1
    fi
done

echo "build finished"