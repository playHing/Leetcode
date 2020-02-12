/*
 * @lc app=leetcode id=108 lang=java
 *
 * [108] Convert Sorted Array to Binary Search Tree
 */
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    public TreeNode sortedArrayToBST(int[] nums) {
        if (nums == null || nums.length == 0) return null;
        return builder(nums, 0, nums.length-1);
    }

    private TreeNode builder(int[] nums, int l, int h) {
        int m = (l + h) / 2;
        TreeNode node = new TreeNode(nums[m]);
        if (l == h) return node;
        if (h - l == 1) {
            TreeNode tmp = new TreeNode(nums[h]);
            tmp.left = node;
            return tmp;
        }
        node.left = builder(nums, l, m-1);
        node.right = builder(nums, m+1, h);
        return node;
    }
}

