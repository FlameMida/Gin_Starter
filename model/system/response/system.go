package response

import "gin-starter/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
