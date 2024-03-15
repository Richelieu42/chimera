package mongodbKit

import (
	"context"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

// NewClientSimply
/*
PS: 单点启动MongoDB，默认使用 27017 端口.

@param uri e.g."mongodb://localhost:27017"
*/
func NewClientSimply(ctx context.Context, uri string) (*mongo.Client, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI(uri)
	// The default is 0, meaning a connection can remain unused indefinitely.
	clientOptions.SetMaxConnIdleTime(0)
	// The default is 0.
	clientOptions.SetMinPoolSize(0)
	// 设置连接池大小（默认: 100）
	// If this is 0, maximum connection pool size is not limited. The default is 100.
	clientOptions.SetMaxPoolSize(100)

	return NewClient(ctx, clientOptions)
}

func NewClient(ctx context.Context, opts ...*options.ClientOptions) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, opts...)
	if err != nil {
		return nil, errorKit.Wrap(err, "fail to connect")
	}

	// 检查连接是否成功
	ctx1, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	if err := client.Ping(ctx1, readpref.Primary()); err != nil {
		_ = client.Disconnect(context.TODO())
		return nil, errorKit.Wrap(err, "fail to ping")
	}

	return client, nil
}
