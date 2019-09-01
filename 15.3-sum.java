/*
 * @lc app=leetcode id=15 lang=java
 *
 * [15] 3Sum
 */
class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        if (nums == null || nums.length < 3) return new LinkedList<>();
        List<List<Integer>> res = new LinkedList<>();
        Set<Item> seemRes = new HashSet<>();
        for (int t = nums.length - 1; t >= 0; --t) {
            Map<Integer, Integer> idxMap = new HashMap<>();
            for (int i = 0; i < t; ++i) {
                if (idxMap.containsKey(nums[i])) {
                    Item item = new Item(nums[idxMap.get(nums[i])], nums[i], nums[t]);
                    if (!seemRes.contains(item)) {
                        res.add(Arrays.asList(new Integer[] {item.a, item.b, item.c}));
                        seemRes.add(item);
                    }
                }
                else
                    idxMap.put(-nums[t]-nums[i], i);
            }
        }
        return res;
    }

    class Item {
        public final int a, b, c;
        public Item(int x, int y, int z) {
            int[] tmp = new int[] {x, y, z};
            Arrays.sort(tmp);
            a = tmp[0]; b = tmp[1]; c = tmp[2];
        }
        
        @Override
        public boolean equals(Object oo) {
            Item o = (Item) oo;
            return a == o.a && b == o.b && c == o.c;
        }
        @Override
        public int hashCode() {
            return a * 11 * 11 + b * 7 + c * 3;
        }
    }
}

