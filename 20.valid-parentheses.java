/*
 * @lc app=leetcode id=20 lang=java
 *
 * [20] Valid Parentheses
 */
class Solution {
    public boolean isValid(String s) {
        if (s == null) return false;
        Stack<Character> stack = new Stack<>();
        for (char c : s.toCharArray()) {
            switch (c) {
                case ')':
                case ']':
                case '}':
                    if (stack.empty() || stack.pop() != lookupPair(c)) return false;
                    break;
                default:
                    stack.push(c);
            }
        }
        return stack.empty();
    }

    private char lookupPair(char c) {
        if (c == ')') return '(';
        if (c == ']') return '[';
        if (c == '}') return '{';
        return '0';
    }

}

