package data_struct

//hashtable

type dict struct {
	EntryBucket [][]DictEntry

	Size uint64

	SizeMask uint64

	Used    uint64
}
