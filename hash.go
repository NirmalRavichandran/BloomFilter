package bloom

func (bf *bloomFilter) getHash(element []byte) (uint32, uint32) {
	bf.hashfn.Reset()
	bf.hashfn.Write(element)
	hash := bf.hashfn.Sum64()
	h1 := uint32(hash & 0xFFFFFFFF)
	h2 := uint32(hash >> 32)
	return h1, h2
}
