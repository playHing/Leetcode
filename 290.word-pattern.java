/*
 * @lc app=leetcode id=290 lang=java
 *
 * [290] Word Pattern
 */
class Solution {
    public boolean wordPattern(String pattern, String str) {
        if (pattern == null || str == null) return false;
        String[] sl = str.split(" ");
        if (sl.length != pattern.length()) return false;
        Map<Character, String> dict = new HashMap<>();
        Set<String> seem = new HashSet<>();
        for (int i = 0; i < pattern.length(); ++i) {
            String word = dict.get(pattern.charAt(i));
            if (word == null) {
                dict.put(pattern.charAt(i), sl[i]);
                if (seem.contains(sl[i])) return false;
                seem.add(sl[i]);
            }
            else if (!word.equals(sl[i])) return false;
        }
        return true;
    }
}

