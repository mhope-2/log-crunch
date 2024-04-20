package utils

func IsInList(s string, list []string) bool {
	for _, item := range list {
		if s == item {
			return true
		}
	}
	return false
}
