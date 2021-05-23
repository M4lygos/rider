package main

import (
	"Rider/data_struct"
	"Rider/data_struct/list"
	"Rider/data_struct/set"
	"Rider/server/core"
	"log"
	"net"
	"os"
	"time"
)
//riderServerInstance
var server *core.Server

func initServer() {
	server.Pid = os.Getpid()
	server.DbNum = 16


	initDb()


	server.Start = time.Now().UnixNano() / 1000000
	//server.AofFilename = DefaultAofFile
	initCommand()

	//LoadData()
}

func initCommand() {
	server.Commands = make(map[string]*core.RiderCommand)
	initStringCommand()

	initLinkedListCommand()

	initOnCommand()

	initSetCommand()
}

func initStringCommand()  {
	getCommand := &core.RiderCommand{Name: "get", Proc: core.GetCommand}
	setCommand := &core.RiderCommand{Name: "set", Proc: core.SetCommand}
	server.Commands["get"] = getCommand
	server.Commands["set"] = setCommand
}

func initLinkedListCommand()  {
	LPushCommand := &core.RiderCommand{Name: "lpush", Proc: core.LPushCommand}
	LPopCommand :=  &core.RiderCommand{Name: "lpop", Proc: core.LPopCommand}
	RPushCommand := &core.RiderCommand{Name: "rpush", Proc: core.RPushCommand}
	RPopCommand :=  &core.RiderCommand{Name: "rpop", Proc: core.RPopCommand}
	server.Commands["lpush"] = LPushCommand
	server.Commands["lpop"] = LPopCommand
	server.Commands["rpush"] = RPushCommand
	server.Commands["rpop"] = RPopCommand
}

func initSetCommand()  {
	SAddCommand := &core.RiderCommand{Name: "sadd", Proc: core.SAddCommand}
	SRemoveCommand := &core.RiderCommand{Name: "sremove", Proc: core.SRemoveCommand}
	server.Commands["sadd"] = SAddCommand
	server.Commands["sremove"] = SRemoveCommand
}

func initOnCommand()  {
	CountKeyCommand := &core.RiderCommand{Name: "count", Proc: core.CountKeyCommand}
	server.Commands["count"] = CountKeyCommand
}



func initDb() {
	server.Db = make([]*core.RiderDb, server.DbNum)
	for i := 0; i < server.DbNum; i++ {
		server.Db[i] = new(core.RiderDb)
		server.Db[i].Dict = make(map[string]*data_struct.RiderObject, 100)
		server.Db[i].ListDict = make(map[string]*list.List, 100)
		server.Db[i].SetDict = make(map[string]*set.Set, 100)
	}
}

func main() {
	//初始化实例
	server = new(core.Server)
	initServer()


	tcpListen, err := net.Listen("tcp","127.0.0.1:6171")
	if err != nil {
		log.Fatal("listen err")
	}
	defer tcpListen.Close()

	for  {
		conn, err := tcpListen.Accept()

		if err != nil {
			log.Printf("connection accept error! info is %s\n", err.Error())
			continue
		}
		log.Printf("%v has been connected", conn.RemoteAddr())
		go Handle(conn)
	}
}


func Handle(conn net.Conn)  {

	remoteIp := conn.RemoteAddr()
	client := server.CreateClient() //创建客户端实例
	defer func() {
		conn.Close()
		log.Printf( "[%v] is disconncted\n", remoteIp.String())
	}()
	for  {
		//messageLength, err := conn.Read(buffer)
		err := client.ReadQueryFromClient(conn) //读取数据
		if err != nil {
			log.Printf("client message has failed , the information is %s, the message length is %v\n", err.Error(), client.Argc)
			return
		}
		err = client.ProcessInputBuffer() //反序列化数据
		if err != nil {
			log.Printf("ProcessInputBuffer err, information is %s", err.Error())
			return
		}
		server.ProcessCommand(client) //执行命令 并将回复内容序列化
		responseConn(conn, client) //回复客户端
	}
}

func responseConn(conn net.Conn, c *core.Client) {
	conn.Write([]byte(c.Buf))
}