package data_struct

import (
	"fmt"
	"testing"
)

func TestHashObject(t *testing.T) {
	o1 := CreateObject(1, 10)
	o2 := CreateObject(1, 10)
	o3 := CreateObject(1, "xty")
	o4 := CreateObject(1, "xty")
	s := make(map[RiderObject]struct{})
	s[*o1] = struct{}{}
	_, ok := s[*o2]
	fmt.Println(ok)
	s[*o3] = struct{}{}
	_, ok = s[*o4]
	fmt.Println(ok)
}