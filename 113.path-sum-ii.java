/*
 * @lc app=leetcode id=113 lang=java
 *
 * [113] Path Sum II
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

    List<List<Integer>> res = new LinkedList<>();
    Stack<Integer> stack = new Stack<>();

    public List<List<Integer>> pathSum(TreeNode root, int sum) {
        if (root != null) fillValidPaths(root, sum);
        return res;
    }

    private void fillValidPaths(TreeNode root, int sum) {
        sum -= root.val;
        stack.push(root.val);
        if (root.left == null && root.right == null && sum == 0) res.add(new LinkedList<>(stack));
        if (root.left != null) fillValidPaths(root.left, sum);
        if (root.right != null) fillValidPaths(root.right, sum);
        stack.pop();
    }
}

