package unittest

import (
	"sort"
	"strings"
)

func sortChar(str string) string {
	// split string into list of char
	res := strings.Split(str, "")
	// sort string's char
	sort.Strings(res)
	// convert to string
	return strings.Join(res, "")
}

func groupAnagram(strList []string) map[string][]string {
	anagramDict := make(map[string][]string)

	// group the word with sorted char for the key
	// e.g : dict { "ahmru" : [ "rumah", "murah" ] }
	for _, str := range strList {
		key := sortChar(str)
		anagramDict[key] = append(anagramDict[key], str)
	}

	return anagramDict
}