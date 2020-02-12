/*
 * @lc app=leetcode id=107 lang=java
 *
 * [107] Binary Tree Level Order Traversal II
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
    public List<List<Integer>> levelOrderBottom(TreeNode root) {
        List<List<Integer>> res = new LinkedList<>();
        if (root == null) return res;
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);
        res.add(0, new LinkedList());
        TreeNode nextLevelFirst = null;

        while (!queue.isEmpty()) {

            TreeNode node = queue.poll();
            if (node == nextLevelFirst) {
                res.add(0, new LinkedList());
                nextLevelFirst = null;
            }
            res.get(0).add(node.val);

            if (node.left != null) {
                queue.add(node.left);
                if (nextLevelFirst == null)
                    nextLevelFirst = node.left;
            }
            if (node.right != null) {
                queue.add(node.right);
                if (nextLevelFirst == null)
                    nextLevelFirst = node.right;
            }
        }

        return res;
    }
}

