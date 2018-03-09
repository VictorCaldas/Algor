package main

import (
	"fmt"
)

type TreeNode struct {
	data   float64
	color  string //More than two color attributes of a binary search tree
	lchild *TreeNode
	rchild *TreeNode
	parent *TreeNode
}

type RBTree struct {
	root   *TreeNode
	cur    *TreeNode
	create *TreeNode
}

func (rbt *RBTree) Add(data float64) {
	rbt.create = new(TreeNode)
	rbt.create.data = data
	rbt.create.color = "red"

	if !rbt.IsEmpty() {
		rbt.cur = rbt.root
		for {
			if data <rbt.cur.data {
				//If you want to insert the value than the value of the current node is small, then the current node to the current node left child, if
				//The left child is empty, insert a new value in the left child
				if rbt.cur.lchild == nil {
					rbt.cur.lchild = rbt.create
					rbt.create.parent = rbt.cur
					break
				} else {
					rbt.cur = rbt.cur.lchild
				}

			} else if data > rbt.cur.data {
				//If you want to insert the value than the value of the current node, the current node to the current node right child, if
				//The right child is empty, insert a new value in the right child
				if rbt.cur.rchild == nil {
					rbt.cur.rchild = rbt.create
					rbt.create.parent = rbt.cur
					break
				} else {
					rbt.cur = rbt.cur.rchild
				}

			} else {
				//If you want to insert the value already exists in the tree, then exit
				return
			}
		}

	} else {
		rbt.root = rbt.create
		rbt.root.color = "black"
		rbt.root.parent = nil
		return
	}

	//Insert node after the repair of the property
	rbt.insertBalanceFixup(rbt.create)
}

func (rbt *RBTree) Delete(data float64) {
	var (
		deleteNode func(node *TreeNode)
		node       *TreeNode = rbt.Search(data)
	)

	deleteNode = func(node *TreeNode) {
		if node.lchild == nil && node.rchild == nil {
			//If the node to delete no child, just delete it (no miss~.~!)
			if node.parent.lchild == node {
				node.parent.lchild = nil
			} else {
				node.parent.rchild = nil
			}

		} else if node.lchild != nil && node.rchild == nil {
			//If the node to delete only the left or the right child, for the parent node of this node pointer pointing to it
			//The child can
			node.lchild.parent = node.parent
			if node.parent.lchild == node {
				node.parent.lchild = node.lchild
			} else {
				node.parent.rchild = node.lchild
			}

		} else if node.lchild == nil && node.rchild != nil {
			node.rchild.parent = node.parent
			if node.parent.lchild == node {
				node.parent.lchild = node.rchild
			} else {
				node.parent.rchild = node.rchild
			}

		} else {
			//If the node to delete both the left and the right child, the node's direct successor to the value of this festival
			//Point, and then delete the direct successor nodes can be
			successor := rbt.GetSuccessor(node.data)
			node.data = successor.data
			deleteNode(successor)
		}
	}

	deleteNode(node)
}

