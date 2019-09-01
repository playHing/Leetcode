/*
 * @lc app=leetcode id=451 lang=java
 *
 * [451] Sort Characters By Frequency
 */
class Solution {

    public String frequencySort(String s) {
        if (s == null) return s;
        int[] freq = new int[256];
        PriorityQueue<Item> queue = new PriorityQueue<>();
        for (int i = 0; i < s.length(); ++i) ++freq[s.charAt(i)];
        for (int i = 0; i < freq.length; ++i) {
            if (freq[i] > 0) queue.add(new Item((char)i, freq[i]));
        }
        char[] res = new char[s.length()];
        for (int i = 0; i < res.length;) {
            Item item = queue.poll();
            for (int j = 0; j < item.f; ++j)
                res[i++] = item.c;
        }
        return new String(res);
    }

    class Item implements Comparable<Item> {
        final char c;
        final int f;
        public Item(char c, int f) {
            this.c = c;
            this.f = f;
        }

        public int compareTo(Item o) {
            if (f > o.f) return -1;
            else if (f < o.f) return 1;
            else return 0;
        }

    }
}

