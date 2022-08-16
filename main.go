package main

import (
	_ "PH_ModuleName_PH/API/content"
	_ "PH_ModuleName_PH/API/user"
	"PH_ModuleName_PH/router"
	_ "PH_ModuleName_PH/tool"
	"log"
)

func main() {
	log.Fatalln(router.Run())
}
