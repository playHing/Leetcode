/*
 * @lc app=leetcode id=438 lang=java
 *
 * [438] Find All Anagrams in a String
 */
class Solution {

    int[] ans = new int[26];

    public List<Integer> findAnagrams(String s, String p) {
        List<Integer> res = new LinkedList<>();
        if (s == null || p == null || s.length() < p.length()) return res;
        int[] freq = new int[26];
        for (int t = 0; t < p.length(); ++t) {
            ans[p.charAt(t)-'a']++;
            freq[s.charAt(t)-'a']++;
        }
        for (int i = 0; ; ++i) {
            if (isAnagram(freq)) res.add(i);
            if (i == s.length() - p.length()) break;
            freq[s.charAt(i)-'a']--;
            freq[s.charAt(i+p.length())-'a']++;
        }
        return res;
    }

    private boolean isAnagram(int[] freq) {
        for (int i = 0; i < 26; ++i) {
            if (freq[i] != ans[i]) return false;
        }
        return true;
    }
}

