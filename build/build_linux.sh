#!/bin/bash

cd go-owasm/libgo_owasm
cargo build --release
cp target/release/deps/libgo_owasm.so ./../api
