package dataSizeKit

var (
	// B 1
	B = &Unit{
		value: 1,
		str:   "B",
	}
	// KB 1024
	KB = &Unit{
		value: 1 << 10,
		str:   "KB",
	}
	// MB 1048576
	MB = &Unit{
		value: 1 << 20,
		str:   "MB",
	}
	// GB 1073741824
	GB = &Unit{
		value: 1 << 30,
		str:   "GB",
	}
	// TB 1099511627776
	TB = &Unit{
		value: 1 << 40,
		str:   "TB",
	}
)

type Unit struct {
	value uint64
	str   string
}

func (su *Unit) GetValue() uint64 {
	if su == nil {
		return B.value
	}
	return su.value
}
func (su *Unit) String() string {
	if su == nil {
		return B.str
	}
	return su.str
}
