package bst

import "github.com/parallelKiller/fundamentals/datastructures/tree"

type binarySTree struct {
	Left  *binarySTree
	Right *binarySTree
	Head  tree.Node
}

type BST interface {
	tree.Tree
	Search(node tree.Node) BST
	Left() BST
	Right() BST
	Head() BST
}

func New(head tree.Node) *binarySTree {
	return &binarySTree{Left: nil, Right: nil, Head: head}
}

func (b *binarySTree) Add(node tree.Node) {
	if node.K < b.Head.K {

	}
}

func (b *binarySTree) Left() BST {
	return b.Left
}

func (b *binarySTree) Right() BST {
	return b.Right
}

func (b *binarySTree) Head() BST {
	return b.Head
}

func (b *binarySTree) Flatten() []tree.Node {
	return []tree.Node{}
}

func (b *binarySTree) Remove(node tree.Node) {}

func (b *binarySTree) Search(node tree.Node) *binarySTree {
	return nil
}
