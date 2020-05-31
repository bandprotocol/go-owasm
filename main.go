package main

import (
	"fmt"
	"io/ioutil"

	"github.com/bandprotocol/go-owasm/api"
)

type Env struct{}

func (e *Env) GetCalldata() []byte {
	return []byte("switza")
}

func (e *Env) GetAskCount() int64 {
	return 10000
}

func (e *Env) GetMinCount() int64 {
	return 20000
}

func (e *Env) GetAnsCount() int64 {
	return 30000
}

func main() {
	fmt.Println("Hello, World!")
	code, _ := ioutil.ReadFile("./wasm/fun2.wasm")
	fmt.Println(api.Run(code, &Env{}))
}
