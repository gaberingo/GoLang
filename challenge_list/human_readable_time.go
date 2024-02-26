package challenge_list

import "fmt"

func HumanReadableTime(seconds int) string {
	var sec = 1
	var min = sec * 60
	var hour = min * 60
	if seconds < 0 {
		return "00:00:00"
	}
	h := seconds / hour
	m := (seconds % hour) / min
	s := (seconds % hour) % min
	return formatTime(h) + ":" + formatTime(m) + ":" + formatTime(s)
}

func formatTime(t int) string {
	return fmt.Sprintf("%02d", t)
}
