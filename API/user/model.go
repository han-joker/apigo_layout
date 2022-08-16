package User

import (
	"PH_ModuleName_PH/tool"
	"log"
)

func init() {
	// migrate
	if err := tool.DB.AutoMigrate(&User{}); err != nil {
		log.Println(err)
	}
}

type User struct {
	tool.Model `json:""`
	Username   string
	Password   string
	Token      string
	Rand       []byte
}
