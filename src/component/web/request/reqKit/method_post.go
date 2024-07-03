package reqKit

import (
	"context"
	"github.com/imroc/req/v3"
)

func Post(ctx context.Context, url string, queryParams map[string][]string) *req.Response {
	r := GetGlobalClient().Post(url)

	for key, s := range queryParams {
		r.AddQueryParams(key, s...)
	}

	r.SetContentType()

	return r.Do(ctx)
}
