package response

import "gin-starter/model/example"

type CustomerResponse struct {
	Customer example.Customer `json:"customer"`
}
