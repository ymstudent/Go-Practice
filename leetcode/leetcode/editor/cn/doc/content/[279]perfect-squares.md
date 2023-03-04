<p>给你一个整数 <code>n</code> ，返回 <em>和为 <code>n</code> 的完全平方数的最少数量</em> 。</p>

<p><strong>完全平方数</strong> 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。例如，<code>1</code>、<code>4</code>、<code>9</code> 和 <code>16</code> 都是完全平方数，而 <code>3</code> 和 <code>11</code> 不是。</p>

<p>&nbsp;</p>

<p><strong>示例&nbsp;1：</strong></p>

<pre>
<strong>输入：</strong>n = <span><code>12</code></span>
<strong>输出：</strong>3 
<strong>解释：</strong><span><code>12 = 4 + 4 + 4</code></span></pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>n = <span><code>13</code></span>
<strong>输出：</strong>2
<strong>解释：</strong><span><code>13 = 4 + 9</code></span></pre>

&nbsp;

<p><strong>提示：</strong></p>

<ul> 
 <li><code>1 &lt;= n &lt;= 10<sup>4</sup></code></li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>广度优先搜索 | 数学 | 动态规划</details><br>

<div>👍 1629, 👎 0</div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1，[手把手刷二叉树系列课程](https://aep.xet.tech/s/3YGcq3) 上线。**

<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

这道题有一些比较有技巧性的数学方法，我这里不探讨数学，就用通用的动态规划思路解决：

题目问和为 `n` 的平方数的最小数量，那么我可以根据和为 `n-1x1, n-2*2, n-3x3...` 的平方数的最小数量推导出来。

这个思路就是状态转移方程，具体看代码，复杂度 `O(N*sqrt(N))`，也是不错的。

**标签：[动态规划](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=1318881141113536512)，[数学](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2122023604245659649)**

## 解法代码

```java
class Solution {
    public int numSquares(int n) {
        // 定义：和为 i 的平方数的最小数量是 dp[i]
        int[] dp = new int[n + 1];
        // base case
        dp[0] = 0;
        // 状态转移方程
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j * j <= i; j++) {
                // i - j * j 只要再加一个平方数 j * j 即可凑出 i
                dp[i] = min(dp[i], dp[i - j * j] + 1);
            }
        }
        return dp[n];
    }
}
```

</details>
</div>



