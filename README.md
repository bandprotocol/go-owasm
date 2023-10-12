# Go-owasm

This repository serves as a wrapper of [Owasm-vm](https://github.com/bandprotocol/owasm/tree/master/packages/vm), enabling you to compile and execute Oracle scripts within Go. Mainly, it's created for [x/oracle](https://github.com/bandprotocol/chain/tree/master/x/oracle) module in BandChain.

## Project structure

This repository contains code written in both Rust and Go. The Rust code is compiled to generate a library (`.dylib`, `.so`, `.a`). This library is linked via cgo and wrapped in Go. The build process is to compile the Rust code into a C library and then link that library with the Go code. Additionally, the package includes pre-compiled libraries, allowing users to import the package and directly utilize the library.

### libgo_owasm

This folder (`./libgo_owasm`) contains Rust code. This code will be compiled into a library that can be linked via cgo and able to be used in Golang. Please see the instructions for compiling it [here](#compile-shared-and-static-libraries).

### api

This folder (`./api`) contains Go code that uses a library from [libgo_owasm](#libgo_owasm). It also includes those compiled libraries in the folder.

### build

This folder (`./build`) contains scripts and docker files that are used to build the library from Rust code.

### .github

This folder (`./github`) contains GitHub Action code that helps with testing, building libraries, and releasing packages.

## Development

If you update owasm-vm or Rust code, you will have to [generate bindings.sh](#generate-bindingsh-for-go) and [compile libraries](#compile-shared-and-static-libraries) to update the compiled library in the project to make go-owasm package use the newer version of owasm-vm library.

### Generate bindings.h for go

```sh
cd libgo_owasm && cargo build --release
```

### Compile shared and static libraries

Currently, you can build libraries on Linux and OS X with x86_64 architecture only. However, if you are working on a platform that is not currently supported, you can push code into GitHub. Then, GitHub Action will help to test, compile the library, and push it to your branch.

You can use the below commands to generate libraries for Linux (x86_64) and OS X (x86_64 and aarch64).

```sh
# Run test
cd libgo_owasm && cargo test

# Build docker images that are used to compile the Rust code
make docker-images

# Run those docker images to build libraries
make releases
```
