/*
 * @lc app=leetcode id=205 lang=java
 *
 * [205] Isomorphic Strings
 */
class Solution {
    public boolean isIsomorphic(String s, String t) {
        if (s == null || t == null) return false;
        char[] ca = t.toCharArray();
        if (ca.length != s.length()) return false;
        Map<Character, Character> dict = new HashMap<>();
        Set<Character> seem = new HashSet<>();
        for (int i = 0; i < s.length(); ++i) {
            Character c = dict.get(s.charAt(i));
            if (c == null) {
                dict.put(s.charAt(i), ca[i]);
                if (seem.contains(ca[i])) return false;
                seem.add(ca[i]);
            }
            else if (c != ca[i]) return false;
        }
        return true;
    }
}

