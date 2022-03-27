package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func NewRouter() routes {
	r := routes{
		router: gin.Default(),
	}

	return r
}

func (r routes) Run(addr ...string) error {
	return r.router.Run(addr...)
}

func (r routes) Handler() http.Handler {
	return r.router
}
