package main

// 阿基米德求最小公因数
// 辗转除余法
func god(m, n int) int {
	var rem = 0

	for n > 0 {
		rem = m % n
		m = n
		n = rem
	}
	return m
}
