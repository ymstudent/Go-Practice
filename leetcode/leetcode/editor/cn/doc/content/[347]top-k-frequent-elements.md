<p>给你一个整数数组 <code>nums</code> 和一个整数 <code>k</code> ，请你返回其中出现频率前 <code>k</code> 高的元素。你可以按 <strong>任意顺序</strong> 返回答案。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入: </strong>nums = [1,1,1,2,2,3], k = 2
<strong>输出: </strong>[1,2]
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入: </strong>nums = [1], k = 1
<strong>输出: </strong>[1]</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li> 
 <li><code>k</code> 的取值范围是 <code>[1, 数组中不相同的元素的个数]</code></li> 
 <li>题目数据保证答案唯一，换句话说，数组中前 <code>k</code> 个高频元素的集合是唯一的</li> 
</ul>

<p>&nbsp;</p>

<p><strong>进阶：</strong>你所设计算法的时间复杂度 <strong>必须</strong> 优于 <code>O(n log n)</code> ，其中 <code>n</code><em>&nbsp;</em>是数组大小。</p>

<details><summary><strong>Related Topics</strong></summary>数组 | 哈希表 | 分治 | 桶排序 | 计数 | 快速选择 | 排序 | 堆（优先队列）</details><br>

<div>👍 1475, 👎 0</div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1，[手把手刷二叉树系列课程](https://aep.xet.tech/s/3YGcq3) 上线。**

<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

首先，肯定要用一个 `valToFreq` [哈希表](https://appktavsiei5995.pc.xiaoe-tech.com/detail/p_6265484ae4b01a4851f65633/6) 把每个元素出现的频率计算出来。

然后，这道题就变成了 [215. 数组中的第 K 个最大元素](/problems/kth-largest-element-in-an-array)，只不过第 215 题让你求数组中元素值 `e` 排在第 `k` 大的那个元素，这道题让你求数组中元素值 `valToFreq[e]` 排在前 `k` 个的元素。

我在 [快速排序详解及运用](https://labuladong.github.io/article/fname.html?fname=快速排序) 中讲过第 215 题，可以用 [优先级队列](https://labuladong.github.io/article/fname.html?fname=二叉堆详解实现优先级队列) 或者快速选择算法解决这道题。这里稍微改一下优先级队列的比较函数，或者改一下快速选择算法中的逻辑即可。

这里我再加一种解法，用计数排序的方式找到钱 `k` 个高频元素，见代码。

**标签：二叉堆，哈希表，快速选择**

## 解法代码

```java
// 用优先级队列解决这道题
class Solution {
    public int[] topKFrequent(int[] nums, int k) {
        // nums 中的元素 -> 该元素出现的频率
        HashMap<Integer, Integer> valToFreq = new HashMap<>();
        for (int v : nums) {
            valToFreq.put(v, valToFreq.getOrDefault(v, 0) + 1);
        }

        PriorityQueue<Map.Entry<Integer, Integer>>
                pq = new PriorityQueue<>((entry1, entry2) -> {
            // 队列按照键值对中的值（元素出现频率）从小到大排序
            return entry1.getValue().compareTo(entry2.getValue());
        });

        for (Map.Entry<Integer, Integer> entry : valToFreq.entrySet()) {
            pq.offer(entry);
            if (pq.size() > k) {
                // 弹出最小元素，维护队列内是 k 个频率最大的元素
                pq.poll();
            }
        }

        int[] res = new int[k];
        for (int i = k - 1; i >= 0; i--) {
            // res 数组中存储前 k 个最大元素
            res[i] = pq.poll().getKey();
        }

        return res;
    }
}

// 用计数排序的方法解决这道题
class Solution2 {
    public int[] topKFrequent(int[] nums, int k) {
        // nums 中的元素 -> 该元素出现的频率
        HashMap<Integer, Integer> valToFreq = new HashMap<>();
        for (int v : nums) {
            valToFreq.put(v, valToFreq.getOrDefault(v, 0) + 1);
        }

        // 频率 -> 这个频率有哪些元素
        ArrayList<Integer>[] freqToVals = new ArrayList[nums.length + 1];
        for (int val : valToFreq.keySet()) {
            int freq = valToFreq.get(val);
            if (freqToVals[freq] == null) {
                freqToVals[freq] = new ArrayList<>();
            }
            freqToVals[freq].add(val);
        }

        int[] res = new int[k];
        int p = 0;
        // freqToVals 从后往前存储着出现最多的元素
        for (int i = freqToVals.length - 1; i > 0; i--) {
            ArrayList<Integer> valList = freqToVals[i];
            if (valList == null) continue;
            for (int j = 0; j < valList.size(); j++) {
                // 将出现次数最多的 k 个元素装入 res
                res[p] = valList.get(j);
                p++;
                if (p == k) {
                    return res;
                }
            }
        }

        return null;
    }
}
```

**类似题目**：
  - [692. 前K个高频单词 🟠](/problems/top-k-frequent-words)
  - [剑指 Offer II 060. 出现频率最高的 k 个数字 🟠](/problems/g5c51o)

</details>
</div>



