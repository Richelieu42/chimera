package reqKit

import (
	"context"
	"github.com/imroc/req/v3"
)

// Get
/*
@param queryParams 可以为nil
*/
func Get(ctx context.Context, url string, queryParams map[string][]string) *req.Response {
	r := GetGlobalClient().Get(url)

	for key, s := range queryParams {
		r.AddQueryParams(key, s...)
	}
	return r.Do(ctx)
}
