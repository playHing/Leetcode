/*
 * @lc app=leetcode id=76 lang=java
 *
 * [76] Minimum Window Substring
 */
class Solution {
    public String minWindow(String s, String t) {
        if (s == null || t == null || s.length() < t.length()) return "";
        int[] freq = new int[256];
        int[] save = new int[256];
        for (int i = 0; i < t.length(); ++i) freq[t.charAt(i)]++;
        for (int i = 0; i < freq.length; ++i) if (freq[i] == 0) freq[i] = -1;
        int i = 0, j = -1, ai = 0, aj = s.length();
        int charLeft = t.length();
        
        while (i < s.length()) {
            if (charLeft > 0 && j < s.length() - 1) {

                if (freq[s.charAt(++j)] > 0) {
                    freq[s.charAt(j)]--;
                    charLeft--;
                }
                else if (freq[s.charAt(j)] == 0)
                    save[s.charAt(j)]++;

            }
            else {

                if (charLeft == 0 && j-i+1 < aj-ai+1) {
                    ai = i;
                    aj = j;
                }

                if (save[s.charAt(i)] > 0) {
                    save[s.charAt(i)]--;
                }
                else if (freq[s.charAt(i)] == 0) {
                    if (j == s.length() - 1) break;
                    freq[s.charAt(i)]++;
                    charLeft++;
                }

                i++;
            }
        }
        return aj == s.length() ? "" : s.substring(ai, aj+1);
    }
}

