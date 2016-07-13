package bst

import (
	//"fmt"
	"github.com/parallelKiller/fundamentals/datastructures/tree"
	"math"

	"github.com/davecgh/go-spew/spew"
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

	elementTreeHead := b.Search(node).(*binarySTree)

	if elementTreeHead.Leaf() {
		parent := b.FindParent(elementTreeHead)
		if parent.left == elementTreeHead {
			b.left = nil
		} else {
			b.right = nil
		}
	} else if elementTreeHead.Left() == (*binarySTree)(nil) {
		*elementTreeHead = *elementTreeHead.Right().(*binarySTree)
	} else if elementTreeHead.Right() == (*binarySTree)(nil) {
		spew.Dump(elementTreeHead)
		*elementTreeHead = *elementTreeHead.Left().(*binarySTree)
	} else {

	}

}

// this allows recursion to work it's course smoothly
func findParentRec(tree *binarySTree, elem *binarySTree, parent *binarySTree) *binarySTree {
	if tree == nil {
		return nil
	}

	println("case A")

	if tree.Head().K == elem.Head().K {
		println("Site B")
		return parent
	} else if tree.Head().K < elem.Head().K {
		println("Site C")
		return findParentRec(tree.Left().(*binarySTree), elem, tree)
	} else {
		println("site B")
		return findParentRec(tree.Right().(*binarySTree), elem, tree)
	}
}

func (b *binarySTree) FindParent(elem *binarySTree) *binarySTree {
	return findParentRec(b, elem, nil)
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
