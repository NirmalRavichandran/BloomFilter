package bloom

import (
	"hash"
	"hash/fnv"
)

type bloomFilter struct {
	bitMap []bool // bit slice
	k      int    // number of hash functions
	m      int    // size of the filter
	n      int    // number of elements in the filter
	hashfn hash.Hash64
}

func NewBloomFilter(noHashes, filterSize int) *bloomFilter {
	bf := new(bloomFilter)
	bf.bitMap = make([]bool, filterSize)
	bf.k = noHashes
	bf.m = filterSize
	bf.hashfn = fnv.New64a()
	bf.n = 0 // initial value before adding any element
	return bf
}

func (bf *bloomFilter) Add(element []byte) {
	h1, h2 := bf.getHash(element)
	for i := 0; i < bf.k; i++ {
		index := (h1 + uint32(i)*h2) % uint32(bf.m)
		bf.bitMap[index] = true
	}
	bf.n++
}

func (bf *bloomFilter) Check(element []byte) bool {
	h1, h2 := bf.getHash(element)
	for i := 0; i < bf.k; i++ {
		index := (h1 + uint32(i)*h2) % uint32(bf.m)
		if !bf.bitMap[index] {
			return false
		}
	}
	return true
}
