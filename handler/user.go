package handler

import (
	"PH_ModuleName_PH/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	router.RegisterGET("/user", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "user",
		})
	})
}
