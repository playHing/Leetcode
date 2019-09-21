/*
 * @lc app=leetcode id=347 lang=java
 *
 * [347] Top K Frequent Elements
 */
class Solution {
    public List<Integer> topKFrequent(int[] nums, int k) {
        List<Integer> res = new LinkedList<>();
        if (nums == null || nums.length == 0) return res;
        Map<Integer, Integer> cntMap = new HashMap<>();
        for (int num : nums) {
            cntMap.put(num, cntMap.getOrDefault(num, 0) + 1);
        }
        PriorityQueue<CntPair> queue = new PriorityQueue<>((p1, p2) -> {
            if (p1.cnt < p2.cnt) return 1;
            else if (p1.cnt > p2.cnt) return -1;
            else return 0;
        });
        for (Map.Entry<Integer, Integer> entry : cntMap.entrySet()) {
            queue.offer(new CntPair(entry.getKey(), entry.getValue()));
        }
        while (!queue.isEmpty() && k-- > 0) {
            res.add(queue.poll().num);
        }
        return res;
    }

    static class CntPair {
        int num;
        int cnt;
        public CntPair(int num, int cnt) {
            this.num = num;
            this.cnt = cnt;
        }
    }
}

