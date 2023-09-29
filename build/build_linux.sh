#!/bin/bash
set -o errexit -o nounset -o pipefail

cd go-owasm/libgo_owasm
/opt/build_gnu_x86_64.sh
/opt/build_gnu_aarch64.sh


# cd go-owasm/libgo_owasm 
# cargo build --release
# cp target/release/deps/libgo_owasm.so ./../api
