/*
 * @lc app=leetcode id=450 lang=java
 *
 * [450] Delete Node in a BST
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
    public TreeNode deleteNode(TreeNode root, int key) {
        if (root == null) return null;
        if (root.val > key) root.left = deleteNode(root.left, key);
        else if (root.val < key) root.right = deleteNode(root.right, key);
        else {
            if (root.left == null || root.right == null) {
                root = (root.left == null) ? root.right : root.left;
            } else {
                TreeNode curNode = root.right;
                while (curNode.left != null) curNode = curNode.left;
                root.val = curNode.val;
                root.right = deleteNode(root.right, curNode.val);
            }
        }
        
        return root;
    }
}

