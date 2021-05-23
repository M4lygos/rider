package core

import (
	"Rider/data_struct"
	"Rider/data_struct/list"
	"Rider/data_struct/set"
	"time"
)

const (
	ARGCOUNT_ERROR = "参数错误!"
	SUCCESS_INFORMARION = "ok"
	DATA_NOT_EXIST = "key does not exist"
	ALL_KEY = "*"
	DICT_KEY = "string"
	LIST_KEY = "list"
)


//Client 与服务端连接之后即创建一个Client结构
type Client struct {
	Cmd      *RiderCommand
	Argv     []*data_struct.RiderObject
	Argc     int
	Db       *RiderDb
	QueryBuf string
	Buf      string
	FakeFlag bool
}



type RiderCommand struct {
	Name string
	Proc cmdFunc
}

type cmdFunc func(c *Client, s *Server)

// Server 服务端实例结构体
type Server struct {
	Db               []*RiderDb
	DbNum            int
	Start            int64
	Port             int32
	RdbFilename      string
	AofFilename      string
	NextClientID     int32
	SystemMemorySize int32
	Clients          int32
	Pid              int
	Commands         map[string]*RiderCommand
	Dirty            int64
	AofBuf           []string
}

//use map[string]* as type dict
//使用Go原生数据结构map作为redis中dict结构体 暂不对dict造轮子
type dict map[string]*data_struct.RiderObject

//链表字典
type linkedListDict map[string]*list.List

//集合字典
type setDict map[string]*set.Set

//RiderDb db结构体
type RiderDb struct {
	Dict    	dict
	ListDict    linkedListDict
	SetDict     setDict
	Expires 	dict
	ID      	int32
}

//内存过期策略结构体
type Expire struct {

}

type ExpireNode struct {
	NodeFlag  string
	Key       string
	LimitTime time.Time
	next      *ExpireNode
}