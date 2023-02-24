<p>给定一个&nbsp;<code>m x n</code> 二维字符网格&nbsp;<code>board</code> 和一个字符串单词&nbsp;<code>word</code> 。如果&nbsp;<code>word</code> 存在于网格中，返回 <code>true</code> ；否则，返回 <code>false</code> 。</p>

<p>单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p> 
<img alt="" src="https://assets.leetcode.com/uploads/2020/11/04/word2.jpg" style="width: 322px; height: 242px;" /> 
<pre>
<strong>输入：</strong>board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
<strong>输出：</strong>true
</pre>

<p><strong>示例 2：</strong></p> 
<img alt="" src="https://assets.leetcode.com/uploads/2020/11/04/word-1.jpg" style="width: 322px; height: 242px;" /> 
<pre>
<strong>输入：</strong>board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
<strong>输出：</strong>true
</pre>

<p><strong>示例 3：</strong></p> 
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/15/word3.jpg" style="width: 322px; height: 242px;" /> 
<pre>
<strong>输入：</strong>board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
<strong>输出：</strong>false
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>m == board.length</code></li> 
 <li><code>n = board[i].length</code></li> 
 <li><code>1 &lt;= m, n &lt;= 6</code></li> 
 <li><code>1 &lt;= word.length &lt;= 15</code></li> 
 <li><code>board</code> 和 <code>word</code> 仅由大小写英文字母组成</li> 
</ul>

<p>&nbsp;</p>

<p><strong>进阶：</strong>你可以使用搜索剪枝的技术来优化解决方案，使其在 <code>board</code> 更大的情况下可以更快解决问题？</p>

<details><summary><strong>Related Topics</strong></summary>数组 | 回溯 | 矩阵</details><br>

<div>👍 1538, 👎 0</div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1，[手把手刷二叉树系列课程](https://aep.xet.tech/s/3YGcq3) 上线，[第 17 期刷题打卡挑战](https://aep.xet.tech/s/2jPp5X) 下周开始，报名从速！。**

<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

这道题和 [一文秒杀所有岛屿题目](https://labuladong.github.io/article/fname.html?fname=岛屿题目) 中讲到的题目非常类似，用 DFS 算法暴力搜索就行了。

注意我们要对已经匹配过的字符做标记，比如用 [图论算法基础](https://labuladong.github.io/article/fname.html?fname=图) 中讲到的 `visited` 数组，或者使用其他方法标记 `board` 中已经匹配过的字符。

如果不做标记的话会出现错误的结果，比如这个测试用例：

```java
[["A","B","C","E"],
 ["S","F","C","S"],
 ["A","D","E","E"]]
```

你如果在其中搜索 `"ABCB"`，按道理不应该搜到，但如果你不对已经匹配过的字符做标记的话，算法可能匹配了第一行前三列的 `"ABC"` 之后又回头匹配了第二列的 `"B"`，导致出现错误的结果。

我就用加负号的方式标记已经匹配过的字符，来避免走回头路，具体看代码吧。

**标签：[DFS 算法](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2122002916411604996)**

## 解法代码

```java
class Solution {

    boolean found = false;

    public boolean exist(char[][] board, String word) {
        int m = board.length, n = board[0].length;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                dfs(board, i, j, word, 0);
                if (found) {
                    return true;
                }
            }
        }
        return false;
    }

    // 从 (i, j) 开始向四周搜索，试图匹配 word[p..]
    void dfs(char[][] board, int i, int j, String word, int p) {
        if (p == word.length()) {
            // 整个 word 已经被匹配完，找到了一个答案
            found = true;
            return;
        }
        if (found) {
            // 已经找到了一个答案，不用再搜索了
            return;
        }
        int m = board.length, n = board[0].length;
        if (i < 0 || j < 0 || i >= m || j >= n) {
            return;
        }
        if (board[i][j] != word.charAt(p)) {
            return;
        }

        // 已经匹配过的字符，我们给它添一个负号作为标记，避免走回头路
        board[i][j] = (char)(-board[i][j]);
        // word[p] 被 board[i][j] 匹配，开始向四周搜索 word[p+1..]
        dfs(board, i + 1, j, word, p + 1);
        dfs(board, i, j + 1, word, p + 1);
        dfs(board, i - 1, j, word, p + 1);
        dfs(board, i, j - 1, word, p + 1);
        board[i][j] = (char)(-board[i][j]);
    }
}
```

</details>
</div>



