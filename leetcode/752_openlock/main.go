package main

import "fmt"

func openLock(deadends []string, target string) int {
	// 初始密码
	const initKeys = "0000"
	// 如果初始密码等于 target
	if initKeys == target {
		return 0
	}
	// 将 deadends 元素初始化到 visited 集合
	deadmap := map[string]bool{}
	for _, ele := range deadends {
		deadmap[ele] = true
	}
	// 如果初始密码存在于 deadmap, 直接返回
	if deadmap[initKeys] {
		return -1
	}
	// 初始化 queue
	queue := []string{initKeys}
	deadmap[initKeys] = true
	step := 0

	for len(queue) > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]

			for j := 0; j < 4; j++ {
				// 向上翻转
				curUp := upOne(cur, j)
				if curUp == target {
					return step + 1
				} else if !contain(deadmap, curUp) {
					queue = append(queue, curUp)
					deadmap[curUp] = true
				}
				// 向下翻转
				curDown := downOne(cur, j)
				if curDown == target {
					return step + 1
				} else if !contain(deadmap, curDown) {
					queue = append(queue, curDown)
					deadmap[curDown] = true
				}
			}
		}
		step++
	}

	return -1
}

func contain(colt map[string]bool, s string) bool {
	if _, ok := colt[s]; ok {
		return true
	}
	return false
}

func upOne(s string, j int) string {
	c := []byte(s)
	if c[j] == '9' {
		c[j] = '0'
	} else {
		c[j] += 1
	}
	return string(c)
}

func downOne(s string, j int) string {
	c := []byte(s)
	if c[j] == '0' {
		c[j] = '9'
	} else {
		c[j] -= 1
	}
	return string(c)
}

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
	fmt.Println(openLock([]string{"0000"}, "8888"))
}
