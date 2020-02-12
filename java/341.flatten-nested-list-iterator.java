/*
 * @lc app=leetcode id=341 lang=java
 *
 * [341] Flatten Nested List Iterator
 */
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * public interface NestedInteger {
 *
 *     // @return true if this NestedInteger holds a single integer, rather than a nested list.
 *     public boolean isInteger();
 *
 *     // @return the single integer that this NestedInteger holds, if it holds a single integer
 *     // Return null if this NestedInteger holds a nested list
 *     public Integer getInteger();
 *
 *     // @return the nested list that this NestedInteger holds, if it holds a nested list
 *     // Return null if this NestedInteger holds a single integer
 *     public List<NestedInteger> getList();
 * }
 */
public class NestedIterator implements Iterator<Integer> {

    Stack<NestedInteger> stack = new Stack<>();

    public NestedIterator(List<NestedInteger> nestedList) {
        if (nestedList != null) {
            pushListItemsToStack(nestedList);
        }
    }

    @Override
    public Integer next() {
        while (!stack.empty()) {
            NestedInteger item = stack.pop();
            if (item.isInteger()) return item.getInteger();
            pushListItemsToStack(item.getList());
        }
        return null;
    }

    @Override
    public boolean hasNext() {
        while (!stack.empty()) {
            if (stack.peek().isInteger()) return true;
            pushListItemsToStack(stack.pop().getList());
        }
        return false;
    }

    private void pushListItemsToStack(List<NestedInteger> nestedList) {
        Stack<NestedInteger> tmp = new Stack<>();
        for (NestedInteger e : nestedList) tmp.push(e);
        while (!tmp.empty()) stack.push(tmp.pop());
    }
}

/**
 * Your NestedIterator object will be instantiated and called as such:
 * NestedIterator i = new NestedIterator(nestedList);
 * while (i.hasNext()) v[f()] = i.next();
 */

