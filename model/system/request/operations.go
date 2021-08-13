package request

import (
	"gin-starter/model/common/request"
	"gin-starter/model/system"
)

type OperationsSearch struct {
	system.Operations
	request.PageInfo
}
