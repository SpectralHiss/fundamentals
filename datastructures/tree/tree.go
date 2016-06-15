package tree

type Node struct {
	K int
	V interface{}
}

type Tree interface {
	Add(node Node)
	Remove(node Node)
	Flatten() []Node
	Leaf() bool
	Height() int
}
