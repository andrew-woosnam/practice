import unittest
from binary_search import *

class BinarySearchTest(unittest.TestCase):
    def test_found(self):
        arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
        target = 6
        self.assertEqual(binary_search(arr, target), 5)

    def test_not_found(self):
        arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
        target = 11
        self.assertEqual(binary_search(arr, target), -1)

    def test_empty_array(self):
        arr = []
        target = 5
        self.assertEqual(binary_search(arr, target), -1)

    def test_single_element_array(self):
        arr = [5]
        target = 5
        self.assertEqual(binary_search(arr, target), 0)

    def test_large_array(self):
        arr = list(range(1000000))
        target = 999999
        self.assertEqual(binary_search(arr, target), 999999)

if __name__ == '__main__':
    unittest.main()