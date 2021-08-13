package v1

import (
	"gin-starter/api/v1/example"
	"gin-starter/api/v1/system"
)

type ApiGroup struct {
	ExampleApi example.ApiGroup
	SystemApi  system.ApiGroup
}

var AppApi = new(ApiGroup)
