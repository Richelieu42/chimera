package dataSizeKit

const (
	B   int64 = 1
	KiB int64 = 1 << 10
	MiB int64 = 1 << 20
	GiB int64 = 1 << 30
	TiB int64 = 1 << 40
	PiB int64 = 1 << 50
	EiB int64 = 1 << 60
	//ZB = 1 << 70
	//YB = 1 << 80
)

//var (
//	// B 1
//	B = &Unit{
//		value: 1,
//		str:   "B",
//	}
//	// KB 1024
//	KB = &Unit{
//		value: 1 << 10,
//		str:   "KB",
//	}
//	// MB 1048576
//	MB = &Unit{
//		value: 1 << 20,
//		str:   "MB",
//	}
//	// GB 1073741824
//	GB = &Unit{
//		value: 1 << 30,
//		str:   "GB",
//	}
//	// TB 1099511627776
//	TB = &Unit{
//		value: 1 << 40,
//		str:   "TB",
//	}
//)
//
//type Unit struct {
//	value uint64
//	str   string
//}
//
//func (su *Unit) GetValue() uint64 {
//	if su == nil {
//		return B.value
//	}
//	return su.value
//}
//func (su *Unit) ToDSN() string {
//	if su == nil {
//		return B.str
//	}
//	return su.str
//}
