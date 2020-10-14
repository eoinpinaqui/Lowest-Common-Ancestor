# LCA.py

class Node:
    def __init__(self, key, value):
        self.key = key
        self.value = value
        self.children = {}


def lowest_common_ancestor(root, node1, node2):
    # If any of the parameters are nil, return an error.
    if root is None or node1 is None or node2 is None:
        return None

    # Traverse the tree to find the answer.
    result = traverse(root, node1, node2)

    # If the LCA has been found, return it with no error.
    if result[1] is not None:
        return result[1]

    return None


def traverse(current, node1, node2):
    # If at the end of a branch, you have not found a branch containing an input node.
    if current is None:
        return False, None

    # Traverse over all of the children to look for a branches that contain the input nodes.
    branches_found = 0
    if current.children is not None:
        for child in current.children:
            result = traverse(child, node1, node2)
            if result[1] is not None:
                return True, result[1]
            elif result[0] is True:
                branches_found = branches_found + 1

    # If the current node is an input node, you have found a branch with an input node.
    if current.key is node1.key or current.key is node2.key:
        branches_found = branches_found + 1

    # If you have found two branches with at least one input node, the current node is the LCA.
    if branches_found >= 2:
        return True, current

    # Return if you have found any branches with at least one input node.
    return branches_found > 0, None
