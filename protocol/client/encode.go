package client

import "fmt"

func BulkString(s string) []byte {
	count := len(s)
	result := fmt.Sprintf("$%d\r\n%s\r\n", count, s)
	return []byte(result)
}