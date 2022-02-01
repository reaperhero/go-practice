package avl

import "log"

type AVLTreeNode struct {
	value int
	high  int
	left  *AVLTreeNode
	right *AVLTreeNode
}

func NewAVLTreeRoot(root int) *AVLTreeNode {
	return &AVLTreeNode{root, 0, nil, nil}
}

func (this *AVLTreeNode) InsertNode(v int) *AVLTreeNode {
	if this == nil {
		return &AVLTreeNode{v, 0, nil, nil}
	}
	if v < this.value {
		this.left = this.left.InsertNode(v)
		this.high = getMax(this.left.getNodeHigh(), this.right.getNodeHigh()) + 1
		if this.left.getNodeHigh()-this.right.getNodeHigh() == 2 {
			if v < this.left.value {
				return this.rightRotation()
			} else {
				return this.leftRightRotation()
			}
		}
	} else {
		this.right = this.right.InsertNode(v)
		this.high = getMax(this.left.getNodeHigh(), this.right.getNodeHigh()) + 1
		if this.right.getNodeHigh()-this.left.getNodeHigh() == 2 {
			if v < this.right.value {
				return this.rightLeftRotation()
			} else {
				return this.leftRotation()
			}
		}
	}
	return this
}

func (this *AVLTreeNode) leftRotation() *AVLTreeNode {
	node := this.right
	this.right = node.left
	node.left = this
	this.high = getMax(this.left.getNodeHigh(), this.right.getNodeHigh()) + 1
	node.high = getMax(node.left.getNodeHigh(), node.right.getNodeHigh()) + 1
	return node
}

func (this *AVLTreeNode) rightRotation() *AVLTreeNode {
	node := this.left
	this.left = node.right
	node.right = this
	this.high = getMax(this.left.getNodeHigh(), this.right.getNodeHigh()) + 1
	node.high = getMax(node.left.getNodeHigh(), node.right.getNodeHigh()) + 1
	return node
}

func (this *AVLTreeNode) leftRightRotation() *AVLTreeNode {
	this.left = this.left.leftRotation()
	return this.rightRotation()
}

func (this *AVLTreeNode) rightLeftRotation() *AVLTreeNode {
	this.right = this.right.rightRotation()
	return this.leftRotation()
}

func (this *AVLTreeNode) getNodeHigh() int {
	if this == nil {
		return -1
	}
	return this.high
}

func (this *AVLTreeNode) PrintSortTree() {
	if this == nil {
		return
	}
	this.left.PrintSortTree()
	log.Println(this.value)
	this.right.PrintSortTree()
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
