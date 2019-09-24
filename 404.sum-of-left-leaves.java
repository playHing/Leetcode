/*
 * @lc app=leetcode id=404 lang=java
 *
 * [404] Sum of Left Leaves
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
    public int sumOfLeftLeaves(TreeNode root) {
        if (root == null) return 0;
        TreeNode leftNode = root.left;
        int leftSum = leftNode != null && 
            leftNode.left == null && leftNode.right == null ? 
            leftNode.val : sumOfLeftLeaves(root.left);
        int rightSum = sumOfLeftLeaves(root.right);
        return leftSum + rightSum;
    }
}

