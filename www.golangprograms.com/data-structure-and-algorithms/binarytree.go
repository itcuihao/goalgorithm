package main

// 二元树是一个有根树，并且每个节点最多有2个子节点。非空的二叉树，若树叶总数为 n0，分支度为2的总数为 n2，则 n0 = n2 + 1。
// 一棵深度为k，且有 2^(k+1)-1 个节点的二元树，称为满二元树（Full Binary Tree）。这种树的特点是每一层上的节点数都是最大节点数。而在一棵二元树中，除最后一层外，若其馀层都是满的，并且最后一层或者是满的，或者是在右边缺少连续若干节点，则此二元树为完全二元树（Complete Binary Tree）。具有n个节点的完全二元树的深度为 log_2n+1。深度为k的完全二元树，至少有 2^k 个节点，至多有 2^(k+1)-1个节点。

// 遍历二叉树：L、D、R分别表示遍历左子树、访问根结点和遍历右子树，则先(根)序遍历二叉树的顺序是DLR，中(根)序遍历二叉树的顺序是LDR，后(根)序遍历二叉树的顺序是LRD。还有按层遍历二叉树。这些方法的时间复杂度都是O(n)，n为结点个数。

// 如果T2是由有序树T转换而来的二叉树，那么T中结点的前序就是T2中结点的前序，T中结点的后序就是T2中结点的中序。任何一棵二叉树的叶结点在先序、中序和后序遍历中的相对次序不发改变。设n,m为一棵二叉树上的两个结点，在中序遍历时，n在m前的条件是n在m的左方。前序序列和中序序列相同的二叉树为空树或任一结点均无左孩子的非空二叉树；中序序列和后序序列相同的二叉树为空树或任一结点均无右孩子的非空二叉树；前序序列和后序序列相同的二叉树为空树或仅有一个结点的二叉树。
import (
	"fmt"
	"io"
	"os"
)

// BinaryNode 节点
type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

// BinaryTree 树
type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(data int64) *BinaryTree {
	if t.root == nil {
		// root节点为nil
		t.root = &BinaryNode{
			data:  data,
			left:  nil,
			right: nil,
		}
	} else {
		// root节点不为空构造二叉树
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data int64) {
	if n == nil {
		return
	} else if data <= n.data { //判断左右节点
		if n.left == nil {
			// 左节点为nil
			n.left = &BinaryNode{
				data:  data,
				left:  nil,
				right: nil,
			}
		} else {
			// 左节点不为nil，继续构造二叉树
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			// 右节点为nil
			n.right = &BinaryNode{
				data:  data,
				left:  nil,
				right: nil,
			}
		} else {
			// 右节点不为nil，继续构造二叉树
			n.right.insert(data)
		}
	}
}

func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func main() {
	tree := &BinaryTree{}
	tree.insert(1).
		insert(2).
		insert(3)
	print(os.Stdout, tree.root, 0, 'M')
}
