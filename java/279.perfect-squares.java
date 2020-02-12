/*
 * @lc app=leetcode id=279 lang=java
 *
 * [279] Perfect Squares
 */
class Solution {
    public int numSquares(int n) {
        if (n <= 0) return 0;
        boolean[] visited = new boolean[n+1];
        Queue<Integer> queue = new LinkedList<>();
        visited[n] = true;
        queue.offer(n);
        int dist = 1;
        
        while (!queue.isEmpty()) {

            int len = queue.size();
            for (int i = 0; i < len; ++i) {
                int idx = queue.poll();
                int val = 1;
                while (val * val < idx) {
                    int cand = idx - val * val;
                    if (!visited[cand]) {
                        visited[cand] = true;
                        queue.offer(cand);
                    }
                    val++;
                }
                if (val * val == idx) return dist;
            }
            
            dist++;
        }

        return n;
    }
}

