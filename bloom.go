package bloom

type bloomFilter struct {
	// Yet to add fields
}

func NewBloomFilter() *bloomFilter {
	return &bloomFilter{}
}

func (bf *bloomFilter) Add(element []byte) {}

func (bf *bloomFilter) Check(element []byte) bool {
	return false
}
