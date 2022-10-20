<p>给你一个字符串 <code>s</code> ，请你反转字符串中 <strong>单词</strong> 的顺序。</p>

<p><strong>单词</strong> 是由非空格字符组成的字符串。<code>s</code> 中使用至少一个空格将字符串中的 <strong>单词</strong> 分隔开。</p>

<p>返回 <strong>单词</strong> 顺序颠倒且 <strong>单词</strong> 之间用单个空格连接的结果字符串。</p>

<p><strong>注意：</strong>输入字符串 <code>s</code>中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>s = "<span><code>the sky is blue</code></span>"
<strong>输出：</strong>"<span><code>blue is sky the</code></span>"
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>s = " &nbsp;hello world &nbsp;"
<strong>输出：</strong>"world hello"
<strong>解释：</strong>反转后的字符串中不能存在前导空格和尾随空格。
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>s = "a good &nbsp; example"
<strong>输出：</strong>"example good a"
<strong>解释：</strong>如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>1 &lt;= s.length &lt;= 10<sup>4</sup></code></li> 
 <li><code>s</code> 包含英文大小写字母、数字和空格 <code>' '</code></li> 
 <li><code>s</code> 中 <strong>至少存在一个</strong> 单词</li> 
</ul>

<ul> 
</ul>

<p>&nbsp;</p>

<p><strong>进阶：</strong>如果字符串在你使用的编程语言中是一种可变数据类型，请尝试使用&nbsp;<code>O(1)</code> 额外空间复杂度的 <strong>原地</strong> 解法。</p>

<details><summary><strong>Related Topics</strong></summary>双指针 | 字符串</details><br>

<div>👍 674, 👎 0</div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.0；[第 12 期刷题打卡](https://mp.weixin.qq.com/s/eUG2OOzY3k_ZTz-CFvtv5Q) 最后一天报名；点击这里体验 [刷题全家桶](https://labuladong.gitee.io/algo/images/others/%E5%85%A8%E5%AE%B6%E6%A1%B6.jpg)。**



<p><strong><a href="https://labuladong.github.io/article?qno=151" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

常规方法是用类似 `split` 再 `reverse` 最后 `join` 的方法得到结果，但更巧妙的方法是我在 [二维数组的花式遍历](https://labuladong.github.io/article/fname.html?fname=花式遍历) 中讲到的：**先把整个字符串进行翻转，再把每个单词中的字母翻转**。

比如说，给你输入这样一个字符串：

```shell
s = "hello world labuladong"
```

那么我们先将整个字符串 `s` 反转：

```shell
s = "gnodalubal dlrow olleh"
```

**然后将每个单词分别反转**：

```shell
s = "labuladong world hello"
```

这样，就实现了原地反转所有单词顺序的目的。

整体的思路应该不难，就是细节比较恶心，直接看我写的代码吧。

**详细题解：[二维数组的花式遍历技巧](https://labuladong.github.io/article/fname.html?fname=花式遍历)**

**标签：字符串，[数组双指针](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2120601117519675393)**

## 解法代码

```java
class Solution {
    public String reverseWords(String s) {
        StringBuilder sb = new StringBuilder();
        // 先清洗一下数据，把多于的空格都删掉
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (c != ' ') {
                // 单词中的字母/数字
                sb.append(c);
            } else if (!sb.isEmpty() && sb.charAt(sb.length() - 1) != ' ') {
                // 单词之间保留一个空格
                sb.append(' ');
            }
        }
        // 末尾如果有空格，清除之
        if (sb.charAt(sb.length() - 1) == ' ') {
            sb.deleteCharAt(sb.length() - 1);
        }

        // 清洗之后的字符串
        char[] chars = sb.toString().toCharArray();
        int n = chars.length;
        // 进行单词的翻转，先整体翻转
        reverse(chars, 0, n - 1);
        // 再把每个单词翻转
        for (int i = 0; i < n; ) {
            for (int j = i; j < n; j++) {
                if (j + 1 == n || chars[j + 1] == ' ') {
                    // chars[i..j] 是一个单词，翻转之
                    reverse(chars, i, j);
                    // 把 i 置为下一个单词的首字母
                    i = j + 2;
                    break;
                }
            }
        }
        // 最后得到题目想要的结果
        return new String(chars);
    }

    // 翻转 arr[i..j]
    void reverse(char[] arr, int i, int j) {
        while (i < j) {
            char temp = arr[i];
            arr[i] = arr[j];
            arr[j] = temp;
            i++;
            j--;
        }
    }
}
```

**类似题目**：
  - [1260. 二维网格迁移 🟢](/problems/shift-2d-grid)
  - [48. 旋转图像 🟠](/problems/rotate-image)
  - [54. 螺旋矩阵 🟠](/problems/spiral-matrix)
  - [59. 螺旋矩阵 II 🟠](/problems/spiral-matrix-ii)
  - [剑指 Offer 29. 顺时针打印矩阵 🟢](/problems/shun-shi-zhen-da-yin-ju-zhen-lcof)

</details>
</div>



