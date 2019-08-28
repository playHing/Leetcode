/*
 * @lc app=leetcode id=75 lang=java
 *
 * [75] Sort Colors
 */
class Solution {
    public void sortColors(int[] nums) {
        if (nums == null) return;
        int zeroPtr = 0, twoPtr = nums.length-1;
        
        for (int onePtr = 0; onePtr <= twoPtr && zeroPtr < twoPtr;) {
            if (nums[onePtr] == 0) {
                nums[onePtr] = nums[zeroPtr];
                nums[zeroPtr++] = 0;
                if (onePtr < zeroPtr) onePtr = zeroPtr;
            } else if (nums[onePtr] == 2) {
                nums[onePtr] = nums[twoPtr];
                nums[twoPtr--] = 2;
            } else {
                onePtr++;
            }
        }
    }
}

