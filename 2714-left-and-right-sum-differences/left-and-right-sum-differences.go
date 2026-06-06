func leftRightDifference(nums []int) []int {
    prefixSum := []int{0}
    for i := 0 ; i< len(nums) -1; i++{
        prefixSum = append(prefixSum , nums[i] + prefixSum[i] )
    }

    suffixSum := make([]int , len(nums))
    for i := len(suffixSum) -1 ; i > 0 ; i-- {
        suffixSum[i-1] = suffixSum[i] + nums[i]
    }

    for i := 0 ; i < len(prefixSum) ; i++{
        prefixSum[i] = abs(prefixSum[i] - suffixSum[i])
    }
    return prefixSum
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}