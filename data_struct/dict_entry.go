package data_struct

//哈希表节点结构(键值对)

type DictEntry struct {

	Key   interface{} //健

	Value interface{} //值

	Next  *DictEntry //下一个entry，用于解决hash冲突
}

