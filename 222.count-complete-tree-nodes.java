/*
 * @lc app=leetcode id=222 lang=java
 *
 * [222] Count Complete Tree Nodes
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
    public int countNodes(TreeNode root) {
        int h = height(root);
        if (h == -1) return 0;
        return height(root.right) == h-1 ? 
            (1 << h) + countNodes(root.right) :
            (1 << (h-1)) + countNodes(root.left);
    }

    private int height(TreeNode root) {
        if (root == null) return -1;
        return 1 + height(root.left);
    }
}

