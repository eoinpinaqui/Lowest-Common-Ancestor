package LCA

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// This test checks the successful case with the following tree with root, node1 and node2 as inputs:
//
//                  root
//                 /    \
//              node1   node2
//
func TestLCA_Success1(t *testing.T) {
	node1 := NewNode(1, "toFind1", nil)
	node2 := NewNode(2, "toFind2", nil)
	root := NewNode(0, "root", &[]*Node{node1, node2})

	answer, err := lowestCommonAncestor(root, node1, node2)
	assert.Equal(t, root, answer, "the root should be the LCA of node1 and node2")
	assert.Nil(t, err, "error should not be created when an LCA is found")
}

// This test checks the successful case with the following tree with root, node1 and node2 as inputs:
//
//                  root
//                 /    \
//             node3    node4________
//            /     \         \       \
//          node1  node6____  node9   node10
//         /      /   |     \
//      node5 node2   node7  node8
//
func TestLCA_Success2(t *testing.T) {
	node5 := NewNode(5, "node5", nil)
	node2 := NewNode(2, "node2", nil)
	node7 := NewNode(7, "node7", nil)
	node8 := NewNode(8, "node8", nil)
	node9 := NewNode(9, "node9", nil)
	node10 := NewNode(10, "node10", nil)
	node1 := NewNode(1, "node1", &[]*Node{node5})
	node6 := NewNode(6, "node6", &[]*Node{node2, node7, nil, node8})
	node3 := NewNode(3, "node3", &[]*Node{node1, node6})
	node4 := NewNode(4, "node4", &[]*Node{node9, node10})
	root := NewNode(0, "root", &[]*Node{node3, node4})

	answer, err := lowestCommonAncestor(root, node1, node2)

	assert.Equal(t, node3, answer, "the root should be the LCA of node1 and node2")
	assert.Nil(t, err, "error should not be created when an LCA is found")
}

// This test checks the case where an LCA does not exist between two nodes
func TestLCA_NoLCAExists(t *testing.T) {
	node1 := NewNode(1, "node1", nil)
	node2 := NewNode(2, "node2", nil)
	node3 := NewNode(3, "node3", nil)
	root := NewNode(0, "root", &[]*Node{node1, node3})

	answer, err := lowestCommonAncestor(root, node1, node2)

	assert.Nil(t, answer, "should return nil if there is no LCA")
	assert.NotNil(t, err, "error should be created if there is no LCA")
}

// This test checks whether an appropriate error is created when the root parameter is nil
func TestLCA_NilRoot(t *testing.T) {
	node1 := NewNode(1, "node1", nil)
	node2 := NewNode(2, "node2", nil)

	answer, err := lowestCommonAncestor(nil, node1, node2)
	assert.Nil(t, answer, "should return nil when an error occurs")
	assert.NotNil(t, err, "error should be created when there are nil parameters")
}

// This test checks whether an appropriate error is created when the root parameter is nil
func TestLCA_NilNode1(t *testing.T) {
	node1 := NewNode(1, "node1", nil)
	node2 := NewNode(2, "node2", nil)
	root := NewNode(0, "root", &[]*Node{node1, node2})

	answer, err := lowestCommonAncestor(root, nil, node2)
	assert.Nil(t, answer, "should return nil when an error occurs")
	assert.NotNil(t, err, "error should be created when there are nil parameters")
}

// This test checks whether an appropriate error is created when the root parameter is nil
func TestLCA_NilNode2(t *testing.T) {
	node1 := NewNode(1, "node1", nil)
	node2 := NewNode(2, "node2", nil)
	root := NewNode(0, "root", &[]*Node{node1, node2})

	answer, err := lowestCommonAncestor(root, node1, nil)
	assert.Nil(t, answer, "should return nil when an error occurs")
	assert.NotNil(t, err, "error should be created when there are nil parameters")
}
