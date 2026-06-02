func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
    ans := math.MaxInt32

    minLandFinish , minWaterFinish := math.MaxInt32 , math.MaxInt32

    for idx := range landStartTime{
        minLandFinish = min(minLandFinish , landStartTime[idx] + landDuration[idx])
    }

    for idx := range waterStartTime{
        minWaterFinish = min(minWaterFinish , waterStartTime[idx] + waterDuration[idx])
    }


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