/*
 * @lc app=leetcode id=11 lang=java
 *
 * [11] Container With Most Water
 */
class Solution {
    public int maxArea(int[] height) {
        if (height == null || height.length <= 1) return 0;
        int i = 0, j = height.length - 1;
        int max = 0;
        while (i < j) {
            int vol = calVol(height, i, j);
            if (vol > max) max = vol;
            if (height[i] < height[j]) i++;
            else j--;
        }
        return max;
    }

    private int min(int a, int b) {
        return a < b ? a : b;
    }

    private int calVol(int[] height, int i, int j) {
        return (j-i)*(min(height[i], height[j]));
    }
}

