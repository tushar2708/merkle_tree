package main

import "crypto/sha1"

type Node struct {
	hashVal Hash
	left    Hashable
	right   Hashable
}

func NewNode(l, r Hashable) {

}

func (n Node) hash() Hash {
	var l, r [sha1.Size]byte

	// calculate hash of each child
	l = n.left.hash()
	r = n.right.hash()

	// append the 2 hashes, and recalculate hash
	return ComputeHash(append(l[:], r[:]...))
}

// this function starts building the tree from the bottom
func BuildTree(parts []Hashable) []Hashable {
	var newNodes []Hashable
	var i int
	for i = 0; i < len(parts); i += 2 {
		if i < len(parts)-1 {
			newNodes = append(newNodes, Node{left: parts[i], right: parts[i+1]})
		} else {
			newNodes = append(newNodes, Node{left: parts[i], right: EmptyBlock{}})
		}
	}

	// if we have reached the rooot node, we are done
	if len(newNodes) == 1 {
		return newNodes
	}

	// create upper layer of nodes
	return BuildTree(newNodes)

}
