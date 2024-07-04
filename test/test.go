package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/imroc/req/v3"
)

// RoundRobinBalancer 简单的轮询负载均衡器
type RoundRobinBalancer struct {
	servers []string
	index   int
	mu      sync.Mutex
}

// NewRoundRobinBalancer 创建一个新的轮询负载均衡器
func NewRoundRobinBalancer(servers []string) *RoundRobinBalancer {
	return &RoundRobinBalancer{
		servers: servers,
	}
}

// Next 获取下一个服务器
func (r *RoundRobinBalancer) Next() string {
	r.mu.Lock()
	defer r.mu.Unlock()
	server := r.servers[r.index]
	r.index = (r.index + 1) % len(r.servers)
	return server
}

func main() {
	servers := []string{
		"https://server1.example.com/api",
		"https://server2.example.com/api",
		"https://server3.example.com/api",
	}

	balancer := NewRoundRobinBalancer(servers)
	client := req.C()

	for i := 0; i < 10; i++ { // 模拟10个请求
		go func() {
			server := balancer.Next()
			resp, err := client.R().
				SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=utf-8").
				SetBody("key=value").
				Post(server)

			if err != nil {
				log.Printf("Error: %v", err)
				return
			}

			fmt.Printf("Response from %s: %s\n", server, resp.Status)
		}()
	}

	// 保持程序运行以等待所有 goroutine 完成
	time.Sleep(5 * time.Second)
}
