func groupAnagrams(strs []string) [][]string {

	hashMap := make(map[string][]string)
	for i := 0; i < len(strs); i++ {
		slice_of_zeros := make([]int, 26)
		current_string := strs[i]
		for j := 0; j < len(current_string); j++ {
			letter := rune(current_string[j])
			index := int(letter) - int('a')
			slice_of_zeros[index] += 1
			}
		
		cur_group := strings.Join(strings.Fields(fmt.Sprint(slice_of_zeros)), ",")

		hashMap[cur_group] = append(hashMap[cur_group], current_string)
	}
	
	values := make([][]string, 0, len(hashMap))
	for _, v := range hashMap {
		
		values = append(values, v)
	}
	
	return values
}