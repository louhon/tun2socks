#!/bin/bash
export CGO_ENABLED=1
export GOOS=android
export GOARCH=arm64
export CC=~/Library/Android/sdk/ndk/21.3.6528147/toolchains/llvm/prebuilt/darwin-x86_64/bin/aarch64-linux-android21-clang
go build -o libtun2socks.so -buildmode=c-shared
~/Library/Android/sdk/ndk/21.3.6528147/toolchains/aarch64-linux-android-4.9/prebuilt/darwin-x86_64/bin/aarch64-linux-android-strip libtun2socks.so
cp -f libtun2socks.so ../ppvpn/app/proxied/android/app/src/main/assets/

