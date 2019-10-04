/*
 * @lc app=leetcode id=17 lang=java
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
class Solution {

    String[] keyboard = new String[] {"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"};

    public List<String> letterCombinations(String digits) {
        LinkedList<String> res = new LinkedList<>();
        if (digits != null && digits.length() > 0) {
            combinationHelper(res, digits);
        }
        return res;
    }

    private void combinationHelper(Queue<String> res, String digits) {
        int[] digitsArr = new int[digits.length()];
        for (int i = 0; i < digits.length(); ++i) {
            digitsArr[i] = digits.charAt(i) - '0';
        }
        res.offer("");
        int round = 0;
        while (round != digits.length()) {
            int len = res.size();
            for (int i = 0; i < len; ++i) {
                String s = res.poll();
                for (char c : keyboard[digitsArr[round]].toCharArray()) {
                    res.offer(s + c);
                }
            }
            round++;
        }
    }

}
// @lc code=end

