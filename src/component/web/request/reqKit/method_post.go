package reqKit

import (
	"context"
	"github.com/imroc/req/v3"
)

// Post TODO: 待验证.
/*
@param body Set the request Body, accepts string、[]byte、io.Reader、map and struct.
*/
func Post(ctx context.Context, url string, queryParams map[string][]string, body interface{}) *req.Response {
	r := GetGlobalClient().Post(url)

	for key, s := range queryParams {
		r.AddQueryParams(key, s...)
	}
	r.SetContentType(ContentTypeForm)
	r.SetBody(body)
	return r.Do(ctx)
}
