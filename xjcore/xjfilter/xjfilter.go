package xjfilter

func Int64ArrayFilter(all, filter []int64) (list []int64) {
	for _, item1 := range all {
		find := false
		for _, item2 := range filter {
			if item1 == item2 {
				find = true
				break
			}
		}
		if !find {
			list = append(list, item1)
		}
	}
	return
}
