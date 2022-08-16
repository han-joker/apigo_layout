package User

import (
	"PH_ModuleName_PH/router"
)

func init() {
	// Crud
	group := router.RegisterGroup("/users")
	group.POST("", Post)
	group.DELETE("/:id", Delete)
	group.PUT("/:id", Put)
	group.GET("", List)
	group.GET("/:id", Retrieve)

	// auth
	group.POST("/sign-out", SignOut)

	noAuth := group.Group("")
	noAuth.POST("/sign-in", SignIn)
	noAuth.POST("/sign-up", SignUp)
}
