/*
 * @lc app=leetcode id=215 lang=java
 *
 * [215] Kth Largest Element in an Array
 */
class Solution {

    private int[] nums;

    public int findKthLargest(int[] nums1, int k) {
        if (nums1 == null || nums1.length == 0) return -1;
        nums = new int[nums1.length];
        for (int i = 0; i < nums1.length; ++i) nums[i] = nums1[i];
        return findKthLargest(0, nums.length-1, k);
    }

    private int findKthLargest(int lo, int hi, int k) {
        int pivot = nums[hi];
        int loIdx = lo, hiIdx = hi;

        while (true) {
            while (nums[hiIdx] >= pivot && loIdx < hiIdx) hiIdx--;
            while (nums[loIdx] < pivot && loIdx < hiIdx) loIdx++;
            if (loIdx == hiIdx) {
                if (nums[loIdx] < pivot) loIdx++;
                break;
            }
            swap(loIdx, hiIdx);
        }

        if (loIdx - lo + k == hi - lo + 1) 
            return pivot;
        else if (loIdx - lo + k < hi - lo + 1)
            return findKthLargest(loIdx, hi-1, k);
        else
            return findKthLargest(lo, hiIdx, k - (hi - hiIdx));

    }

    private void swap(int i, int j) {
        int t = nums[i];
        nums[i] = nums[j];
        nums[j] = t;
    }


}

