<div class="title__3Vvk">
 请你设计并实现一个满足&nbsp; 
 <a href="https://baike.baidu.com/item/LRU" target="_blank">LRU (最近最少使用) 缓存</a> 约束的数据结构。
</div>

<div class="title__3Vvk">
 实现 
 <code>LRUCache</code> 类：
</div>

<div class="original__bRMd"> 
 <div> 
  <ul> 
   <li><code>LRUCache(int capacity)</code> 以 <strong>正整数</strong> 作为容量&nbsp;<code>capacity</code> 初始化 LRU 缓存</li> 
   <li><code>int get(int key)</code> 如果关键字 <code>key</code> 存在于缓存中，则返回关键字的值，否则返回 <code>-1</code> 。</li> 
   <li><code>void put(int key, int value)</code>&nbsp;如果关键字&nbsp;<code>key</code> 已经存在，则变更其数据值&nbsp;<code>value</code> ；如果不存在，则向缓存中插入该组&nbsp;<code>key-value</code> 。如果插入操作导致关键字数量超过&nbsp;<code>capacity</code> ，则应该 <strong>逐出</strong> 最久未使用的关键字。</li> 
  </ul> 
 </div>
</div>

<p>函数 <code>get</code> 和 <code>put</code> 必须以 <code>O(1)</code> 的平均时间复杂度运行。</p>

<p>&nbsp;</p>

<p><strong>示例：</strong></p>

<pre>
<strong>输入</strong>
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
<strong>输出</strong>
[null, null, null, 1, null, -1, null, -1, 3, 4]

<strong>解释</strong>
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul> 
 <li><code>1 &lt;= capacity &lt;= 3000</code></li> 
 <li><code>0 &lt;= key &lt;= 10000</code></li> 
 <li><code>0 &lt;= value &lt;= 10<sup>5</sup></code></li> 
 <li>最多调用 <code>2 * 10<sup>5</sup></code> 次 <code>get</code> 和 <code>put</code></li> 
</ul>

<details><summary><strong>Related Topics</strong></summary>设计 | 哈希表 | 链表 | 双向链表</details><br>

<div>👍 2565, 👎 0</div>

<div id="labuladong"><hr>

**通知：[数据结构精品课](https://aep.h5.xeknow.com/s/1XJHEO) 已更新到 V2.1，[手把手刷二叉树系列课程](https://aep.xet.tech/s/3YGcq3) 上线，[第 17 期刷题打卡挑战](https://aep.xet.tech/s/2jPp5X) 下周开始，报名从速！。**



<p><strong><a href="https://labuladong.github.io/article?qno=146" target="_blank">⭐️labuladong 题解</a></strong></p>
<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

PS：这道题在[《算法小抄》](https://mp.weixin.qq.com/s/tUSovvogbR9StkPWb75fUw) 的第 216 页。

要让 `put` 和 `get` 方法的时间复杂度为 `O(1)`，我们可以总结出 `cache` 这个数据结构必要的条件：

1、显然 `cache` 中的元素必须有时序，以区分最近使用的和久未使用的数据，当容量满了之后要删除最久未使用的那个元素腾位置。

2、我们要在 `cache` 中快速找某个 `key` 是否已存在并得到对应的 `val`；

3、每次访问 `cache` 中的某个 `key`，需要将这个元素变为最近使用的，也就是说 `cache` 要支持在任意位置快速插入和删除元素。

哈希表查找快，但是数据无固定顺序；链表有顺序之分，插入删除快，但是查找慢，所以结合二者的长处，可以形成一种新的数据结构：哈希链表 `LinkedHashMap`：

![](https://labuladong.gitee.io/pictures/LRU算法/4.jpg)

至于 `put` 和 `get` 的具体逻辑，可以画出这样一个流程图：

![](https://labuladong.gitee.io/pictures/LRU算法/put.jpg)

根据上述逻辑写代码即可。

**详细题解：[算法就像搭乐高：带你手撸 LRU 算法](https://labuladong.github.io/article/fname.html?fname=LRU算法)**

**标签：[数据结构](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=1318892385270808576)，[设计](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2121998148662362112)**

## 解法代码

```java
class LRUCache {
    int cap;
    LinkedHashMap<Integer, Integer> cache = new LinkedHashMap<>();
    public LRUCache(int capacity) {
        this.cap = capacity;
    }

    public int get(int key) {
        if (!cache.containsKey(key)) {
            return -1;
        }
        // 将 key 变为最近使用
        makeRecently(key);
        return cache.get(key);
    }

    public void put(int key, int val) {
        if (cache.containsKey(key)) {
            // 修改 key 的值
            cache.put(key, val);
            // 将 key 变为最近使用
            makeRecently(key);
            return;
        }

        if (cache.size() >= this.cap) {
            // 链表头部就是最久未使用的 key
            int oldestKey = cache.keySet().iterator().next();
            cache.remove(oldestKey);
        }
        // 将新的 key 添加链表尾部
        cache.put(key, val);
    }

    private void makeRecently(int key) {
        int val = cache.get(key);
        // 删除 key，重新插入到队尾
        cache.remove(key);
        cache.put(key, val);
    }
}
```

**类似题目**：
  - [剑指 Offer II 031. 最近最少使用缓存 🟠](/problems/OrIXps)

</details>
</div>



