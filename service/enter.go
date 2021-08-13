package service

import (
	"gin-starter/service/example"
	"gin-starter/service/system"
)

type Service struct {
	ExampleService example.ServiceGroup
	SystemService  system.ServiceGroup
}

var AppService = new(Service)
