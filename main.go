package main

import (
	"fmt"
	"io/ioutil"

	"github.com/bandprotocol/go-owasm/api"
)

func main() {
	fmt.Println("Hello, World!")
	code, _ := ioutil.ReadFile("./wasm/fun.wasm")
	fmt.Println(api.Run(code, api.Env{
		AskCount: 1000,
	}))
}
