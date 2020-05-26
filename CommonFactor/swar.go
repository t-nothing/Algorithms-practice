package main

// swar算法计算汉明重量
// https://blog.csdn.net/u010320108/article/details/60878085
func Swar(i uint32) uint8 {
	i = (i & 0x55555555) + ((i >> 1) & 0x55555555)
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333)
	i = (i & 0x0f0f0f0f) + ((i >> 4) & 0x0f0f0f0f)
	i = (i * 0x01010101) >> 24
	return uint8(i)
}
