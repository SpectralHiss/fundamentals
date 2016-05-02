package set

type BitVector uint64

func New() *BitVector {
	set := BitVector(0)
	return &set
}

func (b *BitVector) Add(elem uint) {
	*b = *b | (1 << elem)
}

func (b *BitVector) Member(elem uint) bool {
	return (*b >> elem & 1) == 1
}

func (b *BitVector) Union(c *BitVector) *BitVector {
	newSet := BitVector(*b | *c)
	return &newSet
}

func (b *BitVector) Intersect(c *BitVector) *BitVector {
	newSet := BitVector(*b & *c)
	return &newSet
}
