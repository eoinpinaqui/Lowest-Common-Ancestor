import unittest
from LCA import *


class MyTestCase(unittest.TestCase):
    # This test checks the successful case with the following tree with root, node1 and node2 as inputs:
    #
    #                   root
    #                  /    \
    #               node1   node2
    def test_Success1(self):
        node1 = Node(1, "node1")
        node2 = Node(2, "node2")
        root = Node(0, "root")
        root.children = [node1, node2]
        self.assertEqual(root, lowest_common_ancestor(root, node1, node2))

    # This test checks the successful case with the following tree with root, node1 and node2 as inputs:
    #
    #                   root
    #                  /    \
    #               node3   node4 _________
    #              /     \        \        \
    #           node1    node6 __  node9   node10
    #           /       /   |    \
    #       node5   node2  node7  node8
    def test_Success2(self):
        node5 = Node(5, "node5")
        node2 = Node(2, "node2")
        node7 = Node(7, "node7")
        node8 = Node(8, "node8")
        node9 = Node(9, "node9")
        node10 = Node(10, "node10")
        node1 = Node(1, "node1")
        node1.children = [node5]
        node6 = Node(6, "node6")
        node6.children = [node2, node7, node8]
        node3 = Node(3, "node3")
        node3.children = [node1, node6]
        node4 = Node(4, "node4")
        node4.children = [node9, node10]
        root = Node(0, "root")
        root.children = [node3, node4]
        self.assertEqual(node3, lowest_common_ancestor(root, node1, node2))

    # This test checks the case where an LCA does not exist between two nodes
    def test_LCADoesNotExist(self):
        node1 = Node(1, "node1")
        node2 = Node(2, "node2")
        node3 = Node(3, "node3")
        root = Node(0, "root")
        root.children = [node1, node3]
        self.assertEqual(None, lowest_common_ancestor(root, node1, node2))

    # This test checks the case where the root is None
    def test_NilRoot(self):
        self.assertEqual(None, lowest_common_ancestor(None, Node(1, "node1"), Node(2, "node2")))

    # This test checks the case where the root is None
    def test_NilNode1(self):
        self.assertEqual(None, lowest_common_ancestor(Node(0, "root"), None, Node(2, "node2")))

    # This test checks the case where the root is None
    def test_NilNode1(self):
        self.assertEqual(None, lowest_common_ancestor(Node(0, "root"), Node(1, "node1"), None))


if __name__ == '__main__':
    unittest.main()
