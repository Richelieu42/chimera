package redisKit

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// TestSingleNodeMode 测试Redis的Cluster集群模式
func TestSingleNodeMode(test *testing.T) {
	config := Config{
		UserName: "",
		Password: "",
		Mode:     SingleNodeMode,
		Single: &SingleConfig{
			Addr: "127.0.0.1:6379",
			DB:   10,
		},
		MasterSlave: nil,
		Sentinel:    nil,
		Cluster:     nil,
	}
	client, err := NewClient(config)
	if err != nil {
		panic(err)
	}

	if _, err := client.Ping(context.TODO()); err != nil {
		panic(err)
	}

	fmt.Println(client.SetNX(context.TODO(), "1", "222", time.Second*30))
	fmt.Println(client.TTL(context.TODO(), "1"))

	time.Sleep(time.Second * 3)

	fmt.Println(client.SetXX(context.TODO(), "1", "999", 0))
	fmt.Println(client.TTL(context.TODO(), "1"))

	//fmt.Println(client.TTL(context.TODO(), "1"))
	//
	//fmt.Println(client.SetEx(context.TODO(), "1", "222", time.Second*30))
	//fmt.Println(client.TTL(context.TODO(), "1"))
	//
	//time.Sleep(time.Second * 5)
	//
	//fmt.Println(client.SetEx(context.TODO(), "1", "333", time.Second*30))
	//fmt.Println(client.TTL(context.TODO(), "1"))
}

func TestClusterMode(test *testing.T) {
	config := Config{
		UserName: "",
		Password: "",
		Mode:     ClusterMode,
		Cluster: &ClusterConfig{
			Addrs: []string{
				//"127.0.0.1:6380",
				//"127.0.0.1:6381",
				//"127.0.0.1:6382",
				//"127.0.0.1:6383",
				//"127.0.0.1:6384",
				//"127.0.0.1:6385",

				"192.168.80.43:7000",
				"192.168.80.43:7001",
				"192.168.80.27:7002",
				"192.168.80.27:7003",
				"192.168.80.42:7004",
				"192.168.80.42:7005",
			},
		},
	}
	client, err := NewClient(config)
	if err != nil {
		panic(err)
	}

	if str, err := client.Ping(context.TODO()); err != nil {
		panic(err)
	} else {
		fmt.Println(str)
	}

	//for i := 0; i < 100; i++ {
	//	_, _ = client.Set(context.TODO(), strconv.Itoa(i), strconv.Itoa(i), 0)
	//}

	//for i := 0; i < 1000; i++ {
	//	s, err := client.ScanFully(context.TODO(), "*", 10)
	//	if err != nil {
	//		panic(err)
	//	}
	//	if len(s) != 101 {
	//		panic(len(s))
	//	}
	//	fmt.Printf("====== %d\n", len(s))
	//}

	//c := client.GetUniversalClient()

	//sc := c.Scan(context.TODO(), 0, "*", 10)
	//iter := sc.Iterator()
	//for iter.Next(context.TODO()) {
	//	fmt.Println(iter.Val())
	//}
}
