/*
 * @lc app=leetcode id=401 lang=java
 *
 * [401] Binary Watch
 */

// @lc code=start
class Solution {

    class Time {
        final int hour;
        final int minute;
        public Time(int hour, int minute) {
            this.hour = hour;
            this.minute = minute;
        }
        public Time add(Time o) {
            return new Time(hour+o.hour, minute+o.minute);
        }
        public String toString() {
            return hour + ":" + (minute < 10 ? "0" : "") + minute;
        }
    }

    List<String> res = new LinkedList<>();

    public List<String> readBinaryWatch(int num) {
        Time[] times = new Time[] { 
            new Time(0,1), new Time(0,2), new Time(0,4), new Time(0,8),
            new Time(0,16), new Time(0,32),
            new Time(1,0), new Time(2,0), new Time(4,0), new Time(8,0)};
        if (num >= 0) backtrack(num, 0, times, new Time(0,0));
        return res;
    }

    private void backtrack(int num, int sp, Time[] times, Time time) {
        if (num == 0) {
            if (time.hour < 12 && time.minute < 60) res.add(time.toString());
            return;
        }
        for (int i = sp; times.length-i >= num; ++i) {
            backtrack(num-1, i+1, times, time.add(times[i]));
        }
    }
}
// @lc code=end

