package main

import (
	"fmt"
)

type Node struct {
	data int
	height int
	left *Node
	right *Node
}

/*
	when this function receives arguments, those argument variables have different memory address than variables that were sent
	ex. insert(root *Node) receiving variable will have a different memory address than what was sent in function call i.e insert(root)
	so making the receiving variable point to something will not alter the variable that was sent
	Two variables pointing to the same address does not mean they are the same
*/

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func height(root *Node) int {
	if root == nil {
		return 0
	} else {
		return root.height
	}
}

func leftRotate(root *Node) *Node {
	var tmp *Node
	tmp = root
	root = root.right
	tmp.right = root.left
	root.left = tmp
	tmp.height = max(height(tmp.left), height(tmp.right)) + 1
	root.height = max(height(root.left), height(root.right)) + 1
	return root
}

func rightRotate(root *Node) *Node {
	var tmp *Node
	tmp = root
	root = root.left
	tmp.left = root.right
	root.right = tmp
	tmp.height = max(height(tmp.left), height(tmp.right)) + 1
	root.height = max(height(root.left), height(root.right)) + 1
	return root
}

func insert(root *Node, data int) *Node { 
	if root == nil {
		fmt.Println("inserting:", data)
		root = &Node{data, 1, nil, nil} //take height of leaf as 1
	} else {
		if data > root.data {
			root.right = insert(root.right, data)
		} else {
			root.left = insert(root.left, data)
		}
		root.height = max(height(root.left), height(root.right)) + 1
		f := height(root.left) - height(root.right)
		fmt.Println("balance factor on node", root.data, "is", f)
		if f < -1 && data > root.right.data { //case 1: right-right case
			fmt.Println("balancing tree...right-right case")
			root = leftRotate(root)
		} else if f < -1 && data < root.right.data { //case 2: right-left case
			fmt.Println("balancing tree...right-left case")
			root.right = rightRotate(root.right)
			root = leftRotate(root)
		} else if f > 1 && data < root.left.data { //case 3: left-left case
			fmt.Println("balancing tree...left-left case")
			root = rightRotate(root)
		} else if f > 1 && data > root.left.data { //case 4: left-right case
			fmt.Println("balancing tree...left-right case")
			root.left = leftRotate(root.left)
			root = rightRotate(root)
		}
	}
	return root
}

func remove(root *Node, data int) *Node {
	//base case if root is null
	if root == nil {
		//nothing is found
	} else if data == root.data { //if data is found at current node then determine type of node
		if root.left == nil && root.right == nil {
			// case 1: if curent node is leaf node
			root = nil
		} else if root.left == nil {
			//case 2: if current node's left child is null, then make right child as the current node
			root = root.right
		} else if root.right == nil {
			//case 3: if current node's right child is null, then make left child as the current node
			root = root.left
		} else {
			//case 4: if both right and left children are present, then find the minimum value node from right subtree and replace the current node with it's value
			//and delete the minimum value node
			var x *Node
			x = getMin(root.right)
			root.data = x.data
			root.right = remove(root.right, x.data)
		}
	} else { //if data is not found at current node then traverse the tree
		root.left = remove(root.left, data)
		root.right = remove(root.right, data)
	}
	return root
}

func getMin(root *Node) *Node {
	fmt.Println(root.data)
	if root == nil {
		return nil
	} else if root.left == nil {
		return root
	} else {
		root = getMin(root.left)
	}
	return root
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	} else {
		fmt.Println("on node %d", root.data)
		ldepth := maxDepth(root.left)
		rdepth := maxDepth(root.right)

		fmt.Println("comparing %d & %d", ldepth, rdepth)
		if ldepth > rdepth {
			return ldepth+1
		} else {
			return rdepth+1
		}
	}
}

func longestPath(root *Node) []int {
	if root == nil {
		return []int{}
	} else {
		fmt.Println("on node", root.data)
		path := []int{root.data}
		lpath := longestPath(root.left)
		rpath := longestPath(root.right)

		fmt.Println("comparing", lpath, rpath)
		if len(lpath) > len(rpath) {
			return append(path, lpath...)
		} else {
			return append(path, rpath...)
		}
	}
}

func printSideWays(root *Node, indent string) {
	if root != nil {
		printSideWays(root.right, indent + "    ")
		//fmt.Println(indent, root.data, "->" ,root.height)
		fmt.Println(indent, root.data)
		printSideWays(root.left, indent + "    ")
	}
}

func printByLevel(root *Node, level int) {
	if root != nil {
		printByLevel(root.right, level+1)
		fmt.Println(level, root.data)
		printByLevel(root.left, level+1)
	}
}

func checkBalance(root *Node) int {
	if root == nil {
		return 0
	} else {
		lh := checkBalance(root.left)
		rh := checkBalance(root.right)
		fmt.Println("on node", root.data, "balance factor", lh-rh)
		if lh > rh {
			return lh+1
		} else {
			return rh+1
		}
	}
}

func main() {
	var bst *Node
	for i, num := range []int{11, 7, 16, 20, 29, 15, 14} {
		fmt.Println(i, num)
		if bst == nil {
			bst = insert(bst, 10) //first call to insert needs to return the new node created as empty bst will point to that here
		} else {
			insert(bst, num) // no need to write bst = insert(bst, num) as passed bst will retain pointers that we added as children
			printSideWays(bst, "")
		}
	}
	//fmt.Println("printing tree sideways")
	//printSideWays(bst, "")
	//fmt.Println("printing tree by level")
	//printByLevel(bst, 0)
	//remove(bst, 10)
	//printSideWays(bst, "")
	//fmt.Println("subtree height", maxDepth(bst))
	//fmt.Println("longest path", longestPath(bst))
	//balance(bst)
	//printSideWays(bst, "")
}