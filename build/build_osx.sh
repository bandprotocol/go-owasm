#!/bin/bash

# ref from
## https://wapl.es/rust/2019/02/17/rust-cross-compile-linux-to-macos.html
## https://github.com/CosmWasm/wasmvm/blob/v1.1.1/builders/Dockerfile.cross
export OSXCROSS_MP_INC=1
export LIBZ_SYS_STATIC=1

cd go-owasm/libgo_owasm

echo "Starting aarch64-apple-darwin build"
export CC=aarch64-apple-darwin20.4-clang
export CXX=aarch64-apple-darwin20.4-clang++
cargo build --release --target aarch64-apple-darwin

echo "Starting x86_64-apple-darwin build"
export CC=o64-clang
export CXX=o64-clang++
cargo build --release --target x86_64-apple-darwin

# Create a universal library with both archs
lipo -output ./../api/libgo_owasm.dylib -create \
  target/x86_64-apple-darwin/release/deps/libgo_owasm.dylib \
  target/aarch64-apple-darwin/release/deps/libgo_owasm.dylib
