package cn

import "strconv"

//leetcode submit region begin(Prohibit modification and deletion)
func getFolderNames(names []string) []string {
	m := make(map[string]int)
	res := make([]string, 0)
	for _, name := range names {
		if count, ok := m[name]; !ok {
			res = append(res, name)
			m[name] = 0
		} else {
			for i := count+1; ;i++ {
				newName := name + "(" + strconv.Itoa(i) + ")"
				if _, ok := m[newName]; !ok {
					res = append(res, newName)
					m[newName] = 0
					m[name] = i
					break
				}
			}
		}
	}
	return res
}
//leetcode submit region end(Prohibit modification and deletion)
