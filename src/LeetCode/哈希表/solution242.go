package 哈希表

func isAnagram(s string, t string) bool {
	record := [26]int{}

	for _, r := range s {
		record[r-'a']++
	}

	for _, r := range t {
		record[r-'a']--
	}

	// record数组如果有的元素不为零0，说明字符串s和t 一定是谁多了字符或者谁少了字符。
	return record == [26]int{}

}
