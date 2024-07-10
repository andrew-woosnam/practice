import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

public class BinarySearchTest {
    @Test
    public void testBinarySearch_ElementPresent() {
        int[] arr = { 1, 3, 5, 7, 9, 11, 13, 15 };
        assertEquals(0, BinarySearch.binarySearch(arr, 1));
        assertEquals(2, BinarySearch.binarySearch(arr, 5));
        assertEquals(4, BinarySearch.binarySearch(arr, 9));
        assertEquals(6, BinarySearch.binarySearch(arr, 13));
        assertEquals(7, BinarySearch.binarySearch(arr, 15));
    }

    @Test
    public void testBinarySearch_ElementNotPresent() {
        int[] arr = { 1, 3, 5, 7, 9, 11, 13, 15 };
        assertEquals(-1, BinarySearch.binarySearch(arr, 2));
        assertEquals(-1, BinarySearch.binarySearch(arr, 4));
        assertEquals(-1, BinarySearch.binarySearch(arr, 8));
        assertEquals(-1, BinarySearch.binarySearch(arr, 14));
        assertEquals(-1, BinarySearch.binarySearch(arr, 16));
    }

    @Test
    public void testBinarySearch_EmptyArray() {
        int[] emptyArr = {};
        assertEquals(-1, BinarySearch.binarySearch(emptyArr, 10));
    }

    @Test
    public void testBinarySearch_SingleElementArray() {
        int[] singleElementArr = { 5 };
        assertEquals(0, BinarySearch.binarySearch(singleElementArr, 5));
        assertEquals(-1, BinarySearch.binarySearch(singleElementArr, 10));
    }

    @Test
    public void testBinarySearch_DuplicateElements() {
        int[] duplicateArr = { 2, 4, 4, 6, 8, 8, 10 };
        int result;

        result = BinarySearch.binarySearch(duplicateArr, 4);
        assertTrue(result == 1 || result == 2);

        result = BinarySearch.binarySearch(duplicateArr, 8);
        assertTrue(result == 4 || result == 5);

        assertEquals(-1, BinarySearch.binarySearch(duplicateArr, 5));
    }
}