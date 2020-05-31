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

func (e *Env) SetReturnData(data []byte) {
	fmt.Println("set", string(data))
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

func (e *Env) AskExternalData(eid int64, did int64, data []byte) {
	fmt.Println("asked", eid, did, string(data))
}

func (e *Env) GetExternalData(eid int64, vid int64) []byte {
	return []byte("switez")
}

func main() {
	fmt.Println("Hello, World!")
	code, _ := ioutil.ReadFile("./wasm/fun3.wasm")
	api.Prepare(code, &Env{})
	fmt.Println("Hello, World!")
}
