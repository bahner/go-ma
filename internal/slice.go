package internal

func Contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true // return true immediately if the item is found
		}
	}
	return false // return false if the item is not found
}
