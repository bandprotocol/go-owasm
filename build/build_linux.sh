#!/bin/bash
set -o errexit -o nounset -o pipefail

# cd go-owasm/libgo_owasm 
build_gnu_x86_64.sh
build_gnu_aarch64.sh


# cargo build --release
# cp target/release/deps/libgo_owasm.so ./../api
