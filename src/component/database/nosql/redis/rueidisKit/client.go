package rueidisKit

import (
	"github.com/redis/rueidis"
	"github.com/richelieu-yang/chimera/v3/src/component/database/nosql/redis/redisKit"
)

func NewClient(config *redisKit.Config) (client rueidis.Client, err error) {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6379"}})
	if err != nil {
		panic(err)
	}
	defer client.Close()
}
