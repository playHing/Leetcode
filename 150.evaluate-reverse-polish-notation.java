/*
 * @lc app=leetcode id=150 lang=java
 *
 * [150] Evaluate Reverse Polish Notation
 */
class Solution {
    public int evalRPN(String[] tokens) {
        if (tokens == null) return 0;
        Stack<Integer> stack = new Stack<>();
        Map<String, Operator> operateMap = new HashMap<>();
        operateMap.put("+", new Plus());
        operateMap.put("-", new Minus());
        operateMap.put("*", new Times());
        operateMap.put("/", new Divide());
        for (String token : tokens) {
            if (operateMap.containsKey(token)) {
                int val2 = stack.pop(), val1 = stack.pop();
                stack.push(operateMap.get(token).operate(val1, val2));
            } else {
                stack.push(Integer.parseInt(token));
            }
        }
        return stack.pop();
    }
}

interface Operator {
    int operate(int val1, int val2);
}

class Plus implements Operator {
    public int operate(int val1, int val2) {
        return val1 + val2;
    }
}

class Minus implements Operator {
    public int operate(int val1, int val2) {
        return val1 - val2;
    }
}

class Times implements Operator {
    public int operate(int val1, int val2) {
        return val1 * val2;
    }
}

class Divide implements Operator {
    public int operate(int val1, int val2) {
        return val1 / val2;
    }
}

