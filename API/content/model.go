package content

import (
	"PH_ModuleName_PH/tool"
	"log"
)

func init() {
	// migrate
	if err := tool.DB.AutoMigrate(&Content{}); err != nil {
		log.Println(err)
	}
}

type Content struct {
	tool.Model `json:""`
	Title      string `json:"title"`
	Content    string `json:"content"`
}
