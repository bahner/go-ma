package internal

import "strings"

func Contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true // return true immediately if the item is found
		}
	}
	return false // return false if the item is not found
}

func GetFragmentFromDID(did string) string {
	elements := strings.Split(did, "#")

	return elements[len(elements)-1]
}
