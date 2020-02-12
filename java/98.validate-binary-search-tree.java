/*
 * @lc app=leetcode id=98 lang=java
 *
 * [98] Validate Binary Search Tree
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
    public boolean isValidBST(TreeNode root) {
        return validator(root, Integer.MIN_VALUE, Integer.MAX_VALUE);
    }

    private boolean validator(TreeNode root, int lowerBound, int upperBound) {
        if (root == null) return true;
        if (root.val == Integer.MIN_VALUE && root.left != null) return false;
        if (root.val == Integer.MAX_VALUE && root.right != null) return false;
        return lowerBound <= root.val && root.val <= upperBound && 
            validator(root.left, lowerBound, root.val-1) && 
            validator(root.right, root.val+1, upperBound);
    }
}

