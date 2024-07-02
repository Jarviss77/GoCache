package main

// Implemented avl trees to Insert, Delete, and Search for a node in the tree in log(n) time.

type Node struct {
	HashedKey int
	Key string
	Value string
	Height int
	Left  *Node
	Right *Node
}

type AVLTree struct {
	Root *Node
}

func (tree *AVLTree) Insert(key string, value string) {
	tree.Root = insert(tree.Root, hasher(key), key, value)
}

func insert(node *Node, hashedKey int, key string , value string ) *Node {
	// Simple BST insertion
	if node == nil {
		return &Node{HashedKey: hashedKey, Key: key, Value: value, Height: 1}
	}

	if hashedKey < node.HashedKey {
		node.Left = insert(node.Left, hashedKey, key, value)
	} else if hashedKey > node.HashedKey {
		node.Right = insert(node.Right, hashedKey, key, value)
	} else {
		node.Value = value
		return node
	}

	node.Height = max(height(node.Left), height(node.Right)) + 1

	balance := getBalance(node)

	if balance > 1 && hashedKey < node.Left.HashedKey {
		return rightRotate(node)
	} else if balance < -1 && hashedKey > node.Right.HashedKey {
		return leftRotate(node)
	} else if balance > 1 && hashedKey > node.Left.HashedKey {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	} else if balance < -1 && hashedKey < node.Right.HashedKey {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	return node
}

func rightRotate(node *Node) *Node {
	newRoot := node.Left
	temp := newRoot.Right
	newRoot.Right = node
	node.Left = temp

	node.Height = max(height(node.Left), height(node.Right)) + 1
	newRoot.Height = max(height(newRoot.Left), height(newRoot.Right)) + 1

	return newRoot
}

func (tree *AVLTree) Search(key string) (string , bool){
	node := search(tree.Root, hasher(key), key)
	if node == nil {
		return "", false
	}
	return node.Value, true
}

func search(node *Node, hashedkey int, key string) (*Node ) {
	if node == nil || node.HashedKey == hashedkey {
		return node
	}

	if hashedkey < node.HashedKey {
		return search(node.Left, hashedkey, key)
	}
	return search(node.Right, hashedkey, key)
}

func leftRotate(node *Node) *Node {
	newRoot := node.Right
	temp := newRoot.Left
	newRoot.Left = node
	node.Right = temp

	node.Height = max(height(node.Left), height(node.Right)) + 1
	newRoot.Height = max(height(newRoot.Left), height(newRoot.Right)) + 1

	return newRoot
}

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}


