<p>在一个由 <code>'0'</code> 和 <code>'1'</code> 组成的二维矩阵内，找到只包含 <code>'1'</code> 的最大正方形，并返回其面积。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p> 
<img alt="" src="https://assets.leetcode.com/uploads/2020/11/26/max1grid.jpg" style="width: 400px; height: 319px;" /> 
<pre>
<strong>输入：</strong>matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
<strong>输出：</strong>4
</pre>

<p><strong>示例 2：</strong></p> 
<img alt="" src="https://assets.leetcode.com/uploads/2020/11/26/max2grid.jpg" style="width: 165px; height: 165px;" /> 
<pre>
<strong>输入：</strong>matrix = [["0","1"],["1","0"]]
<strong>输出：</strong>1
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>matrix = [["0"]]
<strong>输出：</strong>0
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>m == matrix.length</code></li> 
 <li><code>n == matrix[i].length</code></li> 
 <li><code>1 &lt;= m, n &lt;= 300</code></li> 
 <li><code>matrix[i][j]</code> 为 <code>'0'</code> 或 <code>'1'</code></li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>数组 | 动态规划 | 矩阵</details><br>

<div>👍 1382, 👎 0</div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1，[手把手刷二叉树系列课程](https://aep.xet.tech/s/3YGcq3) 上线，[第 17 期刷题打卡挑战](https://aep.xet.tech/s/2jPp5X) 下周开始，报名从速！。**

<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

关于动态规划的解题步骤和思维方法见前文 [动态规划核心套路](https://labuladong.github.io/article/fname.html?fname=动态规划详解进阶) 和 [动态规划答疑篇](https://labuladong.github.io/article/fname.html?fname=最优子结构)，这里就不赘述了，直接给出最关键的状态转移方程。

这道题不难，关键是你要观察出一个全是 1 的正方形有什么特点，如何根据小的正方形推导出大的正方形（状态转移方程）。

当 `matrix[i][j]` 为 1，且它的左边、上边、左上边都存在正方形时，`matrix[i][j]` 才能够作为一个更大的正方形的右下角：

![](https://labuladong.gitee.io/pictures/短题解/221.jpg)

PS：这个图是我基于测试用例修改的，你明白我的意思就行。

所以我们可以定义这样一个二维 `dp` 数组：

以 `matrix[i][j]` 为右下角元素的最大的全为 1 正方形矩阵的边长为 `dp[i][j]`。

有了这个定义，状态转移方程就是：

```java
if (matrix[i][j] == 1)
    // 类似「水桶效应」，最大边长取决于边长最短的那个正方形
    dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1;
else
    dp[i][j] = 0;
```

题目最终想要的答案就是最大边长 `max(dp[..][..])` 的平方。

**标签：[动态规划](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=1318881141113536512)**

## 解法代码

```java
class Solution {
    public int maximalSquare(char[][] matrix) {
        int m = matrix.length, n = matrix[0].length;
        // 定义：以 matrix[i][j] 为右下角元素的全为 1 正方形矩阵的最大边长为 dp[i][j]。
        int[][] dp = new int[m][n];

        // base case，第一行和第一列的正方形边长
        for (int i = 0; i < m; i++) {
            dp[i][0] = matrix[i][0] - '0';
        }
        for (int j = 0; j < n; j++) {
            dp[0][j] = matrix[0][j] - '0';
        }

        // 进行状态转移
        for (int i = 1; i < m; i++) {
            for (int j = 1; j < n; j++) {
                if (matrix[i][j] == '0') {
                    // 值为 0 不可能是正方形的右下角
                    continue;
                }
                dp[i][j] = Math.min(Math.min(
                    dp[i - 1][j],
                    dp[i][j - 1]),
                    dp[i - 1][j - 1]
                ) + 1;
            }
        }

        int len = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                len = Math.max(len, dp[i][j]);
            }
        }
        return len * len;
    }


}
```

</details>
</div>





