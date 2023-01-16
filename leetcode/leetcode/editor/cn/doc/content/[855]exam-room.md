<p>在考场里，一排有&nbsp;<code>N</code>&nbsp;个座位，分别编号为&nbsp;<code>0, 1, 2, ..., N-1</code>&nbsp;。</p>

<p>当学生进入考场后，他必须坐在能够使他与离他最近的人之间的距离达到最大化的座位上。如果有多个这样的座位，他会坐在编号最小的座位上。(另外，如果考场里没有人，那么学生就坐在 0 号座位上。)</p>

<p>返回&nbsp;<code>ExamRoom(int N)</code>&nbsp;类，它有两个公开的函数：其中，函数&nbsp;<code>ExamRoom.seat()</code>&nbsp;会返回一个&nbsp;<code>int</code>&nbsp;（整型数据），代表学生坐的位置；函数&nbsp;<code>ExamRoom.leave(int p)</code>&nbsp;代表坐在座位 <code>p</code> 上的学生现在离开了考场。每次调用&nbsp;<code>ExamRoom.leave(p)</code>&nbsp;时都保证有学生坐在座位&nbsp;<code>p</code>&nbsp;上。</p>

<p>&nbsp;</p>

<p><strong>示例：</strong></p>

<pre><strong>输入：</strong>["ExamRoom","seat","seat","seat","seat","leave","seat"], [[10],[],[],[],[],[4],[]]
<strong>输出：</strong>[null,0,9,4,2,null,5]
<strong>解释：</strong>
ExamRoom(10) -&gt; null
seat() -&gt; 0，没有人在考场里，那么学生坐在 0 号座位上。
seat() -&gt; 9，学生最后坐在 9 号座位上。
seat() -&gt; 4，学生最后坐在 4 号座位上。
seat() -&gt; 2，学生最后坐在 2 号座位上。
leave(4) -&gt; null
seat() -&gt; 5，学生最后坐在 5 号座位上。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ol> 
 <li><code>1 &lt;= N &lt;= 10^9</code></li> 
 <li>在所有的测试样例中&nbsp;<code>ExamRoom.seat()</code>&nbsp;和&nbsp;<code>ExamRoom.leave()</code>&nbsp;最多被调用&nbsp;<code>10^4</code>&nbsp;次。</li> 
 <li>保证在调用&nbsp;<code>ExamRoom.leave(p)</code>&nbsp;时有学生正坐在座位 <code>p</code> 上。</li> 
</ol>

<details><summary><strong>Related Topics</strong></summary>设计 | 有序集合 | 堆（优先队列）</details><br>

<div>👍 206, 👎 0</div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1，[手把手刷二叉树系列课程](https://aep.xet.tech/s/3YGcq3) 上线。过年前最后一期打卡挑战 [开始报名](https://aep.xet.tech/s/1a9ByX)。**



<p><strong><a href="https://labuladong.github.io/article?qno=855" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

PS：这道题在[《算法小抄》](https://mp.weixin.qq.com/s/tUSovvogbR9StkPWb75fUw) 的第 389 页。

思路很简单：找最长的线段，从中间分隔成两段，中点就是 `seat()` 的返回值；找 `p` 的左右线段，合并成一个线段，这就是 `leave(p)` 的逻辑。

具体代码的逻辑涉及很多细节，请查看详细题解。

**详细题解：[如何调度考生的座位](https://labuladong.github.io/article/fname.html?fname=座位调度)**

**标签：[数据结构](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=1318892385270808576)**

## 解法代码

```java
class ExamRoom {
    // 将端点 p 映射到以 p 为左端点的线段
    private Map<Integer, int[]> startMap;
    // 将端点 p 映射到以 p 为右端点的线段
    private Map<Integer, int[]> endMap;
    // 根据线段长度从小到大存放所有线段
    private TreeSet<int[]> pq;
    private int N;

    public ExamRoom(int N) {
        this.N = N;
        startMap = new HashMap<>();
        endMap = new HashMap<>();
        pq = new TreeSet<>((a, b) -> {
            int distA = distance(a);
            int distB = distance(b);
            // 如果长度相同，就比较索引
            if (distA == distB)
                return b[0] - a[0];
            return distA - distB;
        });
        // 在有序集合中先放一个虚拟线段
        addInterval(new int[]{-1, N});
    }

    public int seat() {
        // 从有序集合拿出最长的线段
        int[] longest = pq.last();
        int x = longest[0];
        int y = longest[1];
        int seat;
        if (x == -1) { // 情况一
            seat = 0;
        } else if (y == N) { // 情况二
            seat = N - 1;
        } else { // 情况三
            seat = (y - x) / 2 + x;
        }
        // 将最长的线段分成两段
        int[] left = new int[]{x, seat};
        int[] right = new int[]{seat, y};
        removeInterval(longest);
        addInterval(left);
        addInterval(right);
        return seat;
    }

    public void leave(int p) {
        // 将 p 左右的线段找出来
        int[] right = startMap.get(p);
        int[] left = endMap.get(p);
        // 合并两个线段成为一个线段
        int[] merged = new int[]{left[0], right[1]};
        removeInterval(left);
        removeInterval(right);
        addInterval(merged);
    }


    /* 增加一个线段 */
    private void addInterval(int[] intv) {
        pq.add(intv);
        startMap.put(intv[0], intv);
        endMap.put(intv[1], intv);
    }

    /* 去除一个线段 */
    private void removeInterval(int[] intv) {
        pq.remove(intv);
        startMap.remove(intv[0]);
        endMap.remove(intv[1]);
    }

    /* 计算一个线段的长度 */
    private int distance(int[] intv) {
        int x = intv[0];
        int y = intv[1];
        if (x == -1) return y;
        if (y == N) return N - 1 - x;
        // 中点和端点之间的长度
        return (y - x) / 2;
    }
}
```

</details>
</div>



