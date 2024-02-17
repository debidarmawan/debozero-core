package helper

func ContainsString(arr []string, item string) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}

	return false
}
