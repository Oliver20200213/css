package render

// template模板中引用的Add函数
func Add(a int, b int) int {
	return a + b
}

// template模板中引用的In函数
func In(item int, arr []int) bool {
	for _, i := range arr {
		if i == item {
			return true
		}
	}
	return false
}
