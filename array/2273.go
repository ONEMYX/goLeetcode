package array

func removeAnagrams(words []string) []string {
	res := []string{words[0]}
	com := func(a, b string) bool {
		if len(a) != len(b) {
			return false
		}
		cnt := make([]int, 26)
		for _, ch := range a {
			cnt[ch-'a']++
		}
		for _, ch := range b {
			cnt[ch-'a']--
		}
		for _, v := range cnt {
			if v != 0 {
				return false
			}
		}
		return true
	}
	for i := 1; i < len(words); i++ {
		if !com(words[i], words[i-1]) {
			res = append(res, words[i])
		}
	}
	return res
}
