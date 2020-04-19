package two_sum

func twoSum(nums []int, target int) []int {
    res := make([]int, 2)
    ma := make(map[int]int, len(nums))

    for i, v := range nums {
        if j, ok := ma[target - v]; ok && i !=j  {
            res[0] = i
            res[1] = j
        }        
        ma[v] = i
    }

    return res
}