//This function is used to perform inserts at the red black tree, the nature of the repair
func (rbt *RBTree) insertBalanceFixup(insertnode *TreeNode) {
	var uncle *TreeNode

	for insertnode.color == "red" && insertnode.parent.color == "red" {
		//Uncle node acquires node insert new (and the parent node with another node)
		if insertnode.parent == insertnode.parent.parent.lchild {
			uncle = insertnode.parent.parent.rchild
		} else {
			uncle = insertnode.parent.parent.lchild
		}

		if uncle != nil && uncle.color == "red" {
			//If Uncle node is red, according as shown in the following image change (-> black, red, ->):
			//     |                        |
			//    1●                       1○ <-new node ptr come here
			//    / \       --------\      / \
			//  2○   ○3     --------/    2●   ●3
			//  /                        /
			//4○ <-new node ptr        4○
			//
			//This situation could have been circulating, know new node PTR refers to exit to the root (root color is black)
			uncle.color, insertnode.parent.color = "black", "black"
			insertnode = insertnode.parent.parent
			if insertnode == rbt.root || insertnode == nil {
				return
			}
			insertnode.color = "red"

		} else {
			//If the node is empty or uncle uncle node is black, in accordance with the changes as shown below:
			//     |                        |
			//    1● <-right rotate        2●
			//    / \       --------\      / \
			//  2○   ●3     --------/    4○   ○1
			//  /                              \
			//4○ <-new node ptr                 ●3
			//
			//Of course, this is just a black uncle node when the node 4, if 2 nodes of the right child, then
			//Can be left at the 2 node rotation, this into the above this kind of situation. The other two cases you want to
			//Want to understand
			if insertnode.parent == insertnode.parent.parent.lchild {
				if insertnode == insertnode.rchild {
					insertnode = insertnode.parent
					rbt.LeftRotate(insertnode)
				}
				insertnode = insertnode.parent
				insertnode.color = "black"
				insertnode = insertnode.parent
				insertnode.color = "red"
				rbt.RightRotate(insertnode)

			} else {
				if insertnode == insertnode.lchild {
					insertnode = insertnode.parent
					rbt.RightRotate(insertnode)
				}
				insertnode = insertnode.parent
				insertnode.color = "black"
				insertnode = insertnode.parent
				insertnode.color = "red"
				rbt.LeftRotate(insertnode)
			}
			return
		}
	}

}

//This function is used to delete the red black tree, repair the nature
func (rbt *RBTree) deleteBalanceFixup() {
	//Not later
}

func (rbt RBTree) GetRoot() *TreeNode {
	if rbt.root != nil {
		return rbt.root
	}
	return nil
}

func (rbt RBTree) IsEmpty() bool {
	if rbt.root == nil {
		return true
	}
	return false
}

func (rbt RBTree) InOrderTravel() {
	var inOrderTravel func(node *TreeNode)

	inOrderTravel = func(node *TreeNode) {
		if node != nil {
			inOrderTravel(node.lchild)
			fmt.Printf("%g ", node.data)
			inOrderTravel(node.rchild)
		}
	}

	inOrderTravel(rbt.root)
}

func (rbt RBTree) Search(data float64) *TreeNode {
	//And Add operations are similar, as long as the current node to the left than small children turn on, turn right on thinking on children is bigger than the other
	//Go all the way to find, know to find to find the value to return
	rbt.cur = rbt.root
	for {
		if data <rbt.cur.data {
			rbt.cur = rbt.cur.lchild
		} else if data > rbt.cur.data {
			rbt.cur = rbt.cur.rchild
		} else {
			return rbt.cur
		}

		if rbt.cur == nil {
			return nil
		}
	}
}

func (rbt RBTree) GetDeepth() int {
	var getDeepth func(node *TreeNode) int

	getDeepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.lchild == nil && node.rchild == nil {
			return 1
		}
		var (
			ldeepth int = getDeepth(node.lchild)
			rdeepth int = getDeepth(node.rchild)
		)
		if ldeepth > rdeepth {
			return ldeepth + 1
		} else {
			return rdeepth + 1
		}
	}

	return getDeepth(rbt.root)
}

func (rbt RBTree) GetMin() float64 {
	//According to the properties of two binary search tree, node in the tree is the smallest node
	if rbt.root == nil {
		return -1
	}
	rbt.cur = rbt.root
	for {
		if rbt.cur.lchild != nil {
			rbt.cur = rbt.cur.lchild
		} else {
			return rbt.cur.data
		}
	}
}

func (rbt RBTree) GetMax() float64 {
	//According to the properties of two binary search tree, node in the tree is the right value of the largest node
	if rbt.root == nil {
		return -1
	}
	rbt.cur = rbt.root
	for {
		if rbt.cur.rchild != nil {
			rbt.cur = rbt.cur.rchild
		} else {
			return rbt.cur.data
		}
	}
}

