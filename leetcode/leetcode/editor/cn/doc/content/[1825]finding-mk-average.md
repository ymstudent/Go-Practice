<p>给你两个整数&nbsp;<code>m</code>&nbsp;和&nbsp;<code>k</code>&nbsp;，以及数据流形式的若干整数。你需要实现一个数据结构，计算这个数据流的 <b>MK 平均值</b>&nbsp;。</p>

<p><strong>MK 平均值</strong>&nbsp;按照如下步骤计算：</p>

<ol> 
 <li>如果数据流中的整数少于 <code>m</code>&nbsp;个，<strong>MK 平均值</strong>&nbsp;为 <code>-1</code>&nbsp;，否则将数据流中最后 <code>m</code>&nbsp;个元素拷贝到一个独立的容器中。</li> 
 <li>从这个容器中删除最小的 <code>k</code>&nbsp;个数和最大的 <code>k</code>&nbsp;个数。</li> 
 <li>计算剩余元素的平均值，并 <strong>向下取整到最近的整数</strong>&nbsp;。</li> 
</ol>

<p>请你实现&nbsp;<code>MKAverage</code>&nbsp;类：</p>

<ul> 
 <li><code>MKAverage(int m, int k)</code>&nbsp;用一个空的数据流和两个整数 <code>m</code>&nbsp;和 <code>k</code>&nbsp;初始化&nbsp;<strong>MKAverage</strong>&nbsp;对象。</li> 
 <li><code>void addElement(int num)</code>&nbsp;往数据流中插入一个新的元素&nbsp;<code>num</code>&nbsp;。</li> 
 <li><code>int calculateMKAverage()</code>&nbsp;对当前的数据流计算并返回 <strong>MK 平均数</strong>&nbsp;，结果需 <strong>向下取整到最近的整数</strong> 。</li> 
</ul>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>
["MKAverage", "addElement", "addElement", "calculateMKAverage", "addElement", "calculateMKAverage", "addElement", "addElement", "addElement", "calculateMKAverage"]
[[3, 1], [3], [1], [], [10], [], [5], [5], [5], []]
<strong>输出：</strong>
[null, null, null, -1, null, 3, null, null, null, 5]

<strong>解释：</strong>
MKAverage obj = new MKAverage(3, 1); 
obj.addElement(3);        // 当前元素为 [3]
obj.addElement(1);        // 当前元素为 [3,1]
obj.calculateMKAverage(); // 返回 -1 ，因为 m = 3 ，但数据流中只有 2 个元素
obj.addElement(10);       // 当前元素为 [3,1,10]
obj.calculateMKAverage(); // 最后 3 个元素为 [3,1,10]
                          // 删除最小以及最大的 1 个元素后，容器为 <span><code>[3]
                          // [3] 的平均值等于 3/1 = 3 ，故返回 3
obj.addElement(5);        // 当前元素为 [3,1,10,5]
obj.addElement(5);        // 当前元素为 [3,1,10,5,5]
obj.addElement(5);        // 当前元素为 [3,1,10,5,5,5]
obj.calculateMKAverage(); // 最后 3 个元素为 [5,5,5]
                          // </code></span>删除最小以及最大的 1 个元素后，容器为 <span><code>[5]<code>
                          // </code>[5] 的平均值等于 5/1 = 5 ，故返回 5<code>
</code></code></span></pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>3 &lt;= m &lt;= 10<sup>5</sup></code></li> 
 <li><code>1 &lt;= k*2 &lt; m</code></li> 
 <li><code>1 &lt;= num &lt;= 10<sup>5</sup></code></li> 
 <li><code>addElement</code> 与&nbsp;<code>calculateMKAverage</code>&nbsp;总操作次数不超过 <code>10<sup>5</sup></code> 次。</li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>设计 | 队列 | 数据流 | 有序集合 | 堆（优先队列）</details><br>

<div>👍 42, 👎 0</div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1，[手把手刷二叉树系列课程](https://aep.xet.tech/s/3YGcq3) 上线。**

</div>



