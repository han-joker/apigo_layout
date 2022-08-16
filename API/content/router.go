package content

import (
	"PH_ModuleName_PH/router"
)

func init() {
	// Crud
	group := router.RegisterGroup("/contents")
	group.POST("", Post)
	group.DELETE("/:id", Delete)
	group.PUT("/:id", Put)
	group.GET("", List)
	group.GET("/:id", Retrieve)
}
