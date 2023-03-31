package 哈希表

func isHappy(n int) bool {
	m := make(map[int]bool)
	//如果重复便进入了循环，直接退出
	for n != 1 && !m[n] {
		n, m[n] = getSum(n), true
	}
	return n == 1
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		//算每一位的平方和
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
