/**
Runtime: 4 ms, faster than 95.22% of Go online submissions for Two Sum.
Memory Usage: 3.4 MB, less than 40.38% of Go online submissions for Two Sum.
**/
func twoSum(nums []int, target int) []int {
    m := make(map[int]int, len(nums))
    for i, num := range nums {
        if v, found := m[target - num]; found {
            return []int{v,i}
        }
        m[num] = i
    }    
    return nil
}
