package gzipKit

import (
	"fmt"
	"testing"
)

func TestCompressComplexly(t *testing.T) {
	json := `'"{\\n    \\"errno\\": 0,\\n    \\"errmsg\\": \\"ok\\",\\n    \\"data\\": {\\n        \\"log_id\\": \\"3754236822\\",\\n        \\"action_rule\\": {\\n            \\"pos_1\\": [],\\n            \\"pos_2\\": [],\\n            \\"pos_3\\": []\\n        }\\n    }\\n}"'`
	fmt.Println("before length:", len(json))
	fmt.Println("------")

	{
		// 压缩
		data, err := CompressComplexly([]byte(json), WithCompressThreshold(-1))
		if err != nil {
			panic(err)
		}
		fmt.Println("after length:", len(data))
		fmt.Println(string(data))
		fmt.Println("------")
	}

	{
		// 不压缩（266 < 300）
		data, err := CompressComplexly([]byte(json), WithCompressThreshold(300))
		if err != nil {
			panic(err)
		}
		fmt.Println("after length:", len(data))
		fmt.Println(string(data))
		fmt.Println("------")
	}
}
