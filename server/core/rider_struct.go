package core

//
////服务端处理每一个链接的结构
//type Client struct {
//
//}
//
////服务端对象
//type Server struct {
//	Dbs []*DataBase
//	DbNum int //默认设置为16 如果修改，需要修改配置文件，本次暂不加入此功能
//	Start int
//	Port  int //端口
//	RdbFilename      string //暂留
//	AofFilename      string //暂留
//	NextClientID     int32
//	SystemMemorySize int32
//	Clients          int32
//	Pid              int
//	Commands   map[string]*RiderCommand
//	Dirty            int64
//	AofBuf           []string // 暂留
//}
//
//type DataBase struct {
//	HashDb *RiderDb
//	ObjectDb *RiderDb
//	ChainDb *RiderDb
//	SetDb *RiderDb
//	ZetDb *RiderDb
//}
//
///*
//一共提供五种数据结构
//分别包括普通类型hash
//hash对象
//链表
//集合
//排序集合
//*/
//const (
//	HASH = "data_struct"
//	OBJECT = "hash_object"
//	CHAIN = "linked_chain"
//	SET = "hash_set"
//	ZET = "sort_set"
//)
//type RiderDb struct {
//	Dict dict
//	DictType string //取值必须为上述五种
//	ID   int32 //DB分区id 取值 < 16
//}
//
//type dict map[string]interface{}
//
//
////命令行结构
//type RiderCommand struct {
//	Name string
//	Proc cmdFunc
//}
//
////命令函数指针
//type cmdFunc func(c *Client, s *Server)