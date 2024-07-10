public class BinarySearch {
    public static int binarySearch(int[] arr, int target) {
        var l = 0;
        var r = arr.length - 1;

        while (l <= r) {
            var mid = l + (r - l) / 2;
            var sample = arr[mid];

            if (sample == target)
                return mid;
            if (sample < target)
                l = mid + 1;
            if (sample > target)
                r = mid - 1;
        }

        return -1;
    }
}