/*
 * Runtime: 4 ms
 * Memory Usage: 3.4 MB
 */
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
