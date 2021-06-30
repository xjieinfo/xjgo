package xjtypes

// Contains checks if str is in list.
func ContainsStr(list []string, str string) bool {
	for _, each := range list {
		if each == str {
			return true
		}
	}
	return false
}

func ContainsInt(list []int, val int) bool {
	for _, each := range list {
		if each == val {
			return true
		}
	}
	return false
}

func ContainsInt64(list []int64, val int64) bool {
	for _, each := range list {
		if each == val {
			return true
		}
	}
	return false
}
