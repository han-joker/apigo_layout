package router

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func Register(httpMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return r.Handle(strings.ToUpper(httpMethod), relativePath, handlers...)
}

func RegisterGroup(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return r.Group(relativePath, handlers...)
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
