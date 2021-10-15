package xjsort

type UintList []uint

func (ints UintList) Len() int {
	return len(ints)
}

func (ints UintList) Less(i, j int) bool { //返回True时会调用Swap方法调换两个元素的位置
	return ints[i] > ints[j] // (i>j)，ints[i] < ints[j] 表示按升序排序，ints[i] > ints[j] 表示按降序排序
}

func (ints UintList) Swap(i, j int) {
	ints[i], ints[j] = ints[j], ints[i]
}

type UintDescList []uint

func (ints UintDescList) Len() int {
	return len(ints)
}

func (ints UintDescList) Less(i, j int) bool { //返回True时会调用Swap方法调换两个元素的位置
	return ints[i] < ints[j] // (i>j)，ints[i] < ints[j] 表示按升序排序，ints[i] > ints[j] 表示按降序排序
}

func (ints UintDescList) Swap(i, j int) {
	ints[i], ints[j] = ints[j], ints[i]
}
