package redisKit

import (
	"context"
)

// Eval
/*
!!!: cluster集群模式下，使用lua脚本要注意（特别是涉及多个Redis key的情况），更多详见".info".

命令说明:	使用 Lua 解释器执行脚本
命令语法:	EVAL script numkeys key [key ...] arg [arg ...]
*/
func (client *Client) Eval(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error) {
	for i, key := range keys {
		keys[i] = client.GetKeyWithPrefix(key)
	}

	return client.universalClient.Eval(ctx, script, keys, args...).Result()
}

// EvalSha
/*
!!!: cluster集群模式下，使用lua脚本要注意（特别是涉及多个Redis key的情况），更多详见".info".
PS: 一般与 ScriptLoad 搭配使用.

命令说明:	根据给定的 sha1 校验码，执行缓存在服务器中的脚本
命令语法:	EVALSHA sha1 numkeys key [key ...] arg [arg ...]
*/
func (client *Client) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) (interface{}, error) {
	for i, key := range keys {
		keys[i] = client.GetKeyWithPrefix(key)
	}

	return client.universalClient.EvalSha(ctx, sha1, keys, args...).Result()
}

// ScriptLoad
/*
命令说明:	将脚本 script 添加到脚本缓存中，但并不立即执行这个脚本
命令语法:	SCRIPT LOAD script
命令返回值:	给定脚本的 SHA1 校验码
*/
func (client *Client) ScriptLoad(ctx context.Context, script string) (string, error) {
	return client.universalClient.ScriptLoad(ctx, script).Result()
}
