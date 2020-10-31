package LCA

import (
	"errors"
)

// The Node struct is used to represent the nodes in the graph.
type Node struct {
	Key      interface{}
	Value    interface{}
	Children *[]*Node
}

func NewNode(key interface{}, value interface{}, children *[]*Node) *Node {
	return &Node{
		Key:      key,
		Value:    value,
		Children: children,
	}
}

// The DAG struct is used to represent a Directed Acyclic Graphs.
type DAG struct {
	Nodes 	[]*Node
}

func NewDAG() *DAG {
	return &DAG{
		Nodes:	make([]*Node, 0),
	}
}

// The addNode function adds a node to the DAG
func (d *DAG) addNode(node *Node) {
	d.Nodes = append(d.Nodes, node)
}

// The LCA function returns the lowest common ancestor of two given nodes in the DAG
func (d *DAG) LCA(node1 *Node, node2 *Node) (*Node, error) {
	// If any of the parameters are nil, return an error.
	if node1 == nil || node2 == nil {
		return nil, errors.New("nil parameters present")
	}

	// Check all nodes to find the LCA.
	for _, node := range d.Nodes {
		_, _, LCA := traverse(node, node1, node2)

		// If the LCA has been found, return it with no error.
		if LCA != nil {
			return LCA, nil
		}
	}

	// If the LCA has not been found, return an error.
	return nil, errors.New("lowest common ancestor not found")
}

// The traverse function traverses the children of a node in search of the two input nodes.
func traverse(current *Node, node1 *Node, node2 *Node) (bool, int, *Node) {
	// If at the end of a branch, you have not found a branch containing an input node.
	if current == nil {
		return false, 0, nil
	}

	// Traverse over all of the children to look for a branches that contain the input nodes.
	branchesFound := 0

	if current.Children != nil {
		for _, child := range *current.Children {
			found, branches, LCA := traverse(child, node1, node2)
			if LCA != nil {
				return true, branches, LCA
			} else if branches >= 2 {
				return true, branches, current
			} else if found == true {
				branchesFound++
			}
		}
	}

	// If you have found  two branches with at least one input node, the current node is the LCA.
	if branchesFound >= 2 {
		return true, branchesFound, current
	}

	// If the current node is an input node, you have found a branch with an input node.
	if (current.Key == node1.Key) || (current.Key == node2.Key) {
		branchesFound++
	}

	// Return if you have found any branches with at least one input node.
	return branchesFound > 0, branchesFound, nil
}
