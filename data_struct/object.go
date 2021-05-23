package data_struct

// RiderObject 是对特定类型的数据的包装
type RiderObject struct {
	ObjectType int
	Ptr        interface{}
}

const (
	ObjectTypeString = iota + 1
	ObjectTypeInt
	ObjectTypeDouble
)


// CreateObject 创建特定类型的object结构
func CreateObject(t int, ptr interface{}) (o *RiderObject) {
	o = new(RiderObject)
	o.ObjectType = t
	o.Ptr = ptr
	return
}
