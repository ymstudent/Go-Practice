<p>给你一个字符串数组，请你将 <strong>字母异位词</strong> 组合在一起。可以按任意顺序返回结果列表。</p>

<p><strong>字母异位词</strong> 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入:</strong> strs = <span><code>["eat", "tea", "tan", "ate", "nat", "bat"]</code></span>
<strong>输出: </strong>[["bat"],["nat","tan"],["ate","eat","tea"]]</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong> strs = <span><code>[""]</code></span>
<strong>输出: </strong>[[""]]
</pre>

<p><strong>示例 3:</strong></p>

<pre>
<strong>输入:</strong> strs = <span><code>["a"]</code></span>
<strong>输出: </strong>[["a"]]</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>1 &lt;= strs.length &lt;= 10<sup>4</sup></code></li> 
 <li><code>0 &lt;= strs[i].length &lt;= 100</code></li> 
 <li><code>strs[i]</code>&nbsp;仅包含小写字母</li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>数组 | 哈希表 | 字符串 | 排序</details><br>

<div>👍 1383, 👎 0</div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1，[手把手刷二叉树系列课程](https://aep.xet.tech/s/3YGcq3) 上线。**

<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

异位词这类问题的关键在于，你如何迅速判断两个字符串是异位词，主要考察数据编码和 [哈希表](https://appktavsiei5995.pc.xiaoe-tech.com/detail/p_6265484ae4b01a4851f65633/6) 的使用：

你是否可以找到一种编码方法，使得字母异位词的编码都相同？找到这种编码方式之后，就可以用一个哈希表存储编码相同的所有异位词，得到最终的答案。

[242. 有效的字母异位词](/problems/valid-anagram) 考察了异位词的编码问题，对字符串排序可以是一种编码方案，如果是异位词，排序后就变成一样的了，但是这样时间复杂度略高，且会修改原始数据。更好的编码方案是利用每个字符的出现次数进行编码，也就是下面的解法代码。

**标签：哈希表，[数组](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2120601117519675393)**

## 解法代码

```java
class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        // 编码到分组的映射
        HashMap<String, List<String>> codeToGroup = new HashMap<>();
        for (String s : strs) {
            // 对字符串进行编码
            String code = encode(s);
            // 把编码相同的字符串放在一起
            codeToGroup.putIfAbsent(code, new LinkedList<>());
            codeToGroup.get(code).add(s);
        }

        // 获取结果
        List<List<String>> res = new LinkedList<>();
        for (List<String> group : codeToGroup.values()) {
            res.add(group);
        }

        return res;
    }

    // 利用每个字符的出现次数进行编码
    String encode(String s) {
        char[] count = new char[26];
        for (char c : s.toCharArray()) {
            int delta = c - 'a';
            count[delta]++;
        }
        return new String(count);
    }
}
```

</details>
</div>



