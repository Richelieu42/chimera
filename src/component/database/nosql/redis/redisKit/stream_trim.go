package redisKit

import "context"

func (client *Client) XTrimMinID(ctx context.Context, key string, minID string) (int64, error) {
	cmd := client.universalClient.XTrimMinID(ctx, key, minID)
	return cmd.Result()
}

func (client *Client) XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) (int64, error) {
	cmd := client.universalClient.XTrimMinIDApprox(ctx, key, minID, limit)
	return cmd.Result()
}

func (client *Client) XTrimMaxLen(ctx context.Context, key string, maxLen int64) (int64, error) {
	cmd := client.universalClient.XTrimMaxLen(ctx, key, maxLen)
	return cmd.Result()
}

func (client *Client) XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) (int64, error) {
	cmd := client.universalClient.XTrimMaxLenApprox(ctx, key, maxLen, limit)
	return cmd.Result()
}
