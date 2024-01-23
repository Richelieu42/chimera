package raftKit

import (
	"fmt"
	"github.com/gogf/gf/v2/container/gtype"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
	"github.com/richelieu-yang/chimera/v2/src/atomicKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v2/src/file/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/micro/raft/raftLogKit"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
	"net"
	"os"
	"path/filepath"
	"time"
)

type (
	RaftNode struct {
		*raft.Raft

		localAddr raft.ServerAddress
		FSM       raft.FSM
		logger    hclog.Logger

		leaderFlag *gtype.Bool
	}
)

// NewRaftNodeAndBootstrapCluster
/*
PS: 将 传参addr 作为id，所以传参中无id.

@param addr 		raft节点的地址（不能为""）
@param nodeAddrs 	raft集群所有节点的地址（至少3个）
@param dir			raft节点的数据目录
@param fsm 			不能为nil
@param logger 		(1) raft节点的日志输出
					(2) 可以为nil（将使用默认值: debug级别、输出到控制台）
*/
func NewRaftNodeAndBootstrapCluster(addr string, addrs []string, dir string, fsm raft.FSM, logger hclog.Logger) (*RaftNode, error) {
	if err := validateKit.Var(addr, "hostname_port"); err != nil {
		return nil, errorKit.Wrap(err, "param addr(%s) is invalid", addr)
	}
	if err := validateKit.Var(addrs, "unique,gte=3,dive,hostname_port"); err != nil {
		return nil, errorKit.Wrap(err, "param addrs(%s) is invalid", addrs)
	}
	if err := fileKit.AssertNotExistOrIsDir(dir); err != nil {
		return nil, err
	}
	if err := interfaceKit.AssertNotNil(fsm, "fsm"); err != nil {
		return nil, err
	}
	if logger == nil {
		logger = raftLogKit.NewLogger(&hclog.LoggerOptions{
			Name:   "RAFT",
			Level:  hclog.LevelFromString("debug"),
			Output: os.Stderr,
		})
	}

	/* (0) config */
	// Richelieu: 此处不需要配置 config.NotifyCh，用不着
	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID(addr)
	config.Logger = logger
	config.SnapshotInterval = 60 * time.Second // 默认120s

	/* (1) logStore */
	logStorePath := filepath.Join(dir, "raft-log.db")
	logStore, err := raftboltdb.NewBoltStore(logStorePath)
	if err != nil {
		return nil, err
	}

	/* (2) stableStore */
	stableStorePath := filepath.Join(dir, "raft-stable.db")
	stableStore, err := raftboltdb.NewBoltStore(stableStorePath)
	if err != nil {
		return nil, err
	}

	/* (3) snapshotStore */
	/*
		@param retain: Control how many snapshots are retained. Must be at least 1.

		e.g.
			raft.NewFileSnapshotStore(opts.dataDir, 1, os.Stderr)
		e.g.1
			raft.NewFileSnapshotStore(raftDir, 2, os.Stderr)
	*/
	snapshotStore, err := raft.NewFileSnapshotStore(dir, 2, os.Stderr)
	if err != nil {
		return nil, err
	}

	/* (4) transport */
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	/*
		@param maxPool		是一个整数，表示连接池的最大容量，用于复用TCP连接以减少开销。如果为0，则使用默认值2。
		@param timeout		是一个time.Duration类型，表示连接超时的时间，用于控制网络延迟的影响。如果为0，则使用默认值10 * time.Second。
		@param logOutput	是一个io.Writer接口，表示日志输出的目标，用于记录传输层的信息。如果为nil，则使用os.Stderr作为输出。

		e.g.
			raft.NewTCPTransport(addr, addr, 2, 5*time.Second, os.Stderr)
		e.g.1
			raft.NewTCPTransport(address.String(), address, 3, 10*time.Second, os.Stderr)
	*/
	maxPool := 3
	timeout := 10 * time.Second
	transport, err := raft.NewTCPTransport(addr, tcpAddr, maxPool, timeout, os.Stderr)
	if err != nil {
		return nil, err
	}

	/* (5) raft node */
	r, err := raft.NewRaft(config, fsm, logStore, stableStore, snapshotStore, transport)
	if err != nil {
		return nil, err
	}
	node := &RaftNode{
		Raft:       r,
		FSM:        fsm,
		logger:     logger,
		leaderFlag: atomicKit.NewBool(false),
	}
	// 监听leader变化（使用此方法无法保证强一致性读，仅做leader变化过程观察）
	go func() {
		for leaderFlag := range node.LeaderCh() {
			currentLeader, _ := node.LeaderWithID()

			if leaderFlag {
				logger.Info(fmt.Sprintf("Be leader and current leader is [%s].", currentLeader))
			} else {
				logger.Warn(fmt.Sprintf("Lose leader and current leader is [%s].", currentLeader))
			}
			node.leaderFlag.Set(leaderFlag)
		}
	}()

	/* (6) Bootstrap */
	var configuration raft.Configuration
	for _, addr := range addrs {
		server := raft.Server{
			ID:      raft.ServerID(addr),
			Address: raft.ServerAddress(addr),
		}
		configuration.Servers = append(configuration.Servers, server)
	}
	node.BootstrapCluster(configuration)

	return node, nil
}

func (node *RaftNode) IsLeader() bool {
	leaderAddr, _ := node.LeaderWithID()
	return node.localAddr == leaderAddr
}

// IsLeader1 当前Raft节点是否是 Leader ？
/*
	监听leader变化（使用此方法无法保证强一致性读，仅做leader变化过程观察）
*/
func (node *RaftNode) IsLeader1() bool {
	return node.leaderFlag.Val()
}
