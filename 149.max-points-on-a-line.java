/*
 * @lc app=leetcode id=149 lang=java
 *
 * [149] Max Points on a Line
 */
class Solution {

    public int maxPoints(int[][] points) {
        int res = 0;
        if (points == null || points.length == 0) return res;
        for (int i = 0; i < points.length; ++i) {
            Map<Slope, Integer> slopeMap = new HashMap<>();
            int coincidence = 0;
            for (int j = i + 1; j < points.length; ++j) {
                Slope key = new Slope(points[i], points[j]);
                if (Slope.COINCIDENCE_KEY.equals(key)) coincidence++;
                else slopeMap.put(key, slopeMap.getOrDefault(key, 0) + 1);
            }
            if (coincidence > res) res = coincidence;
            for (int count : slopeMap.values())
                if (count+coincidence > res) res = count+coincidence;
        }
        return res + 1;
    }

    static class Slope {

        public static final Slope COINCIDENCE_KEY = new Slope(0, 0);

        final int dx, dy;
        public Slope(int[] p1, int[] p2) {
            int tx = p1[0]-p2[0];
            int ty = p1[1]-p2[1];
            int gcdxy = gcd(tx, ty);
            if (gcdxy != 0) {
                tx /= gcdxy;
                ty /= gcdxy;
                dx = tx < 0 ? -tx : tx;
                dy = tx < 0 ? -ty : ty;
            } else {
                dx = 0;
                dy = 0;
            }
        }

        public Slope(int dx, int dy) {
            this.dx = dx;
            this.dy = dy;
        }

        private int gcd(int a, int b) {
            if (b == 0) return a;
            return gcd(b, a % b);
        }

        @Override
        public boolean equals(Object oo) {
            Slope o = (Slope) oo;
            return dx == o.dx && dy == o.dy;
        }

        @Override
        public int hashCode() {
            return dx * dy * 7 + 13;
        }
    }
}

