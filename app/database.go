package main

// Implement avl trees to Insert, Delete, and Search for a node in the tree in log(n) time.

type Node struct {
	HashedKey int
	Key string
	Value string
	Height int
	Left  *Node
	Right *Node
}