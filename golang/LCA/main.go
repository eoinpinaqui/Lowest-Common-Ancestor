package LCA

import (
	"errors"
)

type Node struct {
	Key      interface{}
	Value    interface{}
	Children *[]Node
}

func NewNode(value interface{}, key interface{}, nodes *[]Node) *Node {
	return &Node{
		Key:      key,
		Value:    value,
		Children: nodes,
	}
}

func lowestCommonAncestor(root *Node, node1 *Node, node2 *Node) (*Node, error) {
	// If any of the parameters are nil, return an error.
	if root == nil || node1 == nil || node2 == nil {
		return nil, errors.New("nil parameters present")
	}

	// Traverse the tree to find the answer.
	_, answerNode := traverse(root, node1, node2)

	// If the LCA has been found, return it with no error.
	if answerNode != nil {
		return answerNode, nil
	}

	// If the LCA has not been found, return an error.
	return nil, errors.New("lowest common ancestor not found")
}

func traverse(current *Node, node1 *Node, node2 *Node) (bool, *Node) {
	// If at the end of a branch, you have not found a branch containing an input node.
	if current == nil {
		return false, nil
	}

	// Traverse over all of the children to look for a branches that contain the input nodes.
	branchesFound := 0
	for _, child := range *current.Children {
		found, answer := traverse(&child, node1, node2)
		if answer != nil {
			return true, answer
		} else if found == true {
			branchesFound++
		}
	}

	// If the current node is an input node, you have found a branch with an input node.
	if (current.Key == node1.Key) || (current.Key == node2.Key) {
		branchesFound++
	}

	// If you have found  two branches with at least one input node, the current node is the LCA.
	if branchesFound >= 2 {
		return true, current
	}

	// Return if you have found any branches with at least one input node.
	return branchesFound > 0, nil
}
