/*
 * @lc app=leetcode id=125 lang=java
 *
 * [125] Valid Palindrome
 */
class Solution {
    public boolean isPalindrome(String s) {
        if (s == null || "".equals(s)) return true;
        int i = 0, j = s.length() - 1;
        s = s.toLowerCase();
        while (i < j) {
            while (i < j && !isValid(s.charAt(i))) i++;
            while (i < j && !isValid(s.charAt(j))) j--;
            if (s.charAt(i++) != s.charAt(j--)) {
                System.out.println(s.charAt(i-1) + " " + s.charAt(j+1));
                return false;
            }
        }
        return true;
    }

    private boolean isValid(char c) {
       int v = c - 'a';
       if (0 <= v && v <= 25) return true;
       v =  c - '0';
       if (0 <= v && v <= 9) return true;
       return false;
    }

}

