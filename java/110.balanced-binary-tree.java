/*
 * @lc app=leetcode id=110 lang=java
 *
 * [110] Balanced Binary Tree
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

    boolean bal = true;

    public boolean isBalanced(TreeNode root) {
        height(root);
        return bal;
    }

    private int height(TreeNode root) {
        if (!bal) return 0;
        if (root == null) return -1;
        int lh = height(root.left);
        int rh = height(root.right);
        if (Math.abs(lh - rh) > 1) bal = false;
        return 1 + Math.max(lh, rh);
    }
}

