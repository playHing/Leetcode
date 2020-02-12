/*
 * @lc app=leetcode id=3 lang=java
 *
 * [3] Longest Substring Without Repeating Characters
 */
class Solution {
    public int lengthOfLongestSubstring(String s) {
        if (s == null || s.length() == 0) return 0;
        int max = 1;
        int i = 0, j = -1;
        int[] lookup = new int[256];
        Arrays.fill(lookup, -1);

        while (j < s.length() - 1) {
            if (lookup[s.charAt(++j)] >= i)
                i = lookup[s.charAt(j)] + 1;
            
            lookup[s.charAt(j)] = j;
            if (max < j-i+1) max = j-i+1;
        }
        return max;
    }
}

