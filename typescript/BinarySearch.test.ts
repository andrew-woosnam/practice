import BinarySearch from './BinarySearch';

describe('binarySearch', () => {
    it('should return the index of the target if found', () => {
        expect(BinarySearch([1, 2, 3, 4, 5], 3)).toBe(2);
        expect(BinarySearch([10, 20, 30, 40, 50], 40)).toBe(3);
        expect(BinarySearch([100, 200, 300, 400, 500], 500)).toBe(4);
    });

    it('should return -1 if the target is not found', () => {
        expect(BinarySearch([1, 2, 3, 4, 5], 6)).toBe(-1);
        expect(BinarySearch([10, 20, 30, 40, 50], 15)).toBe(-1);
        expect(BinarySearch([100, 200, 300, 400, 500], 1000)).toBe(-1);
    });

    it('should work with an empty array', () => {
        expect(BinarySearch([], 5)).toBe(-1);
    });

    it('should work with an array of one element', () => {
        expect(BinarySearch([10], 10)).toBe(0);
        expect(BinarySearch([20], 30)).toBe(-1);
    });

    it('should work with an array of duplicate elements', () => {
        expect(BinarySearch([1, 2, 2, 3, 4, 5], 2)).toBe(1);
        expect(BinarySearch([10, 20, 30, 30, 40, 50], 30)).toBe(2);
        expect(BinarySearch([100, 200, 300, 400, 400, 500], 400)).toBe(3);
    });
});