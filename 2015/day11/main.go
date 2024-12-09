package day11

import "strings"

const all = "abcdefghijklmnopqrstuvwxyz"
const illegal = "ilo"

func nextValidPassword(pass string, straightLen int) string {
	pass = increment(pass)
	for !valid(pass, straightLen) {
		pass = increment(pass)
	}
	return pass
}

func valid(pass string, n int) bool {
	if strings.ContainsAny(pass, illegal) {
		return false
	}

	straight := false
	// pass has increasing straight of n chars i.e. abc, bcd, xyz
	for i := 0; i < len(all); i++ {
		if i+n > len(all) {
			return false
		}

		if strings.Contains(pass, all[i:i+n]) {
			straight = true
			break
		}
	}

	return straight && hasPairs(pass)
}

func hasPairs(pass string) bool {
	pairs := 0
	for _, a := range all {
		s := string(a) + string(a)
		if strings.Contains(pass, s) {
			pairs++
			if pairs >= 2 {
				return true
			}
		}
	}
	return false
}

func increment(pass string) string {
	i := len(pass) - 1
	wrap, wrapped := true, true
	seq := ""
	for wrap {
		s := pass[i]
		r := next(rune(s))
		seq = string(r) + seq
		wrap = r == 'a'
		if wrap {
			wrapped = true
		}
		i--
		if i < 0 {
			return "a" + seq
		}
	}
	if wrapped {
		return pass[:i+1] + seq
	} else {
		return pass[:i] + seq
	}
}

func next(r rune) rune {
	if r == 'z' {
		return 'a'
	}

	return rune(all[strings.IndexRune(all, r)+1])
}
