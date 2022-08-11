package router

import "github.com/gin-gonic/gin"

func RegisterGET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.GET(relativePath, handlers...)
}

var r *gin.Engine

func init() {
	r = gin.Default()
}

func Router() *gin.Engine {
	return r
}

func Run(addr ...string) error {
	return r.Run(addr...)
}
