package main

import "fmt"

/*
给你一个可装载重量为 W 的背包和 N 个物品，每个物品有重量和价值两个属性。
其中第i个物品的重量为wt[i]，价值为val[i]，现在让你用这个背包装物品，最多能装的价值是多少？

举个简单的例子，输入如下：
N = 3, W = 4
wt = [2, 1, 3]
val = [4, 2, 3]
算法返回 6，选择前两件物品装进背包，总重量 3 小于W，可以获得最大价值 6。

解决方案:
	第一步: 明确两点 状态 和 选择
		状态: 背包的容量 和 可选择的物品
		选择: 装进背包 和 不装进背包
	第二步: 明确 dp 数组的定义 和 basecase
		dp 数组的定义: dp[i][w] = dp[N+1][W+1]
		dp 数组的意义: 对于前 i 个物品, 当前背包的容量为 w, 这种情况下可以装的最大价值
		basecase: dp[0][..] = dp[..][0] = 0。==> i、w 可以直接从 1 开始循环
		最终答案: dp[N][W]
	第三步: 根据 选择, 思考状态转移的逻辑
		不装: 如果没有把第 i 个物品装入背包, 那么 dp[i][w] = dp[i-1][w] (继承之前的结果)
		装: 如果把第 i 个物品装入了背包, 那么 dp[i][w] = dp[i-1][w-wt[i-1]] + val[i-1]
			由于 i 是从 1 开始的, 所以对 wt 和 val 的取值是 i-1
			dp[i-1][w-wt[i-1]] + val[i-1]: 寻求剩余重量 w-wt[i-1] 限制下能装的最大价值, 加上第i个物品的价值 val[i-1]
										   这就是装第 i 个物品的前提下, 背包可以装的最大价值
		择优(不装, 装)
*/

func knapsack(W, N int, wt, val []int) int {
	// 创建 dp 数组
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, W+1)
	}
	// 循环处理
	for i := 1; i <= N; i++ {
		for w := 1; w <= W; w++ {
			if w-wt[i-1] < 0 {
				// 当前背包容量装不下, 只能选择不装入背包
				dp[i][w] = dp[i-1][w]
			} else {
				// 装入 或者 不装入 背包, 择优
				dp[i][w] = Max(dp[i-1][w-wt[i-1]]+val[i-1], dp[i-1][w])
			}
		}
	}
	return dp[N][W]
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println("W:4 N:3 wt:[2, 1, 3] val:[4, 2, 3] => 6", knapsack(4, 3, []int{2, 1, 3}, []int{4, 2, 3}))
}
