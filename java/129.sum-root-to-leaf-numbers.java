/*
 * @lc app=leetcode id=129 lang=java
 *
 * [129] Sum Root to Leaf Numbers
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

    int total = 0;

    public int sumNumbers(TreeNode root) {
        if (root != null) sumNumberAlongPaths(root, 0);
        return total;
    }

    private void sumNumberAlongPaths(TreeNode root, int sum) {
        sum = 10 * sum + root.val;
        if (root.left == null && root.right == null) total += sum;
        if (root.left != null) sumNumberAlongPaths(root.left, sum);
        if (root.right != null) sumNumberAlongPaths(root.right, sum);
    }
}

