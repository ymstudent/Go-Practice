package cn

//leetcode submit region begin(Prohibit modification and deletion)
func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)            // 构建临接表(k是v的前置课程)， 课程1：[课程0， 课程2]
	inDegrees := make([]int, numCourses) // index：某课程，value：课程的入度

	for _, info := range prerequisites {
		after, prev := info[0], info[1]
		edges[prev] = append(edges[prev], after) // 通过临接表构建课程依赖关系
		inDegrees[after]++                       // 因为学习after之前需要先学prev，所以after的入度+1
	}

	queue := make([]int, 0) // bfs 队列
	for courseIndex, inDegree := range inDegrees {
		if inDegree == 0 { // 将入度为0的课程先压入队列
			queue = append(queue, courseIndex)
		}
	}

	learnedCourseCount := 0 //已学习课程数量
	for len(queue) > 0 {
		learnedCourseCount++                         // 学习课程数量+1
		learnedCourse := queue[0]                    //
		queue = queue[1:]                            // 弹出一个入度为0的课程
		for _, after := range edges[learnedCourse] { // 通过临接表取出该课程的后置课程
			inDegrees[after]--         // 因为learnedCourse已学，所以after对应的入度-1
			if inDegrees[after] == 0 { // 判断after的入度是否为0，即前置课程是否都已学完
				queue = append(queue, after) // 如果前置课程都学完，将其加入入度为0的课程队列
			}
		}
	}

	return learnedCourseCount == numCourses // 判断所有已学过的课程是否和所有必修课程数量相等
}

//leetcode submit region end(Prohibit modification and deletion)
