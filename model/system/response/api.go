package response

import "gin-starter/model/system"

type SysAPIResponse struct {
	Api system.Api `json:"api"`
}

type SysAPIListResponse struct {
	Apis []system.Api `json:"apis"`
}
