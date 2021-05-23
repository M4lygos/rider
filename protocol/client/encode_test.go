package client

import (
	"log"
	"testing"
)

func Test_client(t *testing.T)  {
	log.Println(string(BulkString("set")))
	log.Println(string(BulkString("get")))
}