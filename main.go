package main

import (
	_ "PH_ModuleName_PH/handler"
	"PH_ModuleName_PH/router"
	"log"
)

func main() {
	log.Fatalln(router.Run())
}
