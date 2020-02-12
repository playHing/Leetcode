/*
 * @lc app=leetcode id=202 lang=java
 *
 * [202] Happy Number
 */
class Solution {

    public boolean isHappy(int n) {
        Set<Integer> history = new HashSet<>();
        while (history.add(n)) {
            int sum = 0;
            while (n > 0) {
                sum += sq(n % 10);
                n /= 10;
            }
            if (sum == 1) return true;
            n = sum;
        }
        return false;
    }

    private int sq(int a) {
        return a * a;
    }
}

