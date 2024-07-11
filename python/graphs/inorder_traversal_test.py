import unittest
from typing import List, Optional

from inorder_traversal import TreeNode, Solution


class TestInorderTraversal(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_single_node(self):
        root = TreeNode(1)
        result = self.solution.inorderTraversal(root)
        expected = [1]
        self.assertEqual(result, expected)

    def test_empty_tree(self):
        root = None
        result = self.solution.inorderTraversal(root)
        expected = []
        self.assertEqual(result, expected)

    def test_complete_tree(self):
        # Constructing the binary tree:
        #     1
        #    / \
        #   2   3
        #  / \
        # 4   5
        root = TreeNode(1)
        root.left = TreeNode(2)
        root.right = TreeNode(3)
        root.left.left = TreeNode(4)
        root.left.right = TreeNode(5)

        result = self.solution.inorderTraversal(root)
        expected = [4, 2, 5, 1, 3]
        self.assertEqual(result, expected)

    def test_right_skewed_tree(self):
        # Constructing the binary tree:
        # 1
        #  \
        #   2
        #    \
        #     3
        root = TreeNode(1)
        root.right = TreeNode(2)
        root.right.right = TreeNode(3)

        result = self.solution.inorderTraversal(root)
        expected = [1, 2, 3]
        self.assertEqual(result, expected)

    def test_left_skewed_tree(self):
        # Constructing the binary tree:
        #     3
        #    /
        #   2
        #  /
        # 1
        root = TreeNode(3)
        root.left = TreeNode(2)
        root.left.left = TreeNode(1)

        result = self.solution.inorderTraversal(root)
        expected = [1, 2, 3]
        self.assertEqual(result, expected)

    def test_complex_tree(self):
        # Constructing the binary tree:
        #       1
        #      / \
        #     2   3
        #    /   / \
        #   4   5   6
        #      /
        #     7
        root = TreeNode(1)
        root.left = TreeNode(2)
        root.right = TreeNode(3)
        root.left.left = TreeNode(4)
        root.right.left = TreeNode(5)
        root.right.right = TreeNode(6)
        root.right.left.left = TreeNode(7)

        result = self.solution.inorderTraversal(root)
        expected = [4, 2, 1, 7, 5, 3, 6]
        self.assertEqual(result, expected)


if __name__ == '__main__':
    unittest.main(exit=False)
