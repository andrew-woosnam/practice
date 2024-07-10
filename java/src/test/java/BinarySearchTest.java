import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

public class BinarySearchTest {

    @Test
    public void testBinarySearch() {
        int[] arr = {1, 3, 5, 7, 9, 11, 13, 15};
        
        // Test cases for elements present in the array
        assertEquals(0, BinarySearch.binarySearch(arr, 1));
        assertEquals(2, BinarySearch.binarySearch(arr, 5));
        assertEquals(4, BinarySearch.binarySearch(arr, 9));
        assertEquals(6, BinarySearch.binarySearch(arr, 13));
        assertEquals(7, BinarySearch.binarySearch(arr, 15));
        
        // Test cases for elements not present in the array
        assertEquals(-1, BinarySearch.binarySearch(arr, 2));
        assertEquals(-1, BinarySearch.binarySearch(arr, 4));
        assertEquals(-1, BinarySearch.binarySearch(arr, 8));
        assertEquals(-1, BinarySearch.binarySearch(arr, 14));
        assertEquals(-1, BinarySearch.binarySearch(arr, 16));
        
        // Test case for an empty array
        int[] emptyArr = {};
        assertEquals(-1, BinarySearch.binarySearch(emptyArr, 10));
        
        // Test case for an array with a single element
        int[] singleElementArr = {5};
        assertEquals(0, BinarySearch.binarySearch(singleElementArr, 5));
        assertEquals(-1, BinarySearch.binarySearch(singleElementArr, 10));
        
        // Test case for an array with duplicate elements
        int[] duplicateArr = {2, 4, 4, 6, 8, 8, 10};
        assertEquals(1, BinarySearch.binarySearch(duplicateArr, 4));
        assertEquals(4, BinarySearch.binarySearch(duplicateArr, 8));
        assertEquals(-1, BinarySearch.binarySearch(duplicateArr, 5));
    }
}