func containsDuplicate(nums []int)bool {
    visited_map := make(map[int]bool)
    
    for i := 0;  i<len(nums); i++ {
        cur_num := nums[i]
        _, found := visited_map[cur_num]
        
        if found == true {
            return true
        }
        
        visited_map[cur_num] = true
        }
    return false
}