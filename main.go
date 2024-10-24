package main

import (
	"MPT-CS/DataBase"
	"MPT-CS/Roaming"
)

func main() {
	DataBase.Connect()
	Roaming.ServeApplication()
}
