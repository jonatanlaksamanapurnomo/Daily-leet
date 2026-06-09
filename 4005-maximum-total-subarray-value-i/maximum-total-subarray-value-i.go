func maxTotalValue(nums []int, k int) int64 {
    mostMax := int64(-1)
    mostMin := int64(math.MaxInt64)

    for _,val := range nums {
        mostMax = max(mostMax , int64(val))
        mostMin = min(mostMin , int64(val))
    }   

    diff := mostMax - mostMin
    return diff * int64(k)
}

func min(a,b int64) int64 {
    if a < b {
        return a
    }
    return b
}

func max(a,b int64) int64 {
    if a > b {
        return a
    }
    return b
}