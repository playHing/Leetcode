/*
 * @lc app=leetcode id=127 lang=java
 *
 * [127] Word Ladder
 */
class Solution {
    public int ladderLength(String beginWord, String endWord, List<String> wordList) {
        if (wordList == null || wordList.size() == 0) return 0;
        int d = 2;
        Queue<String> queue = new LinkedList<>();
        queue.offer(beginWord);
        boolean[] visited = new boolean[wordList.size()];
        while (!queue.isEmpty()) {
            int len = queue.size();
            for (int i = 0; i < len; ++i) {
                String word = queue.poll();
                for (int idx = 0; idx < wordList.size(); ++idx) {
                    if (visited[idx]) continue;
                    String cand = wordList.get(idx);
                    if (dist(word, cand) == 1) {
                        if (cand.equals(endWord)) {
                            return d;
                        }
                        visited[idx] = true;
                        queue.offer(cand);
                    }
                }
            }
            d++;
        }
        return 0;
    }

    private int dist(String w1, String w2) {
        int d = 0;
        if (w1.equals(w2)) return 0;
        for (int i  = 0; i < w1.length(); ++i) {
            if (w1.charAt(i) != w2.charAt(i)) d++;
        }
        return d;
    }
}

