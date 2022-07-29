func twoSum(nums []int, target int) []int {
    
    numbers_map := make(map[int]int)
    
    for i := 0; i < len(nums); i ++ {
        cur_number := nums[i]
        diff := target - cur_number
        _, found := numbers_map[diff]
        
        if found {
            return []int{numbers_map[diff], i}
        }
        numbers_map[cur_number] = i
    }
    return []int{}
}