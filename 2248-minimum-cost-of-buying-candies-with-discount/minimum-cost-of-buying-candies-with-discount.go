//6,5,7,9,2,2
// 2,2,5,6,7,9
func minimumCost(cost []int) int {
    sort.Slice(cost, func(i, j int) bool {
        return cost[i] > cost[j]
    })

    total := 0
    for idx, val := range cost{
        if idx % 3 != 2 {
            total += val
        }
    }
    return total
}