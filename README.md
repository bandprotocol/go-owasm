# Go-owasm

Go-owasm, is a repository designed as a Go wrapper for the [owasm-vm](https://github.com/bandprotocol/owasm/tree/master/packages/vm), enabling seamless compilation and execution of Oracle scripts within Go. This project primarily caters to the [x/oracle](https://github.com/bandprotocol/chain/tree/master/x/oracle) module in BandChain.

## Project structure

This repository contains code written in both Rust and Go. The Rust code is compiled to produce a library (`.dylib`, `.so`, `.a`). This library is then linked via cgo and encapsulated within Go. The build process involves compiling the Rust code into a C library and subsequently linking that library with the Go code. Additionally, pre-compiled libraries are included, facilitating users to import the package and directly utilize the library.

### libgo_owasm

The `./libgo_owasm` directory contains Rust code, which is compiled into a library compatible with cgo, thereby enabling its usage within Golang. Please refer to the instructions for compiling it [here](#compile-shared-and-static-libraries).

### api

The `./api` directory, section comprises Go code that interfaces with the library from [libgo_owasm](#libgo_owasm). It also incorporates the pre-compiled libraries within the folder.

### build

The `./build` directory contains scripts and docker files for building the library from Rust code.

### .github

The `./github` directory, section contains GitHub Action code that aids in testing, building libraries, and releasing packages.

## Development

If you make updates to owasm-vm or the Rust code, it's imperative to [generate bindings.sh](#generate-bindingsh-for-go) and [compile libraries](#compile-shared-and-static-libraries) to update the compiled library within the project, thereby ensuring the go-owasm package incorporates the latest version of the Owasm-vm library.

### Generate bindings.h for go

```sh
cd libgo_owasm && cargo build --release
```

### Compile shared and static libraries

Currently, libraries can be built on Linux and OS X with x86_64 architecture only. However, if you're operating on an unsupported platform, you can push the code to GitHub. Subsequently, GitHub Actions will facilitate testing, compiling the library, and pushing it to your branch.

You can use the commands below to generate libraries for Linux (x86_64) and OS X (x86_64 and aarch64).

```sh
# Run test
cd libgo_owasm && cargo test

# Build docker images that are used to compile the Rust code
make docker-images

# Run those docker images to build libraries
make releases
```
