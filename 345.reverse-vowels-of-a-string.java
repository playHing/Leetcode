/*
 * @lc app=leetcode id=345 lang=java
 *
 * [345] Reverse Vowels of a String
 */
class Solution {
    public String reverseVowels(String s) {
        if (s == null || s.length() == 0) return s;
        char[] ss = s.toCharArray();
        int i = 0, j = s.length() - 1;
        while (i < j) {
            while (i < j && !isVowel(ss[i])) i++;
            while (i < j && !isVowel(ss[j])) j--;
            if (i < j) {
                char tmp = ss[i];
                ss[i++] = ss[j];
                ss[j--] = tmp;
            }
        }
        return String.valueOf(ss);
    }

    private boolean isVowel(char c) {
        switch (c) {
            case 'a':
            case 'e':
            case 'i':
            case 'o':
            case 'u':
            case 'A':
            case 'E':
            case 'I':
            case 'O':
            case 'U':
                return true;
            default:
                return false;
        }
    }
}

