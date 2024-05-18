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

	//r.SetPathParam()
	//
	//r.SetQueryParams()
	//r.SetQueryParamsAnyType()
	//
	//r.SetHeaderNonCanonical()
	//
	//r.SetHeaders()
	//r.SetBody()
	//r.SetContentType(header.FormContentType)
	//r.SetPathParams()
	//r.SetQueryParams()
	//r.SetQueryParamsAnyType()
	//
	//r.SetContentType()

	return
}
