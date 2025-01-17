package main

import (
	"MPT-CS/DataBase"
	"MPT-CS/Roaming"
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.Version())
	DataBase.Connect()
	Roaming.ServeApplication()
}
