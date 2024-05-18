package requestKit

import "github.com/imroc/req/v3"

// NewRequest
/*
@param client 可以为nil，将采用 defClient
*/
func NewRequest(client *req.Client) (r *req.Request) {
	if client == nil {
		client = defClient
	}

	r = client.R()

	return
}
