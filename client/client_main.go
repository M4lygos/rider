package main

import (
	proto "Rider/protocol"
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

const remoteAddress = "127.0.0.1:6171"

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	log.Println("Hi Rider")
	tcpAddr, err := net.ResolveTCPAddr("tcp4", remoteAddress)
	if err != nil {
		log.Printf("resolve tcp error, message is %s \n", err.Error())
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Printf("dial tcp error, message is %s \n", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	responseBuffer := make([]byte, 1024)
	for  {
		log.Printf("%s>\n", remoteAddress)
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\n", "", -1)
		send2Server(text, conn)
		num, err := conn.Read(responseBuffer)

		resp, er := proto.DecodeFromBytes(responseBuffer)
		if err != nil {
			log.Printf("get response err message is %s", err.Error())
			break
		}
		if num == 0 {
			log.Printf("%s> nil\n", remoteAddress)
		}else if er != nil {
			log.Printf("%s> err server response, err is %s \n",remoteAddress, er.Error() )
			continue
		}
		log.Println(string(resp.Value))
	}
}

func send2Server(text string, conn *net.TCPConn) (int, error) {
	//n, err := conn.Write([]byte(text))
	messageByte, err := proto.EncodeCmd(text)
	if err != nil {
		return 0, err
	}
	n, err := conn.Write(messageByte)
	if err != nil {
		log.Printf("the message has been send failed please check the information [%s] \n", err.Error())
		return 0, err
	}
	return n, err
}

func checkError(err error) {
	if err != nil {
		log.Println("err ", err.Error())
		os.Exit(1)
	}
}