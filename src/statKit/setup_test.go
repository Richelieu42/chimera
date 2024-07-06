package statKit

import (
	"testing"
)

func TestMustSetup(t *testing.T) {
	MustSetup("_stat.log")

	/* 0 此时进程的CPU使用率很低 */
	//engine := gin.Default()
	//if err := engine.Run(":8888"); err != nil {
	//	panic(engine)
	//}

	/* 1 此时进程的CPU使用率很低 */
	//select {}

	/* 2 此时进程的CPU使用率很高，维持在99%左右 */
	for {
	}
}
