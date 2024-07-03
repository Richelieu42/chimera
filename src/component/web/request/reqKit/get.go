package reqKit

import "context"

func Get(ctx context.Context, url string) {
	GetGlobalClient().Get(url).
		SetContext(ctx).Do()
}
