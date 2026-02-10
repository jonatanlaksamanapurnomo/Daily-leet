func longestBalanced(nums []int) int {
    maxK := findK(nums)
    resp := 0

    for k := 0 ; k <= maxK ; k++{
        resp = max(resp , longestBalanceWindow(nums,k))
    }    

    return resp
}



func longestBalanceWindow(nums []int, k int) int {
    if k <= 0 {
        return 0
    }

    seen := make(map[int]int)
    left := 0
    right := 0
    totalOdd, totalEven := 0, 0
    resp := 0

    for right = 0; right < len(nums); right++ {
        val := nums[right]
        seen[val]++

        if seen[val] == 1 {
            if val%2 == 0 {
                totalEven++
            } else {
                totalOdd++
            }
        }

        for totalEven > k || totalOdd > k {
            valLeft := nums[left]
            seen[valLeft]--

            if seen[valLeft] == 0 {
                if valLeft % 2 == 0 {
                    totalEven--
                }else{
                    totalOdd--
                }
            }
            left++
        }

        if totalEven == k && totalOdd == k {
            resp = max(resp , right - left + 1)
        }
    }

    return resp
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}


func findK(nums []int) int {
    seen := make(map[int]bool)
    totalOdd , totalEven := 0,0
    for _,val := range nums {
        if !seen[val]{
            seen[val] = true

            if val % 2 == 0 {
                totalEven++
            } else {
                totalOdd++
            }
        }
    }
    return min(totalOdd,totalEven)
}

func min(a,b int ) int {
    if a < b {
        return a
    }
    return b
}