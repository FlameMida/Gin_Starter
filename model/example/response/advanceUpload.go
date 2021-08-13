package response

import "gin-starter/model/example"

type FilePathResponse struct {
	FilePath string `json:"filePath"`
}

type FileResponse struct {
	File example.File `json:"file"`
}
