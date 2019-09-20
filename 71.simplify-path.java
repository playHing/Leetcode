/*
 * @lc app=leetcode id=71 lang=java
 *
 * [71] Simplify Path
 */
class Solution {
    public String simplifyPath(String path) {
        if (path == null) return null;
        String[] dirs = path.split("/");
        Stack<String> stack = new Stack<>();
        for (String dir : dirs) {
            //System.out.println("[" + dir + "]");
            switch (dir) {
                case "":
                case ".":
                    break;
                case "..":
                    if (!stack.empty()) stack.pop();
                    break;
                default:
                    stack.push(dir);
            }
        }
        StringBuffer buffer = new StringBuffer();
        while (!stack.empty()) {
            buffer.insert(0, "/" + stack.pop());
        }
        return buffer.length() == 0 ? "/" : buffer.toString();
    }
}

