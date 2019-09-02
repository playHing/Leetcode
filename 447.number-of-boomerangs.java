/*
 * @lc app=leetcode id=447 lang=java
 *
 * [447] Number of Boomerangs
 */
class Solution {
    public int numberOfBoomerangs(int[][] points) {
        int res = 0;
        if (points == null) return res;
        for (int i = 0; i < points.length; ++i) {
            Map<Integer, Integer> distMap = new HashMap<>();
            for (int j = 0; j < points.length; ++j) {
                if (j == i) continue;
                int key = dist(points[i], points[j]);
                distMap.put(key, distMap.getOrDefault(key, 0) + 1);
            }
            for (int count : distMap.values()) {
                res += (count - 1) * count;
            }
        }
        return res;
    }

    private int dist(int[] p1, int[] p2) {
        return (p1[0]-p2[0])*(p1[0]-p2[0])+(p1[1]-p2[1])*(p1[1]-p2[1]);
    }
}

