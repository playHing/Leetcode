/*
 * @lc app=leetcode id=126 lang=java
 *
 * [126] Word Ladder II
 */
class Solution {
    public List<List<String>> findLadders(String beginWord, String endWord, List<String> wordList) {
        if (wordList == null || wordList.size() == 0) return new LinkedList<>();

        boolean foundEndWord = false;
        Map<String, Integer> wordMap = new HashMap<>();
        for (int i = 0; i < wordList.size(); ++i) {
            wordMap.put(wordList.get(i), i);
        }

        Map<String, List<List<String>>> wordChainMap = new HashMap<>();
        boolean[] visited = new boolean[wordList.size()];
        Queue<String> queue = new LinkedList<>();

        List<String> tmpSubList1 = new LinkedList<>();
        tmpSubList1.add(beginWord);
        List<List<String>> tmpList1 = new LinkedList<>();
        tmpList1.add(tmpSubList1);
        wordChainMap.put(beginWord, tmpList1);
        queue.offer(beginWord);

        while (!queue.isEmpty()) {
            Set<Integer> tmpVisited = new HashSet<>();
            int len = queue.size();
            for (int i = 0; i < len; ++i) {
                String word = queue.poll();
                List<List<String>> tmpList = wordChainMap.get(word);
                
                for (int pos = 0; pos < word.length(); ++pos) {
                    char[] wordSeq = word.toCharArray();
                    int old = wordSeq[pos];
                    for (int c = 'a'; c <= 'z'; ++c) {
                        if (c == old) continue;
                        wordSeq[pos] = (char)c;
                        String cand = String.valueOf(wordSeq);
                        if (wordMap.containsKey(cand) && !visited[wordMap.get(cand)]) {
                            List<List<String>> newList = new LinkedList<>();
                            for (int t = 0; t < tmpList.size(); ++t) {
                                List<String> tmpSubList = new LinkedList<>(tmpList.get(t));
                                tmpSubList.add(cand);
                                newList.add(tmpSubList);
                            }
                            if (wordChainMap.containsKey(cand)) {
                                wordChainMap.get(cand).addAll(newList);
                            } else {
                                wordChainMap.put(cand, newList);
                            }
                            if (tmpVisited.add(wordMap.get(cand))) queue.offer(cand);
                        }
                    }
                }
            }
            if (wordChainMap.containsKey(endWord)) return wordChainMap.get(endWord);
            for (int v : tmpVisited) visited[v] = true;
        }

        return new LinkedList<>();
    }
}

