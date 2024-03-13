package main

func letterCombinations(digits string) (res []string) {

	var (
		store = map[string][]string{
			"2": {"a", "b", "c"},
			"3": {"d", "e", "f"},
			"4": {"g", "h", "i"},
			"5": {"j", "k", "l"},
			"6": {"m", "n", "o"},
			"7": {"p", "q", "r", "s"},
			"8": {"t", "u", "v"},
			"9": {"w", "x", "y", "z"},
		}
		next []string
	)

	fc := func(a, b []string) []string {
		if len(a) == 0 {
			return b
		}

		if len(b) == 0 {
			return a
		}
		var (
			t []string
		)
		for _, v := range a {
			for _, _v := range b {
				t = append(t, v+_v)
			}
		}
		return t
	}

	for _, v := range digits {
		if val, ok := store[string(v)]; ok {
			next = fc(next, val)
		}
	}

	return next
}
