package main

import (
	"MPT-CS/DataBase"
	"MPT-CS/Roaming"
	"fmt"
	"log"
	"runtime"
)

func main() {
	fmt.Println(runtime.Version())
	DataBase.Connect()
	err := DataBase.DB.Where("email=?", "geday2@mail.ru")
	if err != nil {
		DataBase.RunMigrations()
		log.Println("tables created")
	}
	DataBase.SetDefault()
	go DataBase.ClearPincodes()
	Roaming.ServeApplication()
}
