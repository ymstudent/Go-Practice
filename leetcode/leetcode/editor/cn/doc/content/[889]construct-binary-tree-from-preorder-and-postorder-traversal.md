<p>给定两个整数数组，<code>preorder</code>&nbsp;和 <code>postorder</code> ，其中 <code>preorder</code> 是一个具有 <strong>无重复</strong> 值的二叉树的前序遍历，<code>postorder</code> 是同一棵树的后序遍历，重构并返回二叉树。</p>

<p>如果存在多个答案，您可以返回其中 <strong>任何</strong> 一个。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<p><img src="https://assets.leetcode.com/uploads/2021/07/24/lc-prepost.jpg" style="height: 265px; width: 304px;" /></p>

<pre>
<strong>输入：</strong>preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
<strong>输出：</strong>[1,2,3,4,5,6,7]
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong> preorder = [1], postorder = [1]
<strong>输出:</strong> [1]
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>1 &lt;= preorder.length &lt;= 30</code></li> 
 <li><code>1 &lt;= preorder[i] &lt;= preorder.length</code></li> 
 <li><code>preorder</code>&nbsp;中所有值都 <strong>不同</strong></li> 
 <li><code>postorder.length == preorder.length</code></li> 
 <li><code>1 &lt;= postorder[i] &lt;= postorder.length</code></li> 
 <li><code>postorder</code>&nbsp;中所有值都 <strong>不同</strong></li> 
 <li>保证 <code>preorder</code>&nbsp;和 <code>postorder</code>&nbsp;是同一棵二叉树的前序遍历和后序遍历</li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>树 | 数组 | 哈希表 | 分治 | 二叉树</details><br>

<div>👍 290, 👎 0</div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1。过年前最后一期打卡挑战 [开始报名](https://aep.xet.tech/s/1a9ByX)。**



<p><strong><a href="https://labuladong.github.io/article?qno=889" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

本题思路为手把手刷二叉树系列教程的专属内容，用 [手把手带你刷二叉树（纲领篇）](https://labuladong.gitee.io/article/fname.html?fname=二叉树总结) 中的两种思维模式秒杀所有二叉树的题目，并延伸到回溯算法和动态规划系列问题。点击 [这里](https://aep.xet.tech/s/2zdvE2) 解锁《手把手刷通二叉树》的全部内容。</details>
</div>



