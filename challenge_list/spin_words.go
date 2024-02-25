package challenge_list

import (
	"strings"
)

func SpinWords(str string) string {
	var spin []string
	split := strings.Split(str, " ")
	for _, text := range split {
		if len(text) > 4 {
			var newString []byte
			for i := len(text); i > 0; i-- {
				newString = append(newString, text[i-1])
			}
			spin = append(spin, string(newString))
		} else {
			spin = append(spin, text)
		}
	}
	return strings.Join(spin, " ")
}
