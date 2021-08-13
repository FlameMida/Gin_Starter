package router

import (
	"gin-starter/router/example"
	"gin-starter/router/system"
)

type Router struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var AppRouter = new(Router)
