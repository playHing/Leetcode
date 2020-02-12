/*
 * @lc app=leetcode id=236 lang=java
 *
 * [236] Lowest Common Ancestor of a Binary Tree
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

    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if(root == null || root == p || root == q)  return root;
        TreeNode left = lowestCommonAncestor(root.left, p, q);
        TreeNode right = lowestCommonAncestor(root.right, p, q);
        if(left != null && right != null)   return root;
        return left != null ? left : right;
    }

    /*
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if (p == q) return p;
        Stack<TreeNode> sp = new Stack<>();
        Stack<TreeNode> sq = new Stack<>();
        if (!findNode(root, p, sp) || !findNode(root, q, sq)) return null;
        while (sp.size() < sq.size()) sq.pop();
        while (sq.size() < sp.size()) sp.pop();
        TreeNode res = sp.pop();
        while (res != sq.pop()) res = sp.pop();
        return res;
    }

    private boolean findNode(TreeNode root, TreeNode target, Stack<TreeNode> path) {
        if (root == null) return false;
        path.push(root);
        if (root == target) return true;
        if (!findNode(root.left, target, path) && !findNode(root.right, target, path)) {
            path.pop();
            return false;
        }
        return true;
    }
    */

}

