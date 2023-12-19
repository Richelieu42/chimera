package asynqKit

import (
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/richelieu-yang/chimera/v2/src/component/database/redis/redisKit"
)

type redisConnOptImpl struct {
	asynq.RedisConnOpt

	uc redis.UniversalClient
}

func (impl *redisConnOptImpl) MakeRedisClient() interface{} {
	return impl.uc
}

// NewAsynqClient
/*
!!!: 必需先setup redis.
*/
func NewAsynqClient() (*asynq.Client, error) {
	client, err := redisKit.GetClient()
	if err != nil {
		return nil, err
	}
	impl := &redisConnOptImpl{
		uc: client.GetUniversalClient(),
	}
	return asynq.NewClient(impl), nil
}
