/*
 * @lc app=leetcode id=101 lang=java
 *
 * [101] Symmetric Tree
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

    // Recursive
    public boolean isSymmetric(TreeNode root) {
        if (root == null) return true;
        return isReflect(root.left, root.right);
    }

    private boolean isReflect(TreeNode p, TreeNode q) {
        if (p == null && q == null) return true;
        if (p == null || q == null) return false;
        return p.val == q.val && 
            isReflect(p.left, q.right) && 
            isReflect(p.right, q.left);
    }

    // Iterative
    public boolean isSymmetric_ITERATIVE(TreeNode root) {
        if (root == null) return true;
        Stack<TreeNode> stack = new Stack<>();
        stack.push(root.left);
        stack.push(root.right);
        while (!stack.empty()) {
            TreeNode p = stack.pop();
            TreeNode q = stack.pop();
            if (p == null && q == null) continue;
            if (p == null || q == null || p.val != q.val) return false;
            stack.push(p.left);
            stack.push(q.right);
            stack.push(p.right);
            stack.push(q.left);
        }
        return true;
    }

}