func (rbt RBTree) GetPredecessor(data float64) *TreeNode {
	getMax := func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		for {
			if node.rchild != nil {
				node = node.rchild
			} else {
				return node
			}
		}
	}

	node := rbt.Search(data)
	if node != nil {
		if node.lchild != nil {
			//If the node has a left child, then the right most direct predecessor it is its left sub tree nodes, because the
			//A node value small node in the left subtree, and these nodes median maximum is the rightmost node
			return getMax(node.lchild)
		} else {
			//If the node has no children left, then along its parent node to find, know a father node parent's right
			//The child is the father node, then the parent node parent node is the direct precursor
			for {
				if node == nil || node.parent == nil {
					break
				}
				if node == node.parent.rchild {
					return node.parent
				}
				node = node.parent
			}
		}
	}

	return nil
}

func (rbt RBTree) GetSuccessor(data float64) *TreeNode {
	getMin := func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		for {
			if node.lchild != nil {
				node = node.lchild
			} else {
				return node
			}
		}
	}

	//The reference function to find the direct precursor of the watch
	node := rbt.Search(data)
	if node != nil {
		if node.rchild != nil {
			return getMin(node.rchild)
		} else {
			for {
				if node == nil || node.parent == nil {
					break
				}
				if node == node.parent.lchild {
					return node.parent
				}
				node = node.parent
			}
		}
	}

	return nil
}

func (rbt *RBTree) Clear() {
	rbt.root = nil
	rbt.cur = nil
	rbt.create = nil
}

/**
 * Rotation diagram (to rotate left as an example):
 *     |                               |
 *     ○ <-left rotate                 ●
 *      \              ----------\    / \
 *       ●             ----------/   ○   ●r
 *      / \                           \
 *    l●   ●r                         l●
 *
 *
 *
 *     |                               |
 *     ○ <-left rotate                 ●
 *      \              ----------\    / \
 *       ●             ----------/   ○   ●
 *        \                           \
 *         ●                          nil <-don't forget it should be nil
 */
func (rbt *RBTree) LeftRotate(node *TreeNode) {
	if node.rchild == nil {
		return
	}

	right_child := node.rchild
	//The rotation of the nodes of the right child left children assigned to the nodes of the right child, here are the best in the following 3 lines of code in order to write,
	//Otherwise, the node right child left child is nil, it is easy to forget the nodes of the right child is set to nil
	node.rchild = right_child.lchild
	if node.rchild != nil {
		node.rchild.parent = node
	}

	//Let the father node pointer to rotate the right child node to the current node parent node. If the parent node is the root node to special treatment
	right_child.parent = node.parent
	if node.parent == nil {
		rbt.root = right_child
	} else {
		if node.parent.lchild == node {
			node.parent.lchild = right_child
		} else {
			node.parent.rchild = right_child
		}
	}

	//The above preparation is completed, you can start rotating, let to rotate right child left child node to the node,
	//Don't forget to put the father node pointer rotation node point to the parent node of the new
	right_child.lchild = node
	node.parent = right_child
}

func (rbt *RBTree) RightRotate(node *TreeNode) {
	//Process and LeftRotate rotation (right) on the contrary
	if node.lchild == nil {
		return
	}

	left_child := node.lchild
	node.lchild = left_child.rchild
	if node.lchild != nil {
		node.lchild.parent = node
	}

	left_child.parent = node.parent
	if node.parent == nil {
		rbt.root = left_child
	} else {
		if node.parent.lchild == node {
			node.parent.lchild = left_child
		} else {
			node.parent.rchild = left_child
		}
	}
	left_child.rchild = node
	node.parent = left_child
}

func main() {
	var rbt RBTree
	rbt.Add(10)
	rbt.Add(8)
	rbt.Add(7)
	rbt.Add(6)
	rbt.Add(5)
	rbt.Add(4)
	rbt.Add(3)
	rbt.Add(2)
	rbt.Add(1)
	fmt.Println(rbt.GetDeepth())
}