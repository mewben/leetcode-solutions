/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 *
 * Runtime: 56 ms
 * Memory Usage: 34.3 MB
 */
var twoSum = function(nums, target) {
    const hashmap = {}
    for (let i = 0, l = nums.length; i < l; i++) {
        const complement = target - nums[i];
        if (hashmap[complement] !== undefined) {
            return [hashmap[complement], i]
        }
        hashmap[nums[i]] = i
    }
    
    return []
}
