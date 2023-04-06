package redisKit

import "context"

// ScriptLoad 加载脚本，返回对应的sha
func (client *Client) ScriptLoad(script string) (string, error) {
	return client.core.ScriptLoad(context.TODO(), script).Result()
}

// EvalSha 执行脚本的sha（可附带keys、args）
func (client *Client) EvalSha(sha1 string, keys []string, args ...interface{}) (interface{}, error) {
	return client.core.EvalSha(context.TODO(), sha1, keys, args...).Result()
}
