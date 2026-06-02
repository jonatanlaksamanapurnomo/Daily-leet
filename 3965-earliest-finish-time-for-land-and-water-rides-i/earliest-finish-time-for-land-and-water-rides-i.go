func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
    ans := 10000

    minLandFinish , minWaterFinish := 10000 , 10000

    for idx := range landStartTime{
        minLandFinish = min(minLandFinish , landStartTime[idx] + landDuration[idx])
    }

    for idx := range waterStartTime{
        minWaterFinish = min(minWaterFinish , waterStartTime[idx] + waterDuration[idx])
    }

    fmt.Println(minLandFinish ,minWaterFinish )

    //min finish + duration
    for idx,waterLeadDuration := range waterDuration{
        waitingTime := waterStartTime[idx] - minLandFinish
        if waitingTime <= 0 {
            waitingTime = 0
        }
        ans = min(ans , minLandFinish + waterLeadDuration + waitingTime )
    }

    for idx,landLeadDuration := range landDuration {
        waitingTime := landStartTime[idx] - minWaterFinish
        if waitingTime <= 0 {
            waitingTime = 0
        }
        ans = min(ans ,minWaterFinish + landLeadDuration + waitingTime)
    }
    return ans
}

func max(a,b int ) int {
    if a > b {
        return a
    }
    return b
}

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}