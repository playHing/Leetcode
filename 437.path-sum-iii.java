/*
 * @lc app=leetcode id=437 lang=java
 *
 * [437] Path Sum III
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

    public int pathSum(TreeNode root, int sum) {
        if (root != null) count(root, sum, false);
        return total;
    }

    private void count(TreeNode root, int sum, boolean started) {
        if (root == null) return;
        if (root.val == sum) total++;
        if (!started) {
            count(root.left, sum, false);
            count(root.right, sum, false);
        }
        count(root.left, sum - root.val, true);
        count(root.right, sum - root.val, true);
    }

}

