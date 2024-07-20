package redisKit

import "context"

func (client *Client) XTrimMaxLen(ctx context.Context, key string, maxLen int64) (int64, error) {
	cmd := client.universalClient.XTrimMaxLen(ctx, key, maxLen)
	return cmd.Result()
}

// XTrimMaxLenApprox
/*
与 XTrimMaxLen 类似，但是它不保证精确地达到指定的长度，而是尝试接近那个长度。这意味着在某些情况下，stream的实际长度可能略高于你设置的值。
这种策略在需要更高的性能而能接受一定误差的情况下很有用。

approx: 大约
*/
func (client *Client) XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) (int64, error) {
	cmd := client.universalClient.XTrimMaxLenApprox(ctx, key, maxLen, limit)
	return cmd.Result()
}

func (client *Client) XTrimMinID(ctx context.Context, key string, minID string) (int64, error) {
	cmd := client.universalClient.XTrimMinID(ctx, key, minID)
	return cmd.Result()
}

func (client *Client) XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) (int64, error) {
	cmd := client.universalClient.XTrimMinIDApprox(ctx, key, minID, limit)
	return cmd.Result()
}
