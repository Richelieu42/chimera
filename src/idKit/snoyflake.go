package idKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/sony/sonyflake"
	"time"
)

var (
	// Decompose returns a set of Sonyflake ID parts.
	Decompose func(id uint64) map[string]uint64 = sonyflake.Decompose
)

// NewSonyFlake 雪花算法.
/*
PS:
(1) 可用作分布式唯一id（前提是合理配置好 MachineID）;
(2) 通过 sonyflake.Sonyflake.NextID() 生成id（貌似是18位的）.

golang实现的雪花算法 https://mp.weixin.qq.com/s/visG_GHtU67xCtsvvG1aPQ

@param settings 可以为nil（但不推荐这么干）
*/
func NewSonyFlake(settings *sonyflake.Settings) (*sonyflake.Sonyflake, error) {
	if settings == nil {
		settings = &sonyflake.Settings{
			StartTime:      time.Now(),
			MachineID:      nil,
			CheckMachineID: nil,
		}
	}

	sf := sonyflake.NewSonyflake(*settings)
	if sf == nil {
		return nil, errorKit.New("sf == nil")
	}
	return sf, nil
}
