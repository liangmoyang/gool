package math

import "unsafe"

// 返回a/b向上舍入最接近的整数		摘录于：Gopher指北
func divRoundUp(n, a uintptr) uintptr {
	return (n + a - 1) / a
}

// 判断x是否为2的n次幂		摘录于：Gopher指北
func isPowerOfTwo(x uintptr) bool {
	return x&(x-1) == 0
}

// 向上将x舍入为a的倍数，例如：x=6，a=4则返回值为8		摘录于：Gopher指北
func alignUp(x, a uintptr) uintptr {
	return (x + a - 1) &^ (a - 1)
}

// 向上将x舍入为a的倍数，例如：x=6，a=4则返回值为4		摘录于：Gopher指北
func alignDown(x, a uintptr) uintptr {
	return x &^ (a - 1)
}

// 布尔转整形		摘录于：Gopher指北
func bool2int(x bool) int {
	return int(uint8(*(*uint8)(unsafe.Pointer(&x))))
}
