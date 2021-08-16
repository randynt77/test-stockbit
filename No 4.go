func groupAnagrams(str []string) map[[26]int][]string {
	mapResult := make(map[[26]int][]string, 0)
	for _, s := range str {
		var alfabetMap [26]int
		runeString := []rune(s)
		for _, runeValue := range runeString {
			alfabetMap[int(runeValue)-97] = 1
		}
		val, ok := mapResult[alfabetMap]
		if !ok {
			mapResult[alfabetMap] = []string{s}
			continue
		}
		val = append(val, s)
		mapResult[alfabetMap] = val

	}
	return mapResult
}

func main() {
	a := []string{"eta", "tae", "eat", "susu", "usus", "budi", "susi", "udib"}
	res := (groupAnagrams(a))
	for _, val := range res {
		fmt.Println(val)
	}

}
