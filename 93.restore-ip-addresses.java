/*
 * @lc app=leetcode id=93 lang=java
 *
 * [93] Restore IP Addresses
 */

// @lc code=start
class Solution {

    List<String> res = new LinkedList<>();

    public List<String> restoreIpAddresses(String s) {
        if (s != null && s.length() >= 4) resolve("", s, 1);
        return res;
    }

    private void resolve(String cur, String s, int segment) {
        if (segment == 5) {
            if (s.length() == 0) res.add(cur.substring(1));
            return;
        }
        if (s.length() >= 1) {
            resolve(cur + "." + s.substring(0,1), s.substring(1), segment+1);
            if (s.charAt(0) == '0') return;
        }
        if (s.length() >= 2) {
            resolve(cur + "." + s.substring(0,2), s.substring(2), segment+1);
        }
        if (s.length() >= 3) {
            if (Integer.valueOf(s.substring(0,3)) <= 255) {
                resolve(cur + "." + s.substring(0,3), s.substring(3), segment+1);
            }
        }

    }
}
// @lc code=end

