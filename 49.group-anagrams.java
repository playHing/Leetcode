/*
 * @lc app=leetcode id=49 lang=java
 *
 * [49] Group Anagrams
 */
class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        List<List<String>> res = new LinkedList<>();
        if (strs == null) return res;
        Map<String, List<String>> anagramsMap = new HashMap<>();
        for (int i = 0; i < strs.length; ++i) {
            String k = key(strs[i]);
            if (anagramsMap.containsKey(k)) anagramsMap.get(k).add(strs[i]);
            else {
                List<String> freshGp = new LinkedList<>();
                freshGp.add(strs[i]);
                res.add(freshGp);
                anagramsMap.put(k, freshGp);
            }
        }
        return res;
    }

    private String key(String word) {
        char[] wordArray = word.toCharArray();
        Arrays.sort(wordArray);
        return new String(wordArray);
    }
}

