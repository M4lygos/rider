package set

import "Rider/data_struct"

type Set struct {
	SetName string
	HashSet map[data_struct.RiderObject]struct{}
}

func NewSet(name string) *Set {
	return &Set{
		SetName: name,
		HashSet: make(map[data_struct.RiderObject]struct{}, 100),
	}
}

func (set *Set) Add(object *data_struct.RiderObject) string {
   if 	_, ok := set.HashSet[*object] ; ok {
   		return "fail"
   }
   set.HashSet[*object] = struct{}{}
   return "ok"
}

func (set *Set) Delete(object *data_struct.RiderObject) string {
	if 	_, ok := set.HashSet[*object] ; !ok {
		return "fail"
	}
	delete(set.HashSet, *object)
	return "ok"
}