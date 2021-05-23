package core

import (
	"Rider/data_struct"
	"Rider/data_struct/list"
	"Rider/data_struct/set"
	proto "Rider/protocol"
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func SetCommand(c *Client, s *Server) {
	objKey := c.Argv[1]
	objValue := c.Argv[2]
	if c.Argc != 3 {
		addReplyError(c, "(error) ERR wrong number of arguments for 'set' command")
	}
	if stringKey, ok1 := objKey.Ptr.(string); ok1 {
		if stringValue, ok2 := objValue.Ptr.(string); ok2 {
			c.Db.Dict[stringKey] = data_struct.CreateObject(data_struct.ObjectTypeString, stringValue)
		}
	}
	s.Dirty++
	addReplyStatus(c, SUCCESS_INFORMARION)
}

// GetCommand get命令实现
func GetCommand(c *Client, s *Server) {
	o := lookupKey(c.Db, c.Argv[1])
	if o != nil {
		addReplyStatus(c, o.Ptr.(string))
	} else {
		addReplyStatus(c, DATA_NOT_EXIST)
	}
}

func CountKeyCommand(client *Client, server *Server)  {
	if client.Argc != 2 {
		addReplyError(client, ARGCOUNT_ERROR)
		return
	}
	flag := client.Argv[1].Ptr.(string)
	var count int
	switch flag {
	case ALL_KEY:
		count = CountOfKey(client.Db)
	case DICT_KEY:
		count = 12312
	case LIST_KEY:
		count = 17
	}
	addReplyStatus(client, strconv.Itoa(count))
}

func LPushCommand(client *Client, server *Server)  {
	if client.Argc < 3 {
		addReplyError(client, ARGCOUNT_ERROR)
		return
	}
	LinkedList := lookupLinkedKey(client.Db, client.Argv[1])

	if LinkedList == nil {
		LinkedList = list.NewList()
		client.Db.ListDict[client.Argv[1].Ptr.(string)] = LinkedList
		server.Dirty ++
	}
	LinkedList.InsertHead(client.Argv[2])
	addReplyStatus(client, SUCCESS_INFORMARION)
}

func LPopCommand(client *Client, server *Server)  {
	if client.Argc !=  2 {
		addReplyError(client, ARGCOUNT_ERROR)
		return
	}
	linkedList := lookupLinkedKey(client.Db, client.Argv[1])
	if linkedList == nil {
		addReplyError(client, DATA_NOT_EXIST)
		return
	}
	node, err := linkedList.DeleteHead()
	if err != nil {
		addReplyError(client, err.Error())
		return
	}
	addReplyStatus(client, node.Val.Ptr.(string))
}

func RPushCommand(client *Client, server *Server)  {
	if client.Argc < 3 {
		addReplyError(client, ARGCOUNT_ERROR)
		return
	}
	LinkedList := lookupLinkedKey(client.Db, client.Argv[1])
	if LinkedList == nil {
		LinkedList = list.NewList()
		client.Db.ListDict[client.Argv[1].Ptr.(string)] = LinkedList
		server.Dirty ++
	}
	LinkedList.InsertTail(client.Argv[2])
	addReplyStatus(client, SUCCESS_INFORMARION)
}

func RPopCommand(client *Client, server *Server)  {
	if client.Argc !=  2 {
		addReplyError(client, ARGCOUNT_ERROR)
		return
	}
	linkedList := lookupLinkedKey(client.Db, client.Argv[1])
	if linkedList == nil {
		addReplyError(client, DATA_NOT_EXIST)
		return
	}
	node, err := linkedList.DeleteTail()
	if err != nil {
		addReplyError(client, err.Error())
		return
	}
	addReplyStatus(client, node.Val.Ptr.(string))
}

func SAddCommand(client *Client, server *Server)  {
	if client.Argc < 3 {
		addReplyError(client, ARGCOUNT_ERROR)
		return
	}
	hashSet := lookupSetKey(client.Db, client.Argv[1])
	if hashSet == nil {
		hashSet = set.NewSet(client.Argv[1].Ptr.(string))
		client.Db.SetDict[client.Argv[1].Ptr.(string)] = hashSet
		server.Dirty ++
	}
	flag := hashSet.Add(client.Argv[2])
	addReplyStatus(client, flag)
}

func SRemoveCommand(client *Client, server *Server)  {
	if client.Argc < 3 {
		addReplyError(client, ARGCOUNT_ERROR)
		return
	}
	hashSet := lookupSetKey(client.Db, client.Argv[1])
	if hashSet == nil {
		addReplyError(client, DATA_NOT_EXIST)
		return
	}
	flag := hashSet.Delete(client.Argv[2])
	addReplyStatus(client, flag)
}

// addReply 添加回复
func addReply(c *Client, o *data_struct.RiderObject) {
	c.Buf = o.Ptr.(string)
}

func addReplyStatus(c *Client, s string) {
	r := proto.NewString([]byte(s))
	addReplyString(c, r)
}

func addReplyError(c *Client, s string) {
	r := proto.NewError([]byte(s))
	addReplyString(c, r)
}

func addReplyString(c *Client, r *proto.Resp) {
	if ret, err := proto.EncodeToBytes(r); err == nil {
		c.Buf = string(ret)
	}
}

// ProcessCommand 执行命令
func (s *Server) ProcessCommand(c *Client) {
	v := c.Argv[0].Ptr
	name, ok := v.(string)
	if !ok {
		log.Println("error cmd")
		os.Exit(1)
	}
	cmd := lookupCommand(name, s)
	if cmd != nil {
		c.Cmd = cmd
		call(c, s)
	} else {
		addReplyError(c, fmt.Sprintf("(error) ERR unknown command '%s'", name))
	}
}

// lookupCommand查找命令
func lookupCommand(name string, s *Server) *RiderCommand {
	if cmd, ok := s.Commands[name]; ok {
		return cmd
	}
	return nil
}

// call 真正调用命令
func call(c *Client, s *Server) {
	dirty := s.Dirty
	c.Cmd.Proc(c, s)
	dirty = s.Dirty - dirty
	if dirty > 0 && !c.FakeFlag {
		//AppendToFile(s.AofFilename, c.QueryBuf)
	}
}
func lookupKey(db *RiderDb, key *data_struct.RiderObject) (ret *data_struct.RiderObject) {
	if o, ok := db.Dict[key.Ptr.(string)]; ok {
		return o
	}
	return nil
}

func CountOfKey(db *RiderDb) int {
	return len(db.Dict)
}

func lookupLinkedKey(db *RiderDb, key *data_struct.RiderObject) *list.List {
	if o, ok := db.ListDict[key.Ptr.(string)]; ok {
		return o
	}
	return nil
}

func lookupSetKey(db *RiderDb, key *data_struct.RiderObject) *set.Set {
	if o, ok := db.SetDict[key.Ptr.(string)]; ok {
		return o
	}
	return nil
}



// CreateClient 连接建立 创建client记录当前连接
func (s *Server) CreateClient() (c *Client) {
	c = new(Client)
	c.Db = s.Db[0]
	c.QueryBuf = ""
	return c
}

// ReadQueryFromClient 读取客户端请求信息
func (c *Client) ReadQueryFromClient(conn net.Conn) (err error) {
	buff := make([]byte, 512)
	n, err := conn.Read(buff)
	if err != nil {
		log.Println("conn.Read err!=nil", err, "---len---", n, conn)
		conn.Close()
		return err
	}
	c.QueryBuf = string(buff)
	return nil
}

// ProcessInputBuffer 处理客户端请求信息
func (c *Client) ProcessInputBuffer() error {
	//r := regexp.MustCompile("[^\\s]+")
	decoder := proto.NewDecoder(bytes.NewReader([]byte(c.QueryBuf)))
	//decoder := proto.NewDecoder(bytes.NewReader([]byte("*2\r\n$3\r\nget\r\n")))
	if resp, err := decoder.DecodeMultiBulk(); err == nil {
		c.Argc = len(resp)
		c.Argv = make([]*data_struct.RiderObject, c.Argc)
		for k, s := range resp {
			c.Argv[k] = data_struct.CreateObject(data_struct.ObjectTypeString, string(s.Value))
		}
		return nil
	}
	return errors.New("ProcessInputBuffer failed")
}
