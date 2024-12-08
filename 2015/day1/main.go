package day1

func findFloor(input string) int {
	start := 0

	for _, r := range input {
		switch r {
		case '(':
			start++
		case ')':
			start--
		}
	}
	return start
}

func firstEnterBasement(input string) int {
	start := 0

	for i, r := range input {
		switch r {
		case '(':
			start++
		case ')':
			start--
		}

		if start == -1 {
			return i + 1
		}
	}
	return 0
}
