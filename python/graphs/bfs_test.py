import unittest
from bfs import traverse_bfs


class TestBFS(unittest.TestCase):

    def setUp(self):
        # Set up a sample graph to be used in the tests
        self.graph = {
            'A': ['B', 'C'],
            'B': ['A', 'D', 'E'],
            'C': ['A', 'F'],
            'D': ['B'],
            'E': ['B', 'F'],
            'F': ['C', 'E']
        }

    def test_bfs_from_a(self):
        result = traverse_bfs(self.graph, 'A')
        expected_nodes = set(['A', 'B', 'C', 'D', 'E', 'F'])
        result_set = set(result)

        # Ensure all expected nodes are visited
        self.assertEqual(result_set, expected_nodes)
        # Ensure no duplicates
        self.assertEqual(len(result), len(result_set))
        # Ensure start node is the first in the result
        self.assertEqual(result[0], 'A')

    def test_bfs_from_d(self):
        result = traverse_bfs(self.graph, 'D')
        expected_nodes = set(['D', 'B', 'A', 'E', 'C', 'F'])
        result_set = set(result)

        # Ensure all expected nodes are visited
        self.assertEqual(result_set, expected_nodes)
        # Ensure no duplicates
        self.assertEqual(len(result), len(result_set))
        # Ensure start node is the first in the result
        self.assertEqual(result[0], 'D')

    def test_empty_graph(self):
        result = traverse_bfs({}, 'A')
        expected = []
        self.assertEqual(result, expected)

    def test_node_not_in_graph(self):
        result = traverse_bfs(self.graph, 'Z')
        expected = ['Z']
        self.assertEqual(result, expected)


if __name__ == '__main__':
    unittest.main(exit=False)
