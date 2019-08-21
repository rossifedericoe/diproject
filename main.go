package main

import (
	"fmt"
	"github.com/rossifedericoe/diproject/app/config/wire"
)

func main() {
	fmt.Println("Iniciando app")
	app := wire.InitializeApp()
	app.Run()
}

