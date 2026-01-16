package bloom

import "hash"

type bloomFilter struct {
	bitMap []uint32
	k      int
	m      int
	n      int
	hashfn hash.Hash64
}

func NewBloomFilter() *bloomFilter {
	return &bloomFilter{}
}

func (bf *bloomFilter) Add(element []byte) {}

func (bf *bloomFilter) Check(element []byte) bool {
	return false
}
