package yop

func InStringArray(arr []string, el string) bool {
	for _, val := range arr {
		if val == el {
			return true
		}
	}
	return false
}
