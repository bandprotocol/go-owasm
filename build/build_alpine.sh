#!/bin/sh

cd go-owasm/libgo_owasm
cargo build --release --target x86_64-unknown-linux-musl --example muslc
cp target/x86_64-unknown-linux-musl/release/examples/libmuslc.a ./../api/libgo_owasm_muslc.x86_64.a
