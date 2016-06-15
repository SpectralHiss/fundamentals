package bst

import (
	"math"

	"github.com/parallelKiller/fundamentals/datastructures/tree"
)

type binarySTree struct {
	left  *binarySTree
	right *binarySTree
	head  tree.Node
}

type BST interface {
	tree.Tree
	Search(node tree.Node) BST
	Left() BST
	Right() BST
	Head() tree.Node
}

func New(treeHead tree.Node) *binarySTree {
	return &binarySTree{left: nil, right: nil, head: treeHead}
}

func (b *binarySTree) Add(node tree.Node) {

	leaf := &binarySTree{
		left:  nil,
		right: nil,
		head:  node,
	}

	if node.K <= b.Head().K {

		if b.Left() != (*binarySTree)(nil) {
			b.Left().Add(node)
		} else {
			b.left = leaf
		}
	} else {
		if b.Right() != (*binarySTree)(nil) {
			b.Right().Add(node)
		} else {
			b.right = leaf
		}
	}

}

func (b *binarySTree) dfs(f func(tree.Node)) {

	if b.Leaf() {
		f(b.Head())
		return
	}

	if b.Left() != (*binarySTree)(nil) {
		b.Left().(*binarySTree).dfs(f)
	}

	f(b.Head())

	if b.Right() != (*binarySTree)(nil) {
		b.Right().(*binarySTree).dfs(f)
	}

}

func (b *binarySTree) Height() int {
	if b.Leaf() {
		return 0
	}

	leftHeight := b.Left().Height()
	rightHeight := b.Right().Height()

	if b.Left() == (*binarySTree)(nil) {
		return 1 + rightHeight
	}

	if b.Right() == (*binarySTree)(nil) {
		return 1 + leftHeight
	}

	return 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))
}

func (b *binarySTree) Left() BST {
	return b.left
}

func (b *binarySTree) Right() BST {
	return b.right
}

func (b *binarySTree) Head() tree.Node {
	return b.head
}

func (b *binarySTree) Leaf() bool {
	return b.right == nil && b.left == nil
}

func (b *binarySTree) Flatten() []tree.Node {
	res := []tree.Node{}

	b.dfs(func(node tree.Node) {
		res = append(res, node)
	})
	return res
}

func (b *binarySTree) Remove(node tree.Node) {

	elementTreeHead := b.Search(node)

	if elementTreeHead.Leaf() {
		elementTreeHead = nil
	} else if elementTreeHead.(*binarySTree).Left() == (*binarySTree)(nil) {
		elementTreeHead = elementTreeHead.Right()
	} else if elementTreeHead.(*binarySTree).Right() == (*binarySTree)(nil) {
		elementTreeHead = elementTreeHead.Left()
	} else {

	}

}

func (b *binarySTree) Search(node tree.Node) BST {
	nodeWeight := node.K
	headWeight := b.Head().K

	if nodeWeight == headWeight {
		return b
	} else if nodeWeight <= headWeight {
		if b.left == nil {
			return nil
		} else {
			return b.Left().Search(node)
		}
	} else {
		if b.right == nil {
			return nil
		} else {
			return b.Right().Search(node)
		}
	}
}
