// performs binary search on a sorted array, returning the target's index if found, else -1
export default function BinarySearch(arr: number[], target: number): number {
    var l = 0;
    var r = arr.length - 1;

    while (l <= r) {
        var middle = Math.floor(l + (r - l) / 2);
        var sample = arr[middle];

        if (sample === target) return middle;
        else if (sample < target) {
            l = middle + 1;
        } else { // sample > target
            r = middle - 1
        }
    }

    return -1
}
