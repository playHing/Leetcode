/*
 * @lc app=leetcode id=454 lang=java
 *
 * [454] 4Sum II
 */
class Solution {
    public int fourSumCount(int[] A, int[] B, int[] C, int[] D) {
        int res = 0;
        int[][] arrays = new int[][] {A, B, C, D};
        for (int[] item : arrays) Arrays.sort(item);

        for (int d = 0; d < D.length; ++d) {
            for (int c = 0; c < C.length; ++c) {
                int a = 0, b = B.length - 1;
                while (true) {
                    int sum = A[a] + B[b] + C[c] + D[d];
                    if (sum == 0) {
                        int ta = a, tb = b;
                        while (++ta < A.length && A[ta] == A[a]);
                        while (--tb >= 0 && B[tb] == B[b]);
                        res += (ta-a)*(b-tb);
                        if (ta == A.length || tb == -1) break;
                        a = ta;
                        b = tb;
                    }
                    else if (sum < 0 && a < A.length - 1) a++;
                    else if (sum > 0 && b > 0) b--;
                    else break;
                }
            }
        }

        return res;
    }
}

