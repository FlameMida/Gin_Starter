package response

import "gin-starter/model/example"

type ExaFileResponse struct {
	File example.FileTransfer `json:"file"`
}
