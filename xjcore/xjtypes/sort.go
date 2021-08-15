package xjtypes

import "strings"

func SelectSortSring(arr []string) []string {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ { //只剩一个元素不需要索引
			min := i                          //标记索引
			for j := i + 1; j < length; j++ { //每次选出一个极小值
				/*if arr[min] < arr[j] {
					min = j //保存极小值的索引
				}*/
				if strings.Compare(arr[min], arr[j]) > 0 {
					min = j
				}
			}
			if i != min {
				arr[i], arr[min] = arr[min], arr[i] //数据交换
			}
		}
		return arr
	}
}

func SelectSortInt(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	} else {
		for i := 0; i < length-1; i++ { //只剩一个元素不需要索引
			min := i                          //标记索引
			for j := i + 1; j < length; j++ { //每次选出一个极小值
				if arr[min] < arr[j] {
					min = j //保存极小值的索引
				}
			}
			if i != min {
				arr[i], arr[min] = arr[min], arr[i] //数据交换
			}
		}
		return arr
	}
}
