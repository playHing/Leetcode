/*
 * @lc app=leetcode id=454 lang=java
 *
 * [454] 4Sum II
 */
class Solution {
    public int fourSumCount(int[] A, int[] B, int[] C, int[] D) {
        int res = 0;
        Map<Integer, Integer> map = new HashMap<>();
        for (int d = 0; d < D.length; ++d) {
            for (int c = 0; c < C.length; ++c) {
                map.put(-C[c]-D[d], map.getOrDefault(-C[c]-D[d], 0) + 1);
            }
        }
        for (int b = 0; b < B.length; ++b) {
            for (int a = 0; a < A.length; ++a) {
                if (map.containsKey(A[a]+B[b]))
                    res += map.get(A[a]+B[b]);
            }
        }

        return res;
    }
}

