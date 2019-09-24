/*
 * @lc app=leetcode id=257 lang=java
 *
 * [257] Binary Tree Paths
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

    public List<String> binaryTreePaths(TreeNode root) {
        List<String> res = new LinkedList<>();
        if (root != null) drawPaths(root, "", res);
        return res;
    }

    public void drawPaths(TreeNode root, String path, List<String> res) {
        if (root.left == null && root.right == null) res.add(path + root.val);
        if (root.left != null) drawPaths(root.left, path + root.val + "->", res);
        if (root.right != null) drawPaths(root.right, path + root.val + "->", res);
    }
}

