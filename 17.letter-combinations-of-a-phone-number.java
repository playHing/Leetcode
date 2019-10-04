/*
 * @lc app=leetcode id=17 lang=java
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start
class Solution {

    String[] keyboard = new String[] {"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"};

    public List<String> letterCombinations(String digits) {
        List<String> res = new LinkedList<>();
        if (digits != null && digits.length() > 0)
            combinationHelper(res, new Stack<>(), digits);
        return res;
    }

    private void combinationHelper(List<String> res, Stack<Character> stack, String digits) {
        if (digits.length() == 0) {
            List<Character> chars = new ArrayList<>(stack);
            StringBuilder sb = new StringBuilder();
            for (Character ch: chars) sb.append(ch);
            res.add(sb.toString());
            return;
        }
        int d = digits.charAt(0) - '0';
        for (char c : keyboard[d].toCharArray()) {
            stack.push(c);
            combinationHelper(res, stack, digits.substring(1));
            stack.pop();
        }
    }

}
// @lc code=end

