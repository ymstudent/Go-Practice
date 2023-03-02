package cn

//leetcode submit region begin(Prohibit modification and deletion)
type Trie struct {
	isWord   bool      // 从根节点至此是否是一个完整的单词（即这个节点是否是一个单词的结尾）
	children [26]*Trie // 利用数组的下标作为26个字母，值则为子节点
}

func Constructor() Trie {
	return Trie{}
}

// 按word的字符，从根节点往下走，遇到nil，构建一个新的节点，如果节点已存在就继续往下走
func (this *Trie) Insert(word string) {
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if this.children[c] == nil {
			this.children[c] = new(Trie)
		}
		this = this.children[c]
	}
	this.isWord = true
}

// 按照word的字符，从根节点往下走，如果遇到节点为nil，表示单词不在前缀树中
// 如果顺利走到最后一个字符，判断这个字符的节点是不是一个完整的单词
func (this *Trie) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if this.children[c] == nil {
			return false
		}
		this = this.children[c]
	}
	return this.isWord
}

// 和上面一样，只是如果走到了最后一个字符，直接返回true即可，我们并不关系这个字符是不是单词结尾
func (this *Trie) StartsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		c := prefix[i] - 'a'
		if this.children[c] == nil {
			return false
		}
		this = this.children[c]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
//leetcode submit region end(Prohibit modification and deletion)